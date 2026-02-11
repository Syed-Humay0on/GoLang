package main

import (
	"fmt"
	"time"
)

func main() {
	jobChannel := make(chan string)

	// Background mein heavy task
	go func() {
		time.Sleep(3 * time.Second) // Image download mein 3 sec lagenge
		jobChannel <- "Image Downloaded!"
	}()

	// SELECT: Jo pehle ho jaye!
	select {
	case res := <-jobChannel:
		fmt.Println(res)
	case <-time.After(2 * time.Second): // 2 second ka timer
		fmt.Println("Error: Download took too long! Timeout.")
	}
}

