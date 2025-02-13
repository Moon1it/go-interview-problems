package main

import "fmt"

func sum(nums ...int64) *int64 {
	var sum int64
	for _, num := range nums {
		sum += num
	}
	return &sum
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result, "\n", *result)
}
