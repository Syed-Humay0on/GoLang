package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // ← yeh line bohot zaroori

	defer fmt.Printf("Worker %d: cleanup done, bye\n", id)

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: jobs khatam → nikal raha\n", id)
				return
			}
			fmt.Printf("Worker %d → job %d shuru\n", id, job)
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("Worker %d → job %d khatam\n", id, job)

		case <-ctx.Done():
			fmt.Printf("Worker %d: cancel/timeout mila → %v\n", id, ctx.Err())
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // ← context cleanup

	var wg sync.WaitGroup // ← yahan sync package se aaya

	jobs := make(chan int, 10)

	// 4 workers shuru
	for i := 1; i <= 4; i++ {
		wg.Add(1)                    // ← har worker ke liye +1
		go worker(ctx, i, jobs, &wg) // ← pointer pass kar rahe
	}

	// kuch jobs daal do
	for j := 1; j <= 12; j++ {
		jobs <- j
	}
	close(jobs) // ← jobs khatam bata diya

	// Ab wait karo sab workers khatam hone ka
	wg.Wait() // ← yahan rukega jab tak sab Done() na call karen

	fmt.Println("Sab workers band ho gaye → graceful shutdown!")
}
