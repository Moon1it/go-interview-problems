package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, jobCh <-chan int, resultCh chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobCh:
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			case resultCh <- job * job:
			}
		}
	}
}

func main() {
	maxWorkers := 3
	tasks := 10

	jobCh := make(chan int, maxWorkers)
	resultCh := make(chan int, maxWorkers)

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	now := time.Now()

	wg := sync.WaitGroup{}
	for j := 0; j < maxWorkers; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, jobCh, resultCh)
		}()
	}

	go func() {
		defer close(jobCh)
		for i := 0; i < tasks; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				select {
				case <-ctx.Done():
					return
				case jobCh <- i:
				}
			}
		}
	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for val := range resultCh {
		fmt.Println(val)
	}

	fmt.Println("Time taken:", time.Since(now))
}
