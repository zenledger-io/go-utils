package httpclient

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestGet(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewDefault()

	tcs := map[string]struct {
		Query       map[string]string
		Headers     map[string]string
		Opts        []RequestOption
		Handler     http.HandlerFunc
		ConfirmFunc func(*testing.T, *http.Response, error)
	}{
		"when all request params are present": {
			Query:   map[string]string{"q1": "a"},
			Headers: map[string]string{"h1": "b"},
			Opts:    []RequestOption{RequestJSON{}},
			Handler: func(w http.ResponseWriter, r *http.Request) {
				q := r.URL.Query()
				if q.Get("q1") != "a" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				if r.Header.Get("h1") != "b" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				if r.Header.Get("Content-Type") != jsonContentType {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				w.WriteHeader(http.StatusOK)
			},
			ConfirmFunc: func(t *testing.T, r *http.Response, err error) {
				require.Equal(t, http.StatusOK, r.StatusCode)
				require.NoError(t, err)
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			srv := httptest.NewServer(tc.Handler)
			defer srv.Close()

			req := RequestGet(srv.URL, tc.Query, tc.Headers)
			resp, err := req.Send(ctx, c, tc.Opts...)
			tc.ConfirmFunc(t, resp, err)
		})
	}
}
