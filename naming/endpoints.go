package naming

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Endpoint struct {
}

type Manager struct {
	cancel context.CancelFunc
}

func AddEndpoint(client *clientv3.Client) {
	// ctx, cancel := context.WithCancel(context.Background())
}
