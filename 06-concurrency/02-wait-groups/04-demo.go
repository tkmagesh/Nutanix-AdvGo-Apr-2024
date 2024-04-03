package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	wg := &sync.WaitGroup{}

	flag.IntVar(&count, "count", 0, "# of goroutines to spin")
	flag.Parse()
	fmt.Printf("Starting %d goroutines... hit ENTER to start\n", count)
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1)    // increment the wg counter by 1
		go fn(i, wg) // scheduling the execution of fn through the scheduler
	}
	wg.Wait() // block until the wg counter becomes 0 (default = 0)
	fmt.Println("Done!")
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second) // make the delay random
	fmt.Printf("fn[%d] completed\n", id)
}
