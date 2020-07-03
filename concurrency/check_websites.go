/*
Utility which returns a map of each URL checked to a boolean value - true for a good response, false for a bad response.

*/
package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string, inParallel bool) map[string]bool {
	if inParallel {
		return _checkInParallel(wc, urls)
	} else {
		return _checkInSeries(wc, urls)
	}
}

func _checkInParallel(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}

func _checkInSeries(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
