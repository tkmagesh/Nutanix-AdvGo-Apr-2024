package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 10
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 20
	}()

	go func() {
		time.Sleep(2 * time.Second)
		d3 := <-ch3
		fmt.Println("ch3 :", d3)
	}()

	for i := 0; i < 3; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println("ch1 :", d1)
		case d2 := <-ch2:
			fmt.Println("ch2 :", d2)
		case ch3 <- 30:
			fmt.Println("data sent to ch3")
		}
	}
}
