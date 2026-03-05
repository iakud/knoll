package resolver

import (
	"context"
	"sync"

	"github.com/iakud/knoll/naming/endpoints"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Watcher interface {
	UpdateState(endpoints []endpoints.Endpoint)
}

type Resolver struct {
	manager *endpoints.Manager
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
}

func NewResolver(client *clientv3.Client, target string, watcher Watcher) (*Resolver, error) {
	manager, err := endpoints.NewManager(client, target)
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

func (r *Resolver) watch(ctx context.Context, wch endpoints.WatchChannel, watcher Watcher) {
	defer r.wg.Done()

	allUps := make(map[string]*endpoints.Update)
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
				case endpoints.Add:
					allUps[up.Key] = up

				case endpoints.Delete:
					delete(allUps, up.Key)
				}
			}
			eps := convertToEndpoint(allUps)
			watcher.UpdateState(eps)
		}
	}
}

func convertToEndpoint(ups map[string]*endpoints.Update) []endpoints.Endpoint {
	var eps []endpoints.Endpoint
	for _, up := range ups {
		ep := endpoints.Endpoint{
			Addr:     up.Endpoint.Addr,
			Metadata: up.Endpoint.Metadata,
		}
		eps = append(eps, ep)
	}
	return eps
}
