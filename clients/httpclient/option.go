package httpclient

import "net/http"

const (
	jsonContentType = "application/json; charset=utf-8"
)

type RequestOption interface {
	Configure(*http.Request) *http.Request
}

type RequestSetQuery struct {
	Query map[string]string
}

func (opt RequestSetQuery) Configure(r *http.Request) *http.Request {
	q := r.URL.Query()
	handleKeyValuesFunc(q.Set, opt.Query)
	r.URL.RawQuery = q.Encode()

	return r
}

type RequestAddQuery struct {
	Query map[string]string
}

func (opt RequestAddQuery) Configure(r *http.Request) *http.Request {
	q := r.URL.Query()
	handleKeyValuesFunc(q.Add, opt.Query)
	r.URL.RawQuery = q.Encode()

	return r
}

type RequestSetHeaders struct {
	Headers map[string]string
}

func (opt RequestSetHeaders) Configure(r *http.Request) *http.Request {
	handleKeyValuesFunc(r.Header.Set, opt.Headers)

	return r
}

type RequestAddHeaders struct {
	Headers map[string]string
}

func (opt RequestAddHeaders) Configure(r *http.Request) *http.Request {
	handleKeyValuesFunc(r.Header.Add, opt.Headers)

	return r
}

type RequestContentType struct {
	ContentType string
}

func (opt RequestContentType) Configure(r *http.Request) *http.Request {
	r.Header.Set("Content-Type", opt.ContentType)
	return RequestSetHeaders{Headers: map[string]string{"Content-Type": opt.ContentType}}.Configure(r)
}

type RequestJSON struct{}

func (opt RequestJSON) Configure(r *http.Request) *http.Request {
	return RequestContentType{ContentType: jsonContentType}.Configure(r)
}

// helpers

func handleKeyValuesFunc(f func(string, string), m map[string]string) {
	for k, v := range m {
		f(k, v)
	}
}
