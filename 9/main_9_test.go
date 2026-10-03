package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.M) {
	code := t.Run()

	os.Exit(code)
}

func TestConvert(t *testing.T) {
	src := make(chan uint8)
	dst := convert(src)

	go func() {
		defer close(src)

		for i := uint8(0); i <= 7; i++ {
			src <- i
		}
	}()

	nums := make([]float64, 0, 8)

	for val := range dst {
		nums = append(nums, val)
	}

	assert.Equal(t, []float64{0, 1, 8, 27, 64, 125, 216, 343}, nums)
}
