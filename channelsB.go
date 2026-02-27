package main

import "fmt"

func main() {
	// 1. Create a "pipe" (channel) for integers
	dataChan := make(chan int)

	// 2. Start a Producer (Goroutine)
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println("Producer: sending", i)
			dataChan <- i // Send data into channel
		}
		close(dataChan) // Signal that no more data is coming
	}()

	// 3. Consumer (Main Goroutine)
	// 'range' automatically waits for data and stops when channel closes
	for msg := range dataChan {
		fmt.Println("Consumer: received", msg)
	}
}
