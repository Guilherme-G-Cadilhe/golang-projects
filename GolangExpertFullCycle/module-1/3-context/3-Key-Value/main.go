package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha1234")

	bookHotel(ctx, "Bob")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(name, token)
}
