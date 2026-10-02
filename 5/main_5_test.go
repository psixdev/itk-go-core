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

func TestIntersection(t *testing.T) {
	t.Run("both are empty", func(t *testing.T) {
		has, inter := intersection([]int{}, []int{})

		assert.False(t, has)
		assert.Empty(t, inter)
	})

	t.Run("first is empty", func(t *testing.T) {
		has, inter := intersection([]int{}, []int{1, 2})

		assert.False(t, has)
		assert.Empty(t, inter)
	})

	t.Run("second is empty", func(t *testing.T) {
		slice := []int{1, 2}

		has, inter := intersection(slice, []int{})

		assert.False(t, has)
		assert.Empty(t, inter)
	})

	t.Run("both aren't empty", func(t *testing.T) {
		slice := []int{1, 2, 3, 2, 3, 4, 5, 5}

		has, inter := intersection(slice, []int{2, 4, 4, 5, 5, 7})

		assert.True(t, has)
		assert.Equal(t, []int{2, 2, 4, 5, 5}, inter)
		assert.NotSame(t, &slice[0], &inter[0])
	})
}
