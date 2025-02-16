package main

import (
	"fmt"
	"sync"
)

// До версии Go 1.22 переменная i объявляется один раз на весь цикл,
// и каждая горутина захватывает общую переменную i. К моменту
// выполнения горутин цикл уже завершится, и i примет значение 5.
// Вывод: 5 5 5 5 5 (порядок может быть любым).

// Начиная с Go 1.22, даже в классическом цикле каждая итерация
// имеет свою копию i. Это означает, что горутины будут работать
// с копиями i.
// Вывод: 0 1 2 3 4 (порядок может быть любым).

func Example1() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
}
