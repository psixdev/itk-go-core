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

func TestMerge(t *testing.T) {
	t.Run("without data", func(t *testing.T) {
		chs := []chan int{make(chan int), make(chan int), make(chan int)}
		cr := merge(chs...)

		go func() {
			for _, ch := range chs {
				close(ch)
			}
		}()

		nums := make([]int, 0)

		for val := range cr {
			nums = append(nums, val)
		}

		assert.Empty(t, nums)
	})

	t.Run("with data", func(t *testing.T) {
		chs := []chan int{make(chan int), make(chan int), make(chan int)}
		cr := merge(chs...)

		go func() {
			for _, ch := range chs {
				defer close(ch)
			}

			chs[0] <- 0
			chs[0] <- 0
			chs[1] <- 1
			chs[2] <- 2
			chs[0] <- 0
			chs[2] <- 2
		}()

		nums := make([]int, 0, 6)

		for val := range cr {
			nums = append(nums, val)
		}

		assert.ElementsMatch(t, []int{0, 0, 0, 1, 2, 2}, nums)
	})
}
