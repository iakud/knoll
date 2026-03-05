package naming

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type Register struct {
	client  *clientv3.Client
	manager *Manager
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewRegister(client *clientv3.Client, target string, key string, endpoint Endpoint) (*Register, error) {
	if !strings.HasPrefix(key, target+"/") {
		return nil, fmt.Errorf("register: endpoint key should be prefixed with '%s/' got: '%s'", target, key)
	}
	manager, err := NewManager(client, target)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	r := &Register{
		client:  client,
		manager: manager,
		ctx:     ctx,
		cancel:  cancel,
	}
	go r.registerEndpoint(key, endpoint)
	return r, nil
}

func (r *Register) Close() {
	r.cancel()
}

func (r *Register) registerEndpoint(key string, endpoint Endpoint) {
	var tempDelay time.Duration // how long to sleep on add endpoint failure
	for {
		session, err := concurrency.NewSession(r.client)
		if err != nil {
			select {
			case <-r.ctx.Done():
				return
			default:
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				slog.Error(fmt.Sprintf("register: session error: %v; retrying in %v", err, tempDelay))
				time.Sleep(tempDelay)
				continue
			}
		}
		tempDelay = 0
		if err = r.serveEndpoint(session, key, endpoint); err != nil {
			return
		}
	}
}

func (r *Register) serveEndpoint(session *concurrency.Session, key string, endpoint Endpoint) error {
	defer session.Close()

	if err := r.manager.AddEndpoint(r.ctx, key, endpoint, clientv3.WithLease(session.Lease())); err != nil {
		slog.Error("register: add endpoint", "error", err, "key", key, "endpoint", endpoint)
		return nil
	}
	slog.Info("register: ok", "key", key, "endpoint", endpoint)

	defer func() {
		if err := r.manager.DeleteEndpoint(r.ctx, key, clientv3.WithLease(session.Lease())); err != nil {
			slog.Error("register: delete endpoint", "error", err, "key", key)
		}
	}()

	select {
	case <-session.Done():
		return nil
	case <-r.ctx.Done():
		return r.ctx.Err()
	}
}
