package main

import (
	"context"
	"fmt"
	"time"
)

func level3(ctx context.Context) { // sabse neeche wala function
	fmt.Println("Level 3 shuru")

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Level 3: 5 second pura kaam khatam")
	case <-ctx.Done():
		fmt.Printf("Level 3: ruk gaya → %v\n", ctx.Err())
	}
}

func level2(ctx context.Context) { // middle layer
	fmt.Println("Level 2 shuru")
	level3(ctx) // ← same ctx pass kar diya (chain continue)
	fmt.Println("Level 2 khatam")
}

func level1(ctx context.Context) { // top layer
	fmt.Println("Level 1 shuru")
	level2(ctx) // ← same ctx pass
	fmt.Println("Level 1 khatam")
}

func main() {
	root := context.Background()

	// Middle mein timeout daal rahe hain (real chain banaya)
	ctx, cancel := context.WithTimeout(root, 1700*time.Millisecond)
	defer cancel()

	go level1(ctx) // ← sirf root se shuru kiya

	time.Sleep(3 * time.Second)
	fmt.Println("Main khatam")
}
