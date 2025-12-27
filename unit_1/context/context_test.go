package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("Testing request to  a server", func(t *testing.T) {
		data := "Hello this is your boy"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, req)

		if response.Body.String() != data {
			t.Errorf("Got %v want %v", response.Body.String(), data)
		}

	})

	t.Run("Testing the server if the request stops before 1 second", func(t *testing.T) {
		data := "Hello guys"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelContext, cancel := context.WithCancel(req.Context())

		time.AfterFunc(200*time.Millisecond, cancel)

		req = req.WithContext(cancelContext)
		response := &SpyResponseWriter{}

		server.ServeHTTP(response, req)

		if response.written {
			t.Errorf("Should have not written this bro")
		}

	})

}
