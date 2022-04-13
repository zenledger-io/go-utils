package httpclient

import (
	"context"
	"io"
	"net"
	"net/http"
)

type Client interface {
	Do(r *http.Request, opts ...RequestOption) (*http.Response, error)
	Send(ctx context.Context, method, url string, body io.Reader, opts ...RequestOption) (*http.Response, error)
	SendRequest(ctx context.Context, r *Request, opts ...RequestOption) (*http.Response, error)
	Get(ctx context.Context, url string, opts ...RequestOption) (*http.Response, error)
	Delete(ctx context.Context, url string, opts ...RequestOption) (*http.Response, error)
	Post(ctx context.Context, url string, body io.Reader, opts ...RequestOption) (*http.Response, error)
	Put(ctx context.Context, url string, body io.Reader, opts ...RequestOption) (*http.Response, error)
	Patch(ctx context.Context, url string, body io.Reader, opts ...RequestOption) (*http.Response, error)
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

func (c *client) Send(ctx context.Context, method, url string, body io.Reader, opts ...RequestOption) (*http.Response, error) {
	req := &Request{
		URL:    url,
		Method: method,
		Body:   body,
	}

	return c.SendRequest(ctx, req, opts...)
}

func (c *client) SendRequest(ctx context.Context, r *Request, opts ...RequestOption) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, r.Method, r.URL, r.Body)
	if err != nil {
		return nil, err
	}

	prependOpts := r.Options()
	if len(prependOpts) > 0 {
		opts = append(prependOpts, opts...)
	}

	return c.Do(req, opts...)
}

func (c *client) Get(ctx context.Context, url string, opts ...RequestOption) (*http.Response, error) {
	return c.Send(ctx, http.MethodGet, url, nil, opts...)
}

func (c *client) Delete(ctx context.Context, url string, opts ...RequestOption) (*http.Response, error) {
	return c.Send(ctx, http.MethodDelete, url, nil, opts...)
}

func (c *client) Post(ctx context.Context, url string, body io.Reader, opts ...RequestOption) (*http.Response, error) {
	return c.Send(ctx, http.MethodPost, url, body, opts...)
}

func (c *client) Put(ctx context.Context, url string, body io.Reader, opts ...RequestOption) (*http.Response, error) {
	return c.Send(ctx, http.MethodPut, url, body, opts...)
}

func (c *client) Patch(ctx context.Context, url string, body io.Reader, opts ...RequestOption) (*http.Response, error) {
	return c.Send(ctx, http.MethodPatch, url, body, opts...)
}
