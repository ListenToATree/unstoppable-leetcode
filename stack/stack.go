package stack

import (
	"errors"
	"sync"
)

type Stack[T any] struct {
	mutex sync.Mutex
	arr   []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		arr: []T{},
	}
}

// Push adds to the top
func (s *Stack[T]) Push(i T) {
	s.mutex.Lock()
	s.arr = append(s.arr, i)
	s.mutex.Unlock()
}

// Peek returns the top item but not remove it
func (s *Stack[T]) Peek() (T, error) {
	if len(s.arr) == 0 {
		return getZero[T](), errors.New("no item left")
	}
	s.mutex.Lock()
	item := s.arr[s.Size()-1]
	s.mutex.Unlock()
	return item, nil
}

// Pop returns the top item
func (s *Stack[T]) Pop() (T, error) {
	ret, err := s.Peek()
	if err != nil {
		return getZero[T](), err
	}
	s.mutex.Lock()
	s.arr = s.arr[0 : s.Size()-1]
	s.mutex.Unlock()
	return ret, nil
}

func (s *Stack[T]) Size() int {
	return len(s.arr)
}

func getZero[T any]() T {
	var result T
	return result
}
