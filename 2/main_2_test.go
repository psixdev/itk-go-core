package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestCreateRandomSlice(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := createRandomSlice(0, 10)

		assert.Empty(t, slice)
	})

	t.Run("30 elements", func(t *testing.T) {
		callCount := 0

		oldFunc := getRandomInt
		getRandomInt = func(max int) int {
			callCount++
			return max - 1
		}
		defer func() { getRandomInt = oldFunc }()

		const length = 30
		const max = 5

		slice := createRandomSlice(length, max)

		assert.Len(t, slice, length)
		assert.Equal(t, length, callCount)
		assert.Condition(t, func() bool {
			for _, v := range slice {
				if v != max-1 {
					return false
				}
			}
			return true
		}, "Elements should be equal %d", max-1)
	})
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{name: "empty slice", slice: []int{}, expected: []int(nil)},
		{name: "odd number slice", slice: []int{1, 3, 5, 7}, expected: []int(nil)},
		{
			name:     "non-empty slice",
			slice:    []int{2, 0, 1, 1, 2, 3, 4, 5, 4},
			expected: []int{2, 0, 2, 4, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rSlice := sliceExample(tt.slice)

			assert.Equal(t, tt.expected, rSlice)
		})
	}
}

func TestAddElement(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		rSlice := addElement(slice, 3)

		assert.Equal(t, []int{3}, rSlice)
	})

	t.Run("non-empty slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		rSlice := addElement(slice, 4)

		assert.Equal(t, []int{1, 2, 3, 4}, rSlice)
		assert.NotSame(t, &slice[0], &rSlice[0])
	})
}

func TestCopySlice(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		rSlice := copySlice(slice)

		assert.Equal(t, slice, rSlice)
	})

	t.Run("non-empty slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		rSlice := copySlice(slice)

		assert.Equal(t, slice, rSlice)
		assert.NotSame(t, &slice[0], &rSlice[0])
	})
}

func TestRemoveElement(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		rSlice := removeElement(slice, 2)

		assert.Equal(t, []int(nil), rSlice)
	})

	t.Run("non-empty slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		rSlice := removeElement(slice, 2)

		assert.Equal(t, []int{1, 2}, rSlice)
		assert.NotSame(t, &slice[0], &rSlice[0])
	})
}
