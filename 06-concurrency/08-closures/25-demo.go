package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("i :", i)
		}(i)
	}
	wg.Wait()
}
