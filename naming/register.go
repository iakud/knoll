package naming

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func Register[T any](client *clientv3.Client, target string, key string, endpoint Endpoint[T]) (*Manager[T], error) {
	manager, err := NewManager[T](client, target)
	if err != nil {
		return nil, err
	}
	go register(client, manager, key, endpoint)
	return manager, nil
}

func register[T any](client *clientv3.Client, m *Manager[T], key string, endpoint Endpoint[T]) {
	for {
		if err := add(client, m, key, endpoint); err != nil {
			fmt.Printf("naming: addEndpoint error: %v", err)
			// FIXME： 增加重试间隔
			continue
		}
		return
	}
}

const defaultTTL = 60

func add[T any](client *clientv3.Client, m *Manager[T], key string, endpoint Endpoint[T]) error {
	resp, err := client.Grant(m.Context(), defaultTTL)
	if err != nil {
		return err
	}
	if err := m.AddEndpoint(key, endpoint, clientv3.WithLease(resp.ID)); err != nil {
		return err
	}
	return keepAlive(m.Context(), client, resp.ID)
}

func keepAlive(ctx context.Context, client *clientv3.Client, id clientv3.LeaseID) error {
	kaCtx, kaCancel := context.WithCancel(ctx)
	defer kaCancel()
	keepAlive, err := client.KeepAlive(kaCtx, id)
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
		return fmt.Errorf("naming: keep alive channel closed")
	case <-ctx.Done():
		return nil
	}
}
