package httpclient

import (
	"context"
	"io"
	"net/http"
)

type Request struct {
	URL     string
	Method  string
	Headers map[string]string
	Query   map[string]string
	Body    io.Reader
}

func RequestGet(url string, query map[string]string, headers map[string]string) *Request {
	return &Request{
		URL:     url,
		Method:  http.MethodGet,
		Headers: headers,
		Query:   query,
		Body:    nil,
	}
}

func RequestDelete(url string, query map[string]string, headers map[string]string) *Request {
	return &Request{
		URL:     url,
		Method:  http.MethodDelete,
		Headers: headers,
		Query:   query,
		Body:    nil,
	}
}

func RequestPost(url string, query map[string]string, body io.Reader, headers map[string]string) *Request {
	return &Request{
		URL:     url,
		Method:  http.MethodPost,
		Headers: headers,
		Query:   query,
		Body:    body,
	}
}

func RequestPatch(url string, query map[string]string, body io.Reader, headers map[string]string) *Request {
	return &Request{
		URL:     url,
		Method:  http.MethodPatch,
		Headers: headers,
		Query:   query,
		Body:    body,
	}
}

func RequestPut(url string, query map[string]string, body io.Reader, headers map[string]string) *Request {
	return &Request{
		URL:     url,
		Method:  http.MethodPut,
		Headers: headers,
		Query:   query,
		Body:    body,
	}
}

func (r *Request) Send(ctx context.Context, client Client, opts ...RequestOption) (*http.Response, error) {
	return client.Send(ctx, r, opts...)
}
