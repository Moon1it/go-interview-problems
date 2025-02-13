package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testData := []struct {
		slice  []int
		target int
		result int
	}{
		{
			slice:  []int{1},
			target: 1,
			result: 0,
		},
		{
			slice:  []int{1, 4, 5, 13},
			target: 13,
			result: 3,
		},
		{
			slice:  []int{},
			target: 0,
			result: -1,
		},
		{
			slice:  []int{1, 2, 5, 10, 32, 54, 60, 89},
			target: 2,
			result: 1,
		},
	}

	for _, test := range testData {
		r := BinarySearch(test.slice, test.target)
		t.Log(test.result, r)
		assert.Equal(t, test.result, r)
	}
}
