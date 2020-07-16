package _select

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func ExampleRacer() {
	slowServer := "http://www.facebook.com"
	fastServer := "http://really.fast.server"

	fmt.Println(Racer(fastServer, slowServer))
	// Output: http://really.fast.server
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Racer("http://url.a.com", "http://url.b.com")
	}
}

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(1 * time.Millisecond)

	comparisonUrlSet := []struct {
		name     string
		urlA     string
		urlB     string
		expected string
	}{
		{"fastest second input",
			slowServer.URL, fastServer.URL, fastServer.URL},
		{"fastest first input",
			fastServer.URL, slowServer.URL, fastServer.URL},
	}

	for _, comparison := range comparisonUrlSet {
		t.Run(comparison.name, func(t *testing.T) {

			want := comparison.expected
			got := Racer(comparison.urlA, comparison.urlB)

			if got != want {
				t.Errorf("got %q, wanted %q", got, want)
			}
		})
	}
}

func makeDelayedServer(duration time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
