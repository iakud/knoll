package naming

import (
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

func (m *Manager[T]) addEndpoint(key string, endpoint endpoints.Endpoint) error {
	session, err := m.addSessionEndpoint(key, endpoint)
	if err != nil {
		return err
	}
	go m.keepSessionEndpoint(session, key, endpoint)
	return nil
}

func (m *Manager[T]) addSessionEndpoint(key string, endpoint endpoints.Endpoint) (*concurrency.Session, error) {
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

func (m *Manager[T]) keepSessionEndpoint(session *concurrency.Session, key string, endpoint endpoints.Endpoint) {
	err := m.keepAliveCtxCloser(session)
	if err != nil {
		return
	}

	var tempDelay time.Duration // how long to sleep on add failure
	for {
		session, err = m.addSessionEndpoint(key, endpoint)
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
		if err = m.keepAliveCtxCloser(session); err != nil {
			return
		}
	}
}

func (m *Manager[T]) keepAliveCtxCloser(session *concurrency.Session) error {
	select {
	case <-session.Done():
		return nil
	case <-m.ctx.Done():
		return m.ctx.Err()
	}
}
