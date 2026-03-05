package naming

import (
	"context"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Watcher interface {
	UpdateState(endpoints []Endpoint)
}

type Resolver struct {
	manager *Manager
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
}

func NewResolver(client *clientv3.Client, target string, watcher Watcher) (*Resolver, error) {
	manager, err := NewManager(client, target)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())

	wch, err := manager.NewWatchChannel(ctx)
	if err != nil {
		return nil, err
	}
	r := &Resolver{
		manager: manager,
		ctx:     ctx,
		cancel:  cancel,
	}
	r.wg.Add(1)
	go r.watch(ctx, wch, watcher)
	return r, nil
}

func (r *Resolver) Close() {
	r.cancel()
	r.wg.Wait()
}

func (r *Resolver) watch(ctx context.Context, wch WatchChannel, watcher Watcher) {
	defer r.wg.Done()

	allUps := make(map[string]*Update)
	for {
		select {
		case <-ctx.Done():
			return
		case ups, ok := <-wch:
			if !ok {
				return
			}
			for _, up := range ups {
				switch up.Op {
				case Add:
					allUps[up.Key] = up

				case Delete:
					delete(allUps, up.Key)
				}
			}
			eps := convertToEndpoint(allUps)
			watcher.UpdateState(eps)
		}
	}
}

func convertToEndpoint(ups map[string]*Update) []Endpoint {
	var eps []Endpoint
	for _, up := range ups {
		ep := Endpoint{
			Addr:     up.Endpoint.Addr,
			Metadata: up.Endpoint.Metadata,
		}
		eps = append(eps, ep)
	}
	return eps
}
