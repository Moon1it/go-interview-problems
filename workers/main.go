package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, jobCh <-chan int, resultCh chan<- int) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	case job, ok := <-jobCh:
		if !ok {
			return
		}
		resultCh <- job * 2
	}
}

func main() {
	maxWorkers := 5
	tasks := 10

	jobCh := make(chan int, maxWorkers)
	resultCh := make(chan int, maxWorkers)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	for j := 0; j < tasks; j++ {
		go worker(ctx, wg, jobCh, resultCh)
	}

	for i := range tasks {
		jobCh <- i
	}

	for val := range resultCh {
		fmt.Println(val)
	}
}
