package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://website.c.com" {
		return false
	}
	return true
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

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}

}
