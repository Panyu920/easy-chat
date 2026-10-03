package websocketx

import "net/http"

type Dialoptions func(*dialoOption)

type dialoOption struct {
	pattern string
	header  http.Header
}

func newDialoOption(opts ...Dialoptions) *dialoOption {
	o := &dialoOption{
		pattern: "/ws",
		header:  http.Header{},
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func WithClientPattern(pattern string) Dialoptions {
	return func(o *dialoOption) {
		o.pattern = pattern
	}
}

func WithClientHeader(header http.Header) Dialoptions {
	return func(o *dialoOption) {
		o.header = header
	}
}
