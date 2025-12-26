package naming

import (
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type Watcher interface {
	UpdateState(endpoints []Endpoint)
}

func watch(m *Manager, wch endpoints.WatchChannel, watcher Watcher) {
	defer m.wg.Done()

	allUps := make(map[string]*endpoints.Update)
	for {
		select {
		case <-m.ctx.Done():
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

func convertToEndpoint(ups map[string]*endpoints.Update) []Endpoint {
	var eps []Endpoint
	for _, up := range ups {
		ep := Endpoint{
			addr:       up.Endpoint.Addr,
			attributes: up.Endpoint.Metadata.(map[string]any),
		}
		eps = append(eps, ep)
	}
	return eps
}
