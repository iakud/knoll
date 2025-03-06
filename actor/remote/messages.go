package remote

import "github.com/iakud/knoll/actor"

type RemoteUnreachableEvent struct {
	Address string
}

type remoteDeliver struct {
	target  *actor.PID
	sender  *actor.PID
	message any
}
