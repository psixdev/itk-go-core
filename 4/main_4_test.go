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

func TestDifference(t *testing.T) {
	t.Run("both are empty", func(t *testing.T) {
		diff := difference([]string{}, []string{})

		assert.Empty(t, diff)
	})

	t.Run("first is empty", func(t *testing.T) {
		diff := difference([]string{}, []string{"foo", "bar"})

		assert.Empty(t, diff)
	})

	t.Run("second is empty", func(t *testing.T) {
		slice := []string{"foo", "bar"}

		diff := difference(slice, []string{})

		assert.Equal(t, []string{"foo", "bar"}, diff)
		assert.NotSame(t, &slice[0], &diff[0])
	})

	t.Run("both aren't empty", func(t *testing.T) {
		slice := []string{"foo", "bar", "foo"}

		diff := difference(slice, []string{"bar", "baz", "bar"})

		assert.Equal(t, []string{"foo", "foo"}, diff)
		assert.NotSame(t, &slice[0], &diff[0])
	})
}
