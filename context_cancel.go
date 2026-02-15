package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done(): // ← yahan signal aaya!
			fmt.Printf("Worker %d: Cancel mila! Band ho raha hoon.\n", id)
			return

		default:
			fmt.Printf("Worker %d: Kaam chal raha hai...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Worker shuru karo
	go worker(ctx, 1)

	time.Sleep(2 * time.Second) // 2 second baad cancel karo
	fmt.Println("Main: Ab cancel kar raha hoon!")
	cancel() // ← yahan whistle bajti hai

	time.Sleep(1 * time.Second) // thoda wait karo dekhne ke liye
}
