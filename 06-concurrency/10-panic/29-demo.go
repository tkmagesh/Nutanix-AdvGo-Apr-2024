package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		fmt.Println("[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println("[main] Something went wrong :", err)
			return
		}
		fmt.Println("[main] Thank you!")
	}()
	/*
		ch, errCh := divide(100, 0)
		select {
		case result := <-ch:
			fmt.Println("divide result :", result)
		case err := <-errCh:
			fmt.Println("Error dividing :", err)
		}
	*/

	// choosing to ignore the error
	ch, _ := divide(100, 0)
	result := <-ch
	fmt.Println("divide result :", result)
}

func divide(x, y int) (<-chan int, <-chan error) {
	ch := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		defer func() {
			fmt.Println("[divide.func] - deferred")
			if err := recover(); err != nil {
				fmt.Println("[divide.func] - sending error")
				errCh <- err.(error)
			}
		}()
		time.Sleep(2 * time.Second)
		ch <- x / y
	}()
	return ch, errCh
}
