package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *safeCounter) increment(wg *sync.WaitGroup) {
	defer wg.Done()

	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
func main() {
	var wg sync.WaitGroup
	counter := safeCounter{count: 0}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter.increment(&wg)
	}
	wg.Wait()
	fmt.Printf("Final count: %d\n", counter.count)
}
