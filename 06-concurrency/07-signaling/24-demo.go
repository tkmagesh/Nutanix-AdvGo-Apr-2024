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
	fmt.Println("Hit ENTER to stop generating fib series....")
	go func() {
		fmt.Scanln()
		close(fibStopCh)
	}()

	primeStopCh := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(primeStopCh)
	}()

	go func() {
		fibCh := genFib(fibStopCh)
		primeCh := genPrime(primeStopCh)

		wg.Add(1)
		go printFib(wg, fibCh, dataCh)

		wg.Add(1)
		go printPrime(wg, primeCh, dataCh)

		wg.Wait()
		close(dataCh)
	}()

	return dataCh

}

func printFib(wg *sync.WaitGroup, fibCh <-chan int, dataCh chan<- string) {
	defer wg.Done()
	for fibNo := range fibCh {
		dataCh <- fmt.Sprintf("fib : %d", fibNo)
	}
}

func printPrime(wg *sync.WaitGroup, PrimeCh <-chan int, dataCh chan<- string) {
	defer wg.Done()
	for primeNo := range PrimeCh {
		dataCh <- fmt.Sprintf("prime : %d", primeNo)
	}
}

// using "share memory by communicating" (advisable)
func genPrime(stopCh <-chan struct{}) <-chan int {
	primeCh := make(chan int)
	go func() {
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

func genFib(stopCh <-chan struct{}) <-chan int {
	fibCh := make(chan int)
	go func() {
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
