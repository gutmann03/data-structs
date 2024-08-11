package stack

import (
	"errors"
)

const defaultLimit = 1000

var (
	ErrFull  = errors.New("the stack reached the limit")
	ErrEmpty = errors.New("the stack is empty")
)

type Stack[T any] struct {
	elements []T
	index    int
	limit    int
}

func New[T any]() *Stack[T] {
	return NewWithLimit[T](defaultLimit)
}

func NewWithLimit[T any](limit int) *Stack[T] {
	if limit <= 0 {
		panic("not positive limit value")
	}

	return &Stack[T]{
		elements: make([]T, 0, limit),
		index:    0,
		limit:    limit,
	}
}

func (s *Stack[T]) Push(element T) error {
	if s.limit == s.index {
		return ErrFull
	}

	s.elements = append(s.elements, element)
	s.index++

	return nil
}

func (s *Stack[T]) MustPush(element T) {
	if err := s.Push(element); err != nil {
		panic(err)
	}
}

func (s *Stack[T]) Pop() (T, error) {
	if s.index == 0 {
		return *new(T), ErrEmpty
	}

	s.index--
	var element = s.elements[s.index]
	s.elements = s.elements[:s.index]

	return element, nil
}

func (s *Stack[T]) MustPop() T {
	if element, err := s.Pop(); err != nil {
		panic(err)
	} else {
		return element
	}
}

func (s *Stack[T]) Len() int {
	return s.index
}

func (s *Stack[T]) Limit() int {
	return s.limit
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) IsFull() bool {
	return s.Len() == s.Limit()
}
