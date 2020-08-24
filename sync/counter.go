package sync

import "sync"

/*
Simple program to provide a counter object that can be incremented

"A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for.
Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished."

"A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex."

*/

type Counter struct {
	mu    sync.Mutex
	total int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.total++
}

func (c *Counter) Value() int {
	return c.total
}
