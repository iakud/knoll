package naming

import (
	"github.com/go-viper/mapstructure/v2"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type Watcher[T any] interface {
	UpdateState(endpoints []Endpoint[T])
}

func watch[T any](m *Manager[T], wch endpoints.WatchChannel, watcher Watcher[T]) {
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
			eps := convertToEndpoint[T](allUps)
			watcher.UpdateState(eps)
		}
	}
}

func convertToEndpoint[T any](ups map[string]*endpoints.Update) []Endpoint[T] {
	var eps []Endpoint[T]
	for _, up := range ups {
		ep := Endpoint[T]{
			Addr: up.Endpoint.Addr,
		}
		mapstructure.Decode(up.Endpoint.Metadata.(map[string]any), &ep.Metadata)
		eps = append(eps, ep)
	}
	return eps
}
