package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
}

func genNos(ch chan int) {
	count := rand.Intn(20)
	fmt.Println("Count :", count)
	for i := 1; i <= count; i++ {
		ch <- i * 10
	}
	close(ch)
}
