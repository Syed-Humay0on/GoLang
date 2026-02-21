package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Yo! Wassup! This is a goRoutine basic code\nCheck this out")
	// GoRoutines are lightweight threads managed by runtime
	// Enable us to create asynchronous parrallel threaded programs
	compute(5)
	compute(6)
	fmt.Println("code executed synchronously/sequentially without goroutines")
	go compute(5)
	go compute(6)
	fmt.Scanln()
}
