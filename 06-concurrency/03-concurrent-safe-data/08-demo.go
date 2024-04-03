package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

func (c *Counter) Count() int {
	c.Lock()
	defer c.Unlock()
	return c.count
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", counter.Count())
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}
