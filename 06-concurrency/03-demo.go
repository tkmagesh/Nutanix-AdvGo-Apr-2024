package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// runtime.GOMAXPROCS(1)
	for i := 0; i < 100; i++ {
		wg.Add(1) // increment the wg counter by 1
		go f1()   // scheduling the execution of f1 through the scheduler
	}
	f2()
	wg.Wait() // block until the wg counter becomes 0 (default = 0)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second) // make the delay random
	fmt.Println("f1 completed")
	wg.Done() // decrement the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
