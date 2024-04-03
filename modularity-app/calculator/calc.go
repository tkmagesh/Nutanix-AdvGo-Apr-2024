package calculator

import "fmt"

var opCount map[string]int

func init() {
	fmt.Println("calculator package initialized - [calc.go]")
	opCount = make(map[string]int)
}

func OpCount() map[string]int {
	return opCount
}
