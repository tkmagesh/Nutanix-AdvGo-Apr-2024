package main

import (
	"fmt"
	"time"
)

func main() {
	fibCh := genFib()
	for no := range fibCh {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

// using "share memory by communicating" (advisable)

func elapse(d time.Duration) <-chan time.Time {
	timeOut := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeOut <- time.Now()
	}()
	return timeOut
}

func genFib() <-chan int {
	ch := make(chan int)
	go func() {
		timeOut := elapse(5 * time.Second)
	LOOP:
		for x, y := 0, 1; ; x, y = y, x+y {
			select {
			case <-timeOut:
				fmt.Println("timeout... closing the channel")
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
