package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) GetCount() int {
	return c.count
}

func mainSafe() {
	wg := &sync.WaitGroup{}
	counter := Counter{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			counter.Increment()
			fmt.Println("计数器的值为:", counter.GetCount())

		}()

	}
	// Print
	wg.Wait()

}
