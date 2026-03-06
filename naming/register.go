package naming

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type Register struct {
	client *clientv3.Client
	target string
	ctx    context.Context
	cancel context.CancelFunc
}

func RegisterEndpoint(client *clientv3.Client, target string, key string, endpoint Endpoint) (*Register, error) {
	if !strings.HasPrefix(key, target+"/") {
		return nil, fmt.Errorf("naming: register endpoint key should be prefixed with '%s/' got: '%s'", target, key)
	}

	ctx, cancel := context.WithCancel(context.Background())
	r := &Register{
		client: client,
		target: target,
		ctx:    ctx,
		cancel: cancel,
	}
	go r.register(key, endpoint)
	return r, nil
}

func (r *Register) Close() {
	r.cancel()
}

func (r *Register) register(key string, endpoint Endpoint) {
	var tempDelay time.Duration // how long to sleep on add endpoint failure
	for {
		select {
		case <-r.ctx.Done():
			return
		default:
		}
		session, err := concurrency.NewSession(r.client)
		if err != nil {
			if tempDelay == 0 {
				tempDelay = 5 * time.Millisecond
			} else {
				tempDelay *= 2
			}
			if max := 1 * time.Second; tempDelay > max {
				tempDelay = max
			}
			slog.Error(fmt.Sprintf("naming: register error: %v; retrying in %v", err, tempDelay))
			time.Sleep(tempDelay)
			continue
		}
		tempDelay = 0
		if err = r.registerEndpoint(session, key, endpoint); err != nil {
			return
		}
	}
}

func (r *Register) registerEndpoint(session *concurrency.Session, key string, endpoint Endpoint) error {
	defer session.Close()

	manager, err := endpoints.NewManager(r.client, r.target)
	if err != nil {
		return err
	}

	if err := manager.AddEndpoint(r.client.Ctx(), key, endpoints.Endpoint{Addr: endpoint.Addr, Metadata: endpoint.Metadata}, clientv3.WithLease(session.Lease())); err != nil {
		slog.Error("naming: add endpoint", "error", err, "key", key, "endpoint", endpoint)
		return nil
	}
	slog.Info("naming: register ok", "key", key, "endpoint", endpoint)

	defer func() {
		if err := manager.DeleteEndpoint(r.client.Ctx(), key, clientv3.WithLease(session.Lease())); err != nil {
			slog.Error("naming: delete endpoint", "error", err, "key", key)
		}
	}()

	select {
	case <-session.Done():
		return nil
	case <-r.ctx.Done():
		return r.ctx.Err()
	}
}
