package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	ch <- 10
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	ch <- 20
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	ch <- 30
	fmt.Println("chan data : ", <-ch)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	fmt.Println("chan data : ", <-ch)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
}
