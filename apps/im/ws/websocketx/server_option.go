package websocketx

type ServerOptions func(*serverOption)
type serverOption struct {
	pattern string
	auth    IAuth
}

func newServerOptions(opts ...ServerOptions) *serverOption {
	defaultOpt := &serverOption{
		pattern: "/ws",
		auth:    &DefaultAuth{},
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
