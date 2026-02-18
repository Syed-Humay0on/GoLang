package main

import (
	"context"
	"fmt"
)

func sayHello(ctx context.Context, name string) {
	fmt.Printf("Hello %s! (context mila)\n", name)
}
func main() {
	ctx := context.Background()

	sayHello(ctx, "jaani")
	sayHello(ctx, "yaar")
}
