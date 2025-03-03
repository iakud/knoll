package remote

import "github.com/iakud/knoll/actor"

type Remote struct {
	addr      string
	system    *actor.System
	routerPID *actor.PID
}
