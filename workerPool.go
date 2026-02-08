package main

import (
	"fmt"
	"sync"
	"time"
)

func imageWorker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d downloading image %d\n", id, j)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	start := time.Now()

	jobs := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go imageWorker(i, jobs, &wg)
	}
	for j := 1; j <= 20; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	fmt.Printf("20 images processed in: %s\n", time.Since(start))
}
