package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/

	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, ch, 100, 200)
	wg.Wait()
	result := <-ch
	fmt.Println("result :", result)
}

func add(wg *sync.WaitGroup, ch chan int, x, y int) {
	result := x + y
	ch <- result
	wg.Done()
}
