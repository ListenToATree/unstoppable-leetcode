package stack

import (
	"testing"
)

func TestStack_error(t *testing.T) {
	stack := NewStack[int]()
	_, err := stack.Pop()
	if err != nil {
		println(err.Error())
	}
}

func TestStack_happy(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	item, err := stack.Pop()
	if err != nil {
		println(err.Error())
	}
	println(item)
}
