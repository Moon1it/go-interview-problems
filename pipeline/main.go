package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context, nums ...int) <-chan int {
	numCh := make(chan int, len(nums))
	go func() {
		defer close(numCh)
		for _, val := range nums {
			select {
			case <-ctx.Done():
				return
			case numCh <- val:
			}
		}
	}()
	return numCh
}

func sq(ctx context.Context, in <-chan int) <-chan int {
	numCh := make(chan int)
	go func() {
		defer close(numCh)
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}

				time.Sleep(2 * time.Second)

				select {
				case <-ctx.Done():
					return
				case numCh <- val * val:
				}
			}
		}
	}()
	return numCh
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for val := range sq(ctx, gen(ctx, 1, 2, 3, 4, 5)) {
		fmt.Println(val)
	}
}
