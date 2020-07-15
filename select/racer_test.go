package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Fastest first", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(1 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL

		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("Fastest second", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(1 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL

		got := Racer(fastURL, slowURL)

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func makeDelayedServer(duration time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
