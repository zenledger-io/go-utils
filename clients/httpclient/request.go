package httpclient

import (
	"io"
)

type Request struct {
	URL     string
	Method  string
	Body    io.Reader
	Headers map[string]string
	Query   map[string]string
	Opts    []RequestOption
}

func (r *Request) Options() []RequestOption {
	opts := make([]RequestOption, 0, len(r.Opts)+2)

	if len(r.Opts) > 0 {
		opts = append(opts, r.Opts...)
	}

	if len(r.Query) > 0 {
		opts = append(opts, RequestSetQuery{Query: r.Query})
	}

	if len(r.Headers) > 0 {
		opts = append(opts, RequestSetHeaders{Headers: r.Headers})
	}

	return opts
}
