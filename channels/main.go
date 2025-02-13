package main

import (
	"fmt"
	"sync"
)

func Merge(intChs ...<-chan int) <-chan int {
	resultCh := make(chan int)

	wg := sync.WaitGroup{}

	for _, ch := range intChs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range ch {
				resultCh <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

func main() {
	// Тест обеих версий
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Тестируем явную версию
	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i <= 12; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	fmt.Println("Тест MergeExplicit:")
	mergedCh := Merge(ch1, ch2)
	for val := range mergedCh {
		fmt.Printf("%d ", val)
	}
	fmt.Println("\nГотово!")
}
