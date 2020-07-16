/*
Module to determine the fastest returning URL based on an input of two URLs.
*/
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
	_, _ = http.Get(url)
	return time.Since(start)
}
