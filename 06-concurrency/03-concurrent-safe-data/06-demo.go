package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	count++
}
