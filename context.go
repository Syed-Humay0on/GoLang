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
				fmt.Printf("Worker %d: jobs khatam, nikal raha hoon\n", id)
				return
			}
			fmt.Printf("Worker %d -> image %d download shuru\n", id, j)
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("Worker %d → image %d done\n", id, j)

		case <-ctx.Done():
			fmt.Printf("Worker %d: Cancel/timeout mila! Band kar raha hoon\n", id)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)

	defer cancel()

	jobs := make(chan int, 20)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, ctx) // ← ctx pass kar diya
		}(i)
	}

	go func() {
		for j := 1; j <= 15; j++ {
			select {
			case jobs <- j:
			case <-ctx.Done():
				fmt.Println("Producer: Timeout ho gaya, jobs band!")
				return
			}
		}
		close(jobs)
	}()

	wg.Wait()
	fmt.Println("Sab workers band ho gaye — graceful shutdown!")
}
