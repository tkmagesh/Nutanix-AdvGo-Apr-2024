package main

import "fmt"

func main() {
	/*
		fmt.Println(sumInts([]int{1, 2, 3, 4}))
		fmt.Println(sumFloats([]float32{1.5, 2.3, 3.4, 4.7}))
	*/
	fmt.Println(sum([]int{1, 2, 3, 4}))
	fmt.Println(sum([]float32{1.5, 2.3, 3.4, 4.7}))
}

/*
func sum[T int | float32](nos []T) T {
	var result T
	for _, no := range nos {
		result += no
	}
	return result
}
*/

/*
type Ints interface {
	int | int8 | int16 | int32 | int64
}

type Floats interface {
	float32 | float64
}

func sum[T Ints | Floats](nos []T) T {
	var result T
	for _, no := range nos {
		result += no
	}
	return result
}
*/

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128
}

func sum[T Numbers](nos []T) T {
	var result T
	for _, no := range nos {
		result += no
	}
	return result
}

func sumInts(nos []int) int {
	var result int
	for _, no := range nos {
		result += no
	}
	return result
}

func sumFloats(nos []float32) float32 {
	var result float32
	for _, no := range nos {
		result += no
	}
	return result
}
