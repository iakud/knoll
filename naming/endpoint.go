package naming

import (
	"context"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type Endpoint struct {
	Addr     string
	Metadata any
}

type Manager struct {
	client  *clientv3.Client
	manager endpoints.Manager
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
}

func NewManager(client *clientv3.Client, target string) (*Manager, error) {
	manager, err := endpoints.NewManager(client, target)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())

	m := &Manager{
		client:  client,
		manager: manager,
		ctx:     ctx,
		cancel:  cancel,
	}
	return m, nil
}

func (m *Manager) Close() {
	m.cancel()
	m.wg.Wait()
}

func (m *Manager) Context() context.Context {
	return m.ctx
}

func (m *Manager) AddEndpoint(key string, endpoint Endpoint) error {
	return m.startAddEndpoint(key, endpoints.Endpoint{Addr: endpoint.Addr, Metadata: endpoint.Metadata})
}

func (m *Manager) Watch(watcher Watcher) error {
	wch, err := m.manager.NewWatchChannel(m.ctx)
	if err != nil {
		return err
	}
	m.wg.Add(1)
	go watch(m, wch, watcher)
	return nil
}
