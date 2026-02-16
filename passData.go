package main

import (
	"context"
	"fmt"
)

type key string

const userIDKey key = "userID"

func processRequest(ctx context.Context) {
	userID := ctx.Value(userIDKey)
	fmt.Printf("Processing request for user: %v\n", userID)
}

func main() {
	ctx := context.WithValue(context.Background(), userIDKey, 42069)

	processRequest(ctx)
}
