package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	dataCh := genData(stopCh)

	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()

	for data := range dataCh {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func genData(stopCh <-chan struct{}) <-chan string {
	wg := &sync.WaitGroup{}
	dataCh := make(chan string)
	go func() {
		wg.Add(1)
		fibCh := genFib(wg, stopCh)
		wg.Add(1)
		primeCh := genPrime(wg, stopCh)
		go func() {
		LOOP:
			for {
				select {
				case fibNo := <-fibCh:
					dataCh <- fmt.Sprintf("fib : %d", fibNo)
				case primeNo := <-primeCh:
					dataCh <- fmt.Sprintf("prime : %d", primeNo)
				case <-stopCh:
					break LOOP
				}
			}
		}()
		wg.Wait()
		close(dataCh)
	}()
	return dataCh
}

// using "share memory by communicating" (advisable)
func genPrime(wg *sync.WaitGroup, stopCh <-chan struct{}) <-chan int {
	primeCh := make(chan int)
	go func() {
		defer wg.Done()
	LOOP:
		for no := 2; ; no++ {
			select {
			case <-stopCh:
				fmt.Println("stop signal received... closing the prime channel")
				close(primeCh)
				break LOOP
			default:
				if !isPrime(no) {
					continue LOOP
				}
				primeCh <- no
				time.Sleep(500 * time.Millisecond)
			}
		}

	}()
	return primeCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func genFib(wg *sync.WaitGroup, stopCh <-chan struct{}) <-chan int {
	fibCh := make(chan int)
	go func() {
		defer wg.Done()
	LOOP:
		for x, y := 0, 1; ; x, y = y, x+y {
			select {
			case <-stopCh:
				fmt.Println("stop signal received... closing the fib channel")
				close(fibCh)
				break LOOP
			case fibCh <- x:
				time.Sleep(500 * time.Millisecond)
			}
		}

	}()
	return fibCh
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
