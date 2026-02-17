package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, ctx context.Context) {
	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: Jobs khatam!\n", id)
				return
			}
			fmt.Printf("Worker %d → Image %d shuru\n", id, j)
			time.Sleep(300 * time.Millisecond) // kaam
			fmt.Printf("Worker %d → Image %d done\n", id, j)

		case <-ctx.Done():
			fmt.Printf("Worker %d: Context cancel/timeout mila!\n", id)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, ctx)
		}(i)
	}

	// 8 jobs bhejo
	for j := 1; j <= 8; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	fmt.Println("Sab khatam — graceful shutdown!")
}
