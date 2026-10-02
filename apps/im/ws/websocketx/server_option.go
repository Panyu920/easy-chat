package websocketx

import "time"

type ServerOptions func(*serverOption)
type serverOption struct {
	pattern             string
	auth                IAuth
	maxConnIdleDuration time.Duration
}

func newServerOptions(opts ...ServerOptions) *serverOption {
	defaultOpt := &serverOption{
		pattern:             "/ws",
		auth:                &DefaultAuth{},
		maxConnIdleDuration: 10000 * time.Second,
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
