package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	fibCh := genFib(stopCh)

	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()

	for no := range fibCh {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

// using "share memory by communicating" (advisable)

func genFib(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for x, y := 0, 1; ; x, y = y, x+y {
			select {
			case <-stopCh:
				fmt.Println("stop signal received... closing the channel")
				close(ch)
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
			}
		}

	}()
	return ch
}

// using "communicate by sharing memory" (not advisable)
/*
var timedOut bool = false

func elapse(d time.Duration) {
	time.Sleep(d)
	timedOut = true
}

func genFib() <-chan int {
	ch := make(chan int)
	go func() {
		go elapse(5 * time.Second)
		for x, y := 0, 1; ; x, y = y, x+y {
			time.Sleep(500 * time.Millisecond)
			ch <- x
			if timedOut {
				break
			}
		}
		fmt.Println("timeout... closing the channel")
		close(ch)
	}()
	return ch
}
*/
