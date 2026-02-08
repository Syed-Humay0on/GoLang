package main

import (
	"fmt"
	"time"
)

func downloadImage(id int) {
	fmt.Printf("downloading image %d... \n", id)
	time.Sleep(time.Second)
}

func main() {
	start := time.Now()
	for i := 1; i <= 3; i++ {
		downloadImage(i)
	}
	fmt.Printf("Total time taken %s\n", time.Since(start))
}
