package main

import (
	"context"
	"fmt"
	"math/rand/v2"
)

var getRandomInt = rand.IntN

func random(ctx context.Context, max int) <-chan int {
	ch := make(chan int)

	if max <= 0 {
		panic("invalid max value")
	}

	go func() {
		defer close(ch)

		for {
			select {
			case <-ctx.Done():
				return
			case ch <- getRandomInt(max):
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := random(ctx, 10)

	for val := <-ch; val != 0; val = <-ch {
		fmt.Println(val)
	}

	fmt.Println(0)
}
