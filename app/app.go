package app

import (
	"context"
	"os/signal"
	"syscall"
)

var ctx, stop = signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

func Run() error {
	if err := initServices(); err != nil {
		return err
	}
	if err := startServices(); err != nil {
		return err
	}
	defer stopServices()

	// signal
	<-ctx.Done()
	return ctx.Err()
}

func Stop() error {
	stop()
	return nil
}
