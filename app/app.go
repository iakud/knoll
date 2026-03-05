package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var ctx, cancel = context.WithCancel(context.Background())

func Run() error {
	if err := initServices(); err != nil {
		return err
	}
	if err := startServices(); err != nil {
		return err
	}
	defer stopServices()

	//FIXME: signal.NotifyContext()
	// signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	select {
	case s := <-ch:
		slog.Info("app", "signal", s)
	case <-ctx.Done():
		slog.Info("app", "done", ctx.Err())
	}
	return nil
}

func Stop() error {
	cancel()
	return nil
}
