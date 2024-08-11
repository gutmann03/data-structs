package stack_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"data-structs/stack"
)

func TestStack(t *testing.T) {
	t.Run("Constructors", func(t *testing.T) {
		t.Run("New", func(t *testing.T) {
			assert.NotNil(t, stack.New[int]())
		})

		t.Run("NewWithLimit", func(t *testing.T) {
			assert.PanicsWithValue(t, "not positive limit value", func() { stack.NewWithLimit[int](-1) })
			assert.PanicsWithValue(t, "not positive limit value", func() { stack.NewWithLimit[int](0) })
			assert.NotNil(t, stack.NewWithLimit[int](1))
		})
	})

	t.Run("Getters", func(t *testing.T) {
		var s *stack.Stack[int]
		tests := []struct {
			action  func(*testing.T)
			length  int
			limit   int
			isFull  bool
			isEmpty bool
		}{
			{
				action: func(t *testing.T) {
					s = stack.NewWithLimit[int](2)

					_, err := s.Pop()
					assert.ErrorIs(t, err, stack.ErrEmpty)
					assert.PanicsWithError(t, stack.ErrEmpty.Error(), func() { s.MustPop() })

					_, err = s.Pick()
					assert.ErrorIs(t, err, stack.ErrEmpty)
					assert.PanicsWithError(t, stack.ErrEmpty.Error(), func() { s.MustPick() })
				},
				length:  0,
				limit:   2,
				isFull:  false,
				isEmpty: true,
			},
			{
				action: func(t *testing.T) {
					assert.NoError(t, s.Push(1))

					element, err := s.Pick()
					assert.NoError(t, err)
					assert.EqualValues(t, element, 1)
				},
				length:  1,
				limit:   2,
				isFull:  false,
				isEmpty: false,
			},
			{
				action: func(t *testing.T) {
					assert.NoError(t, s.Push(10))

					element, err := s.Pick()
					assert.NoError(t, err)
					assert.EqualValues(t, element, 10)
				},
				length:  2,
				limit:   2,
				isFull:  true,
				isEmpty: false,
			},
			{
				action: func(t *testing.T) {
					assert.ErrorIs(t, s.Push(1), stack.ErrFull)
				},
				length:  2,
				limit:   2,
				isFull:  true,
				isEmpty: false,
			},
			{
				action: func(t *testing.T) {
					_, err := s.Pop()
					assert.NoError(t, err)

					_, err = s.Pop()
					assert.NoError(t, err)
				},
				length:  0,
				limit:   2,
				isFull:  false,
				isEmpty: true,
			},
			{
				action: func(t *testing.T) {
					_, err := s.Pop()
					assert.ErrorIs(t, err, stack.ErrEmpty)

					_, err = s.Pick()
					assert.ErrorIs(t, err, stack.ErrEmpty)
				},
				length:  0,
				limit:   2,
				isFull:  false,
				isEmpty: true,
			},
		}
		for _, test := range tests {
			test.action(t)
			assert.EqualValues(t, test.length, s.Len())
			assert.EqualValues(t, test.limit, s.Limit())
			assert.EqualValues(t, test.isEmpty, s.IsEmpty())
			assert.EqualValues(t, test.isFull, s.IsFull())
		}
	})

	t.Run("Push & Pop & Pick", func(t *testing.T) {
		s := stack.NewWithLimit[string](3)
		s.MustPush("!")
		assert.EqualValues(t, "!", s.MustPick())

		s.MustPush("world")
		assert.EqualValues(t, "world", s.MustPick())

		s.MustPush("Hello")
		assert.EqualValues(t, "Hello", s.MustPick())

		assert.PanicsWithError(t, stack.ErrFull.Error(), func() { s.MustPush("!") })
		assert.EqualValues(t, 3, s.Len())
		assert.True(t, s.IsFull())

		var content = make([]string, 0, s.Len())
		for !s.IsEmpty() {
			content = append(content, s.MustPop())
		}
		assert.PanicsWithError(t, stack.ErrEmpty.Error(), func() { s.MustPop() })
		assert.EqualValues(t, "Hello world !", strings.Join(content, " "))
	})
}
