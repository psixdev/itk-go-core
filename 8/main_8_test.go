package main

import (
	"os"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestWaitGroup(t *testing.T) {
	var wg WaitGroup

	for range 2 {
		var count atomic.Int64
		count.Store(0)

		ch := make(chan struct{})

		for range 3 {
			wg.Add(1)

			go func() {
				defer wg.Done()
				<-ch
				count.Add(1)
			}()
		}

		assert.EqualValues(t, 0, count.Load())

		close(ch)
		wg.Wait()

		assert.EqualValues(t, 3, count.Load())
	}
}
