package stack

import (
	"errors"
)

// defaultLimit defines default limit value for Stack without limit provided.
const defaultLimit = 1000

var (
	// ErrFull indicates that Stack reached its limit and cannot be pushed.
	ErrFull = errors.New("the stack reached the limit")
	// ErrEmpty indicates that Stack reached 0 length and cannot be popped.
	ErrEmpty = errors.New("the stack is empty")
)

// Stack is a generic implementation for basic data structure stack.
// Ordering type: LIFO.
type Stack[T any] struct {
	elements []T
	index    int
	limit    int
}

// New returns pointer to the created Stack structure with default limit value.
func New[T any]() *Stack[T] {
	return NewWithLimit[T](defaultLimit)
}

// NewWithLimit returns pointer to the created Stack structure with provided limit value.
// NOTE: Panics if the limit id <= 0.
func NewWithLimit[T any](limit int) *Stack[T] {
	if limit <= 0 {
		panic("not positive limit value")
	}

	return &Stack[T]{
		elements: make([]T, 0, limit),
		index:    -1,
		limit:    limit,
	}
}

// Push adds element into the Stack.
// Returns ErrFull in case the Stack's limit was reached.
func (s *Stack[T]) Push(element T) error {
	if s.IsFull() {
		return ErrFull
	}

	s.elements = append(s.elements, element)
	s.index++

	return nil
}

// MustPush adds element into the Stack.
// Note: Uses [Stack.Push] and panics in case non-nil error was returned.
func (s *Stack[T]) MustPush(element T) {
	if err := s.Push(element); err != nil {
		panic(err)
	}
}

// Pop removes the last inserted element in the Stack, returning it.
// Returns ErrEmpty in case all elements were popped from the Stack.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), ErrEmpty
	}

	var element = s.elements[s.index]
	s.elements = s.elements[:s.index]
	s.index--

	return element, nil
}

// MustPop removes the last inserted element in the Stack, returning it.
// NOTE: Uses [Stack.Pop] and panics in case non-nil error was returned.
func (s *Stack[T]) MustPop() T {
	if element, err := s.Pop(); err != nil {
		panic(err)
	} else {
		return element
	}
}

// Pick returns the last inserted element in the Stack, without modifying the Stack.
// Returns ErrEmpty in case there is no elements in the Stack.
func (s *Stack[T]) Pick() (T, error) {
	if s.IsEmpty() {
		return *new(T), ErrEmpty
	}

	return s.elements[s.index], nil
}

// MustPick returns the last inserted element in the Stack, without modifying the Stack.
// NOTE: Uses [Stack.Pick] and panics in case non-nil error was returned.
func (s *Stack[T]) MustPick() T {
	if element, err := s.Pick(); err != nil {
		panic(err)
	} else {
		return element
	}
}

// Len returns the number of elements in the Stack.
func (s *Stack[T]) Len() int {
	return s.index + 1
}

// Limit returns the maximum allowed number of elements in the Stack.
func (s *Stack[T]) Limit() int {
	return s.limit
}

// IsEmpty returns true if the Stack has no elements, and false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

// IsFull returns true if the Stack cannot be pushed with any element, and false otherwise.
func (s *Stack[T]) IsFull() bool {
	return s.Len() == s.Limit()
}
