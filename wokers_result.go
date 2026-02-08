package main

import (
	"fmt"
	"sync"
	"time"
)

func myImageWorker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Woker %d: downloading image %d... \n", id, j)
		time.Sleep(time.Millisecond * 500)
		results <- fmt.Sprintf("Image %d done by worker %d", j, id)
	}
}

func main() {
	start := time.Now()

	jobs := make(chan int, 10)
	results := make(chan string, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go myImageWorker(w, jobs, results, &wg)
	}

	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Result:", res)
	}
	fmt.Printf("Total time: %s\n", time.Since(start))
}
