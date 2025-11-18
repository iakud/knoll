package naming

import (
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

func (m *Manager[T]) startAddEndpoint(key string, endpoint endpoints.Endpoint) error {
	go m.keepAliveEndpoint(key, endpoint)
	return nil
}

func (m *Manager[T]) addEndpoint(key string, endpoint endpoints.Endpoint) (*concurrency.Session, error) {
	session, err := concurrency.NewSession(m.client)
	if err != nil {
		return nil, err
	}
	if err := m.manager.AddEndpoint(m.ctx, key, endpoint, clientv3.WithLease(session.Lease())); err != nil {
		session.Close()
		return nil, err
	}
	return session, nil
}

func (m *Manager[T]) keepAliveEndpoint(key string, endpoint endpoints.Endpoint) {
	var tempDelay time.Duration // how long to sleep on add endpoint failure
	for {
		session, err := m.addEndpoint(key, endpoint)
		if err != nil {
			select {
			case <-m.ctx.Done():
				return
			default:
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Printf("naming: add endpoint error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
		}
		tempDelay = 0
		if err = m.keepAliveSessionCloser(session); err != nil {
			return
		}
	}
}

func (m *Manager[T]) keepAliveSessionCloser(session *concurrency.Session) error {
	defer session.Close()

	select {
	case <-session.Done():
		return nil
	case <-m.ctx.Done():
		return m.ctx.Err()
	}
}
