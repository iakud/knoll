package naming

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type Manager struct {
	client *clientv3.Client
	target string
	ctx    context.Context
	cancel context.CancelFunc
}

func Register(client *clientv3.Client, target string, key string, endpoint Endpoint) (*Manager, error) {
	ctx, cancel := context.WithCancel(context.Background())
	r := &Manager{
		client: client,
		target: target,
		ctx:    ctx,
		cancel: cancel,
	}
	go r.register(key, endpoint)
	return r, nil
}

func (r *Manager) Close() {
	r.cancel()
}

func (r *Manager) register(key string, endpoint Endpoint) {
	for {
		if err := r.add(key, endpoint); err != nil {
			fmt.Printf("naming: watch error: %v", err)
			// FIXME： 增加重试间隔
			continue
		}
		return
	}
}

const defaultTTL = 60

func (r *Manager) add(key string, endpoint Endpoint) error {
	resp, err := r.client.Grant(r.ctx, defaultTTL)
	if err != nil {
		return err
	}
	manager, err := endpoints.NewManager(r.client, r.target)
	if err != nil {
		return err
	}
	ep := endpoints.Endpoint{Addr: endpoint.Addr}
	if endpoint.Attributes != nil {
		ep.Metadata = endpoint.Attributes.m
	}
	if err := manager.AddEndpoint(r.ctx, key, ep, clientv3.WithLease(resp.ID)); err != nil {
		return err
	}
	return r.keepAlive(resp.ID)
}

func (r *Manager) keepAlive(id clientv3.LeaseID) error {
	ctx, cancel := context.WithCancel(r.ctx)
	defer cancel()
	keepAlive, err := r.client.KeepAlive(ctx, id)
	if err != nil || keepAlive == nil {
		return err
	}
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		for range keepAlive {
			// eat messages until keep alive channel closes
		}
	}()

	select {
	case <-donec:
		return fmt.Errorf("service: Session closed")
	case <-r.ctx.Done():
		return nil
	}
}
