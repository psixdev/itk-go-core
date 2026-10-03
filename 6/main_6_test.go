package main

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func stubGetRandomInt() func(int) int {
	callsCount := 0

	return func(max int) int {
		callsCount++
		return 8 - callsCount
	}
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestRandom(t *testing.T) {
	t.Run("with invalid max", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		assert.Panics(t, func() {
			random(ctx, 0)
		})
	})

	t.Run("with valid max", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		oldFunc := getRandomInt
		getRandomInt = stubGetRandomInt()
		defer func() { getRandomInt = oldFunc }()

		ch := random(ctx, 10)

		nums := make([]int, 0, 5)

		for val := <-ch; val != 2; val = <-ch {
			nums = append(nums, val)
		}

		cancel()

		for range ch {
		}

		assert.Equal(t, []int{7, 6, 5, 4, 3}, nums)
	})
}
