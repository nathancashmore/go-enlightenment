package sync

/*
Simple program to provide a counter object that can be incremented

"A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for.
Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished."

*/

type Counter struct {
	total int
}

func (c *Counter) Inc() {
	c.total++
}

func (c *Counter) Value() int {
	return c.total
}
