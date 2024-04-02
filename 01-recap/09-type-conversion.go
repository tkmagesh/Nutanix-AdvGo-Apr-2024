package main

import "fmt"

func main() {
	var f float64
	var i int8
	i = 100
	f = float64(i) // use the type name like a function for type conversion
	x := float64(i) * f
	fmt.Printf("type of x : %T\n", x)
	fmt.Println(f, i)
}
