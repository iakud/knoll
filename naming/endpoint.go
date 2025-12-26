package naming

import (
	"context"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type Endpoint struct {
	addr       string
	attributes map[string]any
}

func NewEndpoint() Endpoint {
	return Endpoint{attributes: make(map[string]any)}
}

func (e *Endpoint) Addr() string {
	return e.addr
}

func (e *Endpoint) SetAddr(addr string) {
	e.addr = addr
}

func (e *Endpoint) Get(key string) any {
	return e.attributes[key]
}

func (e *Endpoint) Set(key string, value any) {
	e.attributes[key] = value
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
	return m.startAddEndpoint(key, endpoints.Endpoint{Addr: endpoint.addr, Metadata: endpoint.attributes})
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
