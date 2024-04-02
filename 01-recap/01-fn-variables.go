package main

import "fmt"

func main() {
	var fn func()
	if fn == nil {
		fmt.Println("fn is not a function")
	}

	fn = func() {
		fmt.Println("f1 invoked")
	}
	fn()

	fn = func() {
		fmt.Println("f2 invoked")
	}
	fn()
}
