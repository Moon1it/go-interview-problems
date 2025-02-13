package main

import "fmt"

type NewType int

func (this *NewType) Reverse(slice []int) {
	end := len(slice) - 1

	for i := range slice {
		if i >= (end - i) {
			return
		}
		slice[i], slice[end-i] = slice[end-i], slice[i]
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	t := NewType(10)
	go t.Reverse(slice)

	fmt.Println(slice)
}
