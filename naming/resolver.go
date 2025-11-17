package naming

import (
	"context"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type ResolverState interface {
	UpdateState(Endpoints []endpoints.Endpoint)
}

type Resolver struct {
	client *clientv3.Client
	target string
	state  ResolverState
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func Resolve(client *clientv3.Client, target string, state ResolverState) (*Resolver, error) {
	ctx, cancel := context.WithCancel(context.Background())
	r := &Resolver{
		client: client,
		target: target,
		state:  state,
		ctx:    ctx,
		cancel: cancel,
	}

	manager, err := endpoints.NewManager(client, target)
	if err != nil {
		return nil, err
	}
	wch, err := manager.NewWatchChannel(ctx)
	if err != nil {
		return nil, err
	}

	r.wg.Add(1)
	go r.watch(wch)
	return r, nil
}

func (r *Resolver) Close() {
	r.cancel()
	r.wg.Wait()
}

func (r *Resolver) watch(wch endpoints.WatchChannel) {
	defer r.wg.Done()

	allUps := make(map[string]*endpoints.Update)
	for {
		select {
		case <-r.ctx.Done():
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

			var eps []endpoints.Endpoint
			for _, up := range ups {
				ep := endpoints.Endpoint{
					Addr:     up.Endpoint.Addr,
					Metadata: up.Endpoint.Metadata,
				}
				eps = append(eps, ep)
			}
			r.state.UpdateState(eps)
		}
	}
}
