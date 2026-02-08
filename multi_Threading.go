package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadingImage(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker started downloading image %d \n", id)
	time.Sleep(time.Second)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go downloadingImage(i, &wg)
	}
	wg.Wait()
	fmt.Printf("Total time %s\n", time.Since(start))
}
