package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	for _, i := range []int{1, 2, 3} {
		go func(i int) {
			ctx2, _ := context.WithCancel(ctx)
			context.Background()
			ctx3, _ := context.WithCancel(ctx2)
			_ = ctx3
			fmt.Println(i)

		}(i)
	}

	time.Sleep(1 * time.Second)
}
