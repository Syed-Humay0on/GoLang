package main

import (
	"context"
	"fmt"
	_ "net/http"
	"time"
)

func task(ctx context.Context, id int) {
	select {
	case <-time.After(time.Duration(id) * time.Second):
		fmt.Printf("Task %d done!\n", id)
	case <-ctx.Done():
		fmt.Printf("Task %d bailed: %v\n", id, ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 1; i <= 3; i++ {
		go task(ctx, i)
	}
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
