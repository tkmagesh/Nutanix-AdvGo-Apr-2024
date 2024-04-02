package main

import "fmt"

func main() {

	// v2.0

	// var add, subtract func(int, int)
	add := getLogOperation(add)
	subtract := getLogOperation(subtract)

	// v1.0

	add(100, 200)
	subtract(100, 200)

}

// v2.0

func getLogOperation(operationFn func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("operation started")
		operationFn(x, y)
		fmt.Println("operation completed")
	}
}

/*
func logAdd(x, y int) {
	fmt.Println("operation started")
	add(x, y)
	fmt.Println("operation completed")
}

func logSubtract(x, y int) {
	fmt.Println("operation started")
	subtract(x, y)
	fmt.Println("operation completed")
}
*/

// v1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
