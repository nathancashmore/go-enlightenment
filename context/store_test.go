package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func TestHandler(t *testing.T) {
	data := "Hello World"
	srv := Server(&SpyStore{response: data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	srv.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf("got %s wanted %s", response.Body.String(), data)
	}
}
