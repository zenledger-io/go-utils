package httpclient

import (
	"context"
	"net"
	"net/http"
)

type Client interface {
	Do(r *http.Request, opts ...RequestOption) (*http.Response, error)
	Send(ctx context.Context, r *Request, opts ...RequestOption) (*http.Response, error)
}

func New(cfg Config) Client {
	var rt http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   cfg.DialTimeout,
			KeepAlive: cfg.KeepAlive,
		}).DialContext,
		MaxIdleConns:          cfg.MaxIdleConns,
		IdleConnTimeout:       cfg.IdleConnTimeout,
		TLSHandshakeTimeout:   cfg.TLSHandshakeTimeout,
		ExpectContinueTimeout: cfg.ExpectContinueTimeout,
		ResponseHeaderTimeout: cfg.ResponseHeaderTimeout,
	}

	return &client{
		Client: &http.Client{
			Transport: rt,
			Timeout:   cfg.Timeout,
		},
	}
}

func NewDefault() Client {
	return New(DefaultConfig())
}

type client struct {
	*http.Client
}

func (c *client) Do(r *http.Request, opts ...RequestOption) (*http.Response, error) {
	for _, opt := range opts {
		r = opt.Configure(r)
	}

	return c.Client.Do(r)
}

func (c *client) Send(ctx context.Context, r *Request, opts ...RequestOption) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, r.Method, r.URL, r.Body)
	if err != nil {
		return nil, err
	}

	prependOpts := make([]RequestOption, 0, 2)

	if len(r.Query) > 0 {
		prependOpts = append(prependOpts, RequestSetQuery{Query: r.Query})
	}

	if len(r.Headers) > 0 {
		prependOpts = append(prependOpts, RequestSetHeaders{Headers: r.Headers})
	}

	if len(prependOpts) > 0 {
		opts = append(prependOpts, opts...)
	}

	return c.Do(req, opts...)
}
