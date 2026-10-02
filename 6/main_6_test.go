package main

import (
	"context"
	"os"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func stubGetRandomInt(callsCount *atomic.Int64) func(int) int {
	callsCount.Store(0)

	return func(max int) int {
		return 8 - int(callsCount.Add(1))
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

		var callsCount atomic.Int64

		oldFunc := getRandomInt
		getRandomInt = stubGetRandomInt(&callsCount)
		defer func() { getRandomInt = oldFunc }()

		ch := random(ctx, 10)

		nums := make([]int, 0, 5)

		for val := <-ch; val != 2; val = <-ch {
			nums = append(nums, val)
		}

		cancel()

		for range ch {
			callsCount.Add(-1)
		}

		assert.Equal(t, 7, int(callsCount.Load()))
		assert.Equal(t, []int{7, 6, 5, 4, 3}, nums)
	})
}
