package main

import (
	"fmt"
)

func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/

	ch := make(chan int)

	go add(ch, 100, 200)
	result := <-ch
	fmt.Println("result :", result)
}

func add(ch chan<- int, x, y int) {
	result := x + y
	ch <- result
}
