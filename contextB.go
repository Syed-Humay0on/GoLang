package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
	if err != nil {
		fmt.Println("Oops:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Bummer:", err) // Times out? You’ll see context.DeadlineExceeded
		return
	}
	defer resp.Body.Close()
	fmt.Println("Sweet:", resp.Status)
}
