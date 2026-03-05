package naming

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"github.com/iakud/knoll/naming/internal"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Endpoint struct {
	Addr     string
	Metadata any
}

type Operation uint8

const (
	// Add indicates a new address is added.
	Add Operation = iota
	// Delete indicates an existing address is deleted.
	Delete
)

type Update struct {
	Op       Operation
	Key      string
	Endpoint Endpoint
}

type WatchChannel <-chan []*Update

type UpdateWithOpts struct {
	Update
	Opts []clientv3.OpOption
}

func NewAddUpdateOpts(key string, endpoint Endpoint, opts ...clientv3.OpOption) *UpdateWithOpts {
	return &UpdateWithOpts{Update: Update{Op: Add, Key: key, Endpoint: endpoint}, Opts: opts}
}

func NewDeleteUpdateOpts(key string, opts ...clientv3.OpOption) *UpdateWithOpts {
	return &UpdateWithOpts{Update: Update{Op: Delete, Key: key}, Opts: opts}
}

type Manager struct {
	client *clientv3.Client
	target string
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewManager(client *clientv3.Client, target string) (*Manager, error) {
	if client == nil {
		return nil, errors.New("invalid etcd client")
	}
	if target == "" {
		return nil, errors.New("invalid target")
	}

	ctx, cancel := context.WithCancel(context.Background())

	m := &Manager{
		client: client,
		target: target,
		ctx:    ctx,
		cancel: cancel,
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

func (m *Manager) Update(ctx context.Context, updates []*UpdateWithOpts) (err error) {
	ops := make([]clientv3.Op, 0, len(updates))
	for _, update := range updates {
		if !strings.HasPrefix(update.Key, m.target+"/") {
			return fmt.Errorf("endpoints: endpoint key should be prefixed with '%s/' got: '%s'", m.target, update.Key)
		}

		switch update.Op {
		case Add:
			internalUpdate := &internal.Update{
				Op:       internal.Add,
				Addr:     update.Endpoint.Addr,
				Metadata: update.Endpoint.Metadata,
			}

			var v []byte
			if v, err = json.Marshal(internalUpdate); err != nil {
				return err
			}
			ops = append(ops, clientv3.OpPut(update.Key, string(v), update.Opts...))
		case Delete:
			ops = append(ops, clientv3.OpDelete(update.Key, update.Opts...))
		default:
			return fmt.Errorf("endpoints: bad update op")
		}
	}
	_, err = m.client.KV.Txn(ctx).Then(ops...).Commit()
	return err
}

func (m *Manager) AddEndpoint(ctx context.Context, key string, endpoint Endpoint, opts ...clientv3.OpOption) error {
	return m.Update(ctx, []*UpdateWithOpts{NewAddUpdateOpts(key, endpoint, opts...)})
}

func (m *Manager) DeleteEndpoint(ctx context.Context, key string, opts ...clientv3.OpOption) error {
	return m.Update(ctx, []*UpdateWithOpts{NewDeleteUpdateOpts(key, opts...)})
}

func (m *Manager) NewWatchChannel(ctx context.Context) (WatchChannel, error) {
	key := m.target + "/"
	resp, err := m.client.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}

	initUpdates := make([]*Update, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		var iup internal.Update
		if err := json.Unmarshal(kv.Value, &iup); err != nil {
			slog.Warn("unmarshal endpoint update failed", "key", string(kv.Key), "error", err)
			continue
		}
		up := &Update{
			Op:       Add,
			Key:      string(kv.Key),
			Endpoint: Endpoint{Addr: iup.Addr, Metadata: iup.Metadata},
		}
		initUpdates = append(initUpdates, up)
	}

	upch := make(chan []*Update, 1)
	if len(initUpdates) > 0 {
		upch <- initUpdates
	}
	go m.watch(ctx, resp.Header.Revision+1, upch)
	return upch, nil
}

func (m *Manager) watch(ctx context.Context, rev int64, upch chan []*Update) {
	defer close(upch)

	opts := []clientv3.OpOption{clientv3.WithRev(rev), clientv3.WithPrefix()}
	key := m.target + "/"
	wch := m.client.Watch(ctx, key, opts...)
	for {
		select {
		case <-ctx.Done():
			return
		case wresp, ok := <-wch:
			if !ok {
				slog.Warn("watch closed", "target", m.target)
				return
			}
			if wresp.Err() != nil {
				slog.Warn("watch failed", "target", m.target, "error", wresp.Err())
				return
			}

			deltaUps := make([]*Update, 0, len(wresp.Events))
			for _, e := range wresp.Events {
				var iup internal.Update
				var err error
				var op Operation
				switch e.Type {
				case clientv3.EventTypePut:
					err = json.Unmarshal(e.Kv.Value, &iup)
					op = Add
					if err != nil {
						slog.Warn("unmarshal endpoint update failed", "key", string(e.Kv.Key), "error", err)
						continue
					}
				case clientv3.EventTypeDelete:
					iup = internal.Update{Op: internal.Delete}
					op = Delete
				default:
					continue
				}
				up := &Update{Op: op, Key: string(e.Kv.Key), Endpoint: Endpoint{Addr: iup.Addr, Metadata: iup.Metadata}}
				deltaUps = append(deltaUps, up)
			}
			if len(deltaUps) > 0 {
				upch <- deltaUps
			}
		}
	}
}
