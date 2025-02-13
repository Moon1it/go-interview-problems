package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 2000)
	for i := range slice {
		slice[i] = i
	}

	target := 1950

	fmt.Println(BinarySearch(slice, target))
}

func BinarySearch(slice []int, target int) int {
	low := 0
	high := len(slice) - 1

	count := 0

	for low <= high {
		count++
		mid := (low + high) / 2

		if slice[mid] == target {
			fmt.Println("count", count)
			return mid
		}

		if slice[mid] < target {
			low = mid + 1
		}

		if slice[mid] > target {
			high = mid - 1
		}
	}

	return -1
}
