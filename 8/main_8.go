package main

import (
	"fmt"
	"sync"
	"time"
)

type WaitGroup struct {
	count int
	mu    sync.Mutex
	ch    chan struct{}
}

func (wg *WaitGroup) Add(count int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.count += count

	if wg.count < 0 {
		wg.count = 0
	}

	if wg.count > 0 {
		if wg.ch == nil {
			wg.ch = make(chan struct{})
		}
	} else if wg.ch != nil {
		close(wg.ch)
		wg.ch = nil
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	wg.mu.Lock()
	ch := wg.ch
	wg.mu.Unlock()

	if ch != nil {
		<-ch
	}
}

func main() {
	var wg WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Printf("run job %d\n", i)

			time.Sleep(time.Second)

			fmt.Printf("job %d is done\n", i)
		}()
	}

	wg.Wait()
}
