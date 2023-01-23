package main

import (
	"context"
	"fmt"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "api-key", "super-secret-api-key")
}

func getApiKey(ctx context.Context) {
	apiKey := ctx.Value("api-key")
	fmt.Println(apiKey)
}

func main() {
	ctx := context.Background()
	ctx = enrichContext(ctx)

	getApiKey(ctx)
}
