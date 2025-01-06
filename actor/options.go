package actor

import "log/slog"

type options struct {
	remote Remoter
	logger *slog.Logger
}

var defaultOptions = options{}

type Option func(*options)

func WithRemote(remote Remoter) Option {
	return func(opts *options) {
		opts.remote = remote
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(opts *options) {
		opts.logger = logger
	}
}
