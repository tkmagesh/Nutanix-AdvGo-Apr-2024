package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("attempt to divide by zero")

func main() {
	var divisor int

	for {
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		if q, r, err := divideWrapper(100, divisor); err == nil {
			fmt.Println(q, r)
			break
		}
		fmt.Println("Invalid divisor.. try again!")
	}

}

func divideWrapper(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient, remainder = x/y, x%y
	return
}
