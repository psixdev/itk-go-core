package main

import (
	"fmt"
	"math"
)

func convert(src <-chan uint8) <-chan float64 {
	dst := make(chan float64)

	go func() {
		defer close(dst)

		for val := range src {
			dst <- math.Pow(float64(val), 3)
		}
	}()

	return dst
}

func main() {
	src := make(chan uint8)
	dst := convert(src)

	go func() {
		defer close(src)

		for i := uint8(0); i <= 7; i++ {
			src <- i
		}
	}()

	for val := range dst {
		fmt.Println(val)
	}
}
