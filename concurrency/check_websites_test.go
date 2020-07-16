package concurrency

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://website.c.com" {
		return false
	}
	return true
}

func slowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func ExampleWebsiteChecker() {
	websites := []string{
		"http://website.a.com",
		"http://website.b.com",
		"http://website.c.com",
		"http://website.d.com",
	}

	fmt.Println(CheckWebsites(mockWebsiteChecker, websites, true))
	// Output: map[http://website.a.com:true http://website.b.com:true http://website.c.com:false http://website.d.com:true]
}

func BenchmarkWebsiteChecker(b *testing.B) {
	benchmarks := []struct {
		name       string
		inParallel bool
	}{
		{"Parallel", true},
		{"Series", false},
	}

	urls := make([]string, 100)
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {

			for i := 0; i < len(urls); i++ {
				urls[i] = "http://a.url.of.sorts"
			}

			for i := 0; i < b.N; i++ {
				CheckWebsites(slowWebsiteChecker, urls, bm.inParallel)
			}
		})
	}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://website.a.com",
		"http://website.b.com",
		"http://website.c.com",
		"http://website.d.com",
	}

	want := map[string]bool{
		"http://website.a.com": true,
		"http://website.b.com": true,
		"http://website.c.com": false,
		"http://website.d.com": true,
	}

	t.Run("Check in parallel", func(t *testing.T) {
		got := CheckWebsites(mockWebsiteChecker, websites, true)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})

	t.Run("Check in series", func(t *testing.T) {
		got := CheckWebsites(mockWebsiteChecker, websites, false)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})

}
