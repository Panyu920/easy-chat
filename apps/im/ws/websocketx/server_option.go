package websocketx

import "time"

type ServerOptions func(*serverOption)
type serverOption struct {
	pattern             string
	auth                IAuth
	maxConnIdleDuration time.Duration
	maxTaskConcurrency  int
}

func newServerOptions(opts ...ServerOptions) *serverOption {
	defaultOpt := &serverOption{
		pattern:             "/ws",
		auth:                &DefaultAuth{},
		maxConnIdleDuration: 10000 * time.Second,
		maxTaskConcurrency:  50,
	}

	for _, opt := range opts {
		opt(defaultOpt)
	}

	return defaultOpt
}

func WithServerAuthOption(auth IAuth) ServerOptions {
	return func(opt *serverOption) {
		opt.auth = auth
	}
}

func WithServerMaxTaskConcurrency(maxTaskConcurrency int) ServerOptions {
	return func(opt *serverOption) {
		opt.maxTaskConcurrency = maxTaskConcurrency
	}
}
