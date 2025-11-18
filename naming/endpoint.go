package naming

import (
	"context"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type Endpoint[T any] struct {
	Addr     string
	Metadata T
}

type Manager[T any] struct {
	client  *clientv3.Client
	manager endpoints.Manager
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
}

func NewManager[T any](client *clientv3.Client, target string) (*Manager[T], error) {
	manager, err := endpoints.NewManager(client, target)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())

	m := &Manager[T]{
		client:  client,
		manager: manager,
		ctx:     ctx,
		cancel:  cancel,
	}
	return m, nil
}

func (m *Manager[T]) Close() {
	m.cancel()
	m.wg.Wait()
}

func (m *Manager[T]) Context() context.Context {
	return m.ctx
}

func (m *Manager[T]) AddEndpoint(key string, endpoint Endpoint[T]) error {
	return m.addEndpoint(key, endpoints.Endpoint{Addr: endpoint.Addr, Metadata: endpoint.Metadata})
}

func (m *Manager[T]) Watch(watcher Watcher[T]) error {
	wch, err := m.manager.NewWatchChannel(m.ctx)
	if err != nil {
		return err
	}
	m.wg.Add(1)
	go watch(m, wch, watcher)
	return nil
}
