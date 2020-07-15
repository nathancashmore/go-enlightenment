package _select

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (result string) {

	durationA := measureDuration(urlA)
	durationB := measureDuration(urlB)

	if durationA > durationB {
		return urlB
	}

	return urlA
}

func measureDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
