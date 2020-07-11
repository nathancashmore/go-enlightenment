package _select

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://not.there.at.all.com"

	want := fastURL

	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
