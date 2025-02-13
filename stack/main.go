package main

import "fmt"

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{
		nums: []int{},
	}
}

func (this *Stack) Push(elem int) {
	this.nums = append(this.nums, elem)
}

func (this *Stack) Pop() {
	if len(this.nums) < 1 {
		return
	}
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *Stack) Print() {
	for i := len(this.nums) - 1; i >= 0; i-- {
		fmt.Println(this.nums[i])
	}
}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Pop()
	stack.Pop()

	stack.Push(4)

	stack.Pop()
	stack.Pop()

	stack.Print()
}
