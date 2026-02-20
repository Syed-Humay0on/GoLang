package main

import (
	"context"
	"fmt"
)

func encrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "Request-id", "12345")
}

func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("Request-id")
	fmt.Println(rID)
}

func main() {
	fmt.Println("Go Context basic understanding")
	ctx := context.Background()
	ctx = encrichContext(ctx)
	doSomethingCool(ctx)
}
