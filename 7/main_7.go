package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func merge(chs ...chan int) <-chan int {
	chRes := make(chan int)

	var wg sync.WaitGroup

	go func() {
		defer close(chRes)

		for _, ch := range chs {
			wg.Go(func() {
				for val := range ch {
					chRes <- val
				}
			})
		}

		wg.Wait()
	}()

	return chRes
}

func main() {
	chs := []chan int{make(chan int), make(chan int), make(chan int)}
	cr := merge(chs...)

	go func() {
		for _, ch := range chs {
			defer close(ch)
		}

		for n := 0; n < 7; n++ {
			i := rand.IntN(3)
			chs[i] <- n
		}
	}()

	for val := range cr {
		fmt.Println(val)
	}
}
