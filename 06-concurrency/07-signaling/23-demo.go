package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	dataCh := genData()
	for data := range dataCh {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func genData() <-chan string {
	wg := &sync.WaitGroup{}
	dataCh := make(chan string)

	fibStopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop fib generation...")
	go func() {
		fmt.Scanln()
		// fibStopCh <- struct{}{}
		close(fibStopCh)
	}()

	primeStopCh := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(primeStopCh)
	}()

	stopCh := make(chan struct{})
	go func() {
		wg.Add(1)
		fibCh := genFib(wg, fibStopCh)
		wg.Add(1)
		primeCh := genPrime(wg, primeStopCh)
		go func() {
		LOOP:
			for {
				select {
				case fibNo, isOpen := <-fibCh:
					if isOpen {
						dataCh <- fmt.Sprintf("fib : %d", fibNo)
					}
				case primeNo, isOpen := <-primeCh:
					if isOpen {
						dataCh <- fmt.Sprintf("prime : %d", primeNo)
					}
				case <-stopCh:
					break LOOP
				}
			}
		}()
		wg.Wait()
		close(stopCh)
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
