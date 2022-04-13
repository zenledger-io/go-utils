package httpclient

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Get(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewDefault()

	tcs := map[string]struct {
		Query   map[string]string
		Headers map[string]string
		Opts    []RequestOption
		Handler http.HandlerFunc
	}{
		"when all request params are present": {
			Query:   map[string]string{"q1": "a"},
			Headers: map[string]string{"h1": "b"},
			Opts:    []RequestOption{RequestJSONContentType{}},
			Handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				q := r.URL.Query()
				if q.Get("q1") != "a" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				if r.Header.Get("h1") != "b" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				if r.Header.Get("Content-Type") != ContentTypeJSON {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				w.WriteHeader(http.StatusOK)
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			srv := httptest.NewServer(tc.Handler)
			defer srv.Close()

			opts := append(tc.Opts, RequestSetQuery{Query: tc.Query}, RequestSetHeaders{Headers: tc.Headers})
			resp, err := c.Get(ctx, srv.URL, opts...)
			require.Equal(t, http.StatusOK, resp.StatusCode)
			require.NoError(t, err)
		})
	}
}
