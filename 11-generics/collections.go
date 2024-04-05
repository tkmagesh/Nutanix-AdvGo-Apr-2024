package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Category string
	Units    int
}

func (p Product) GetId() int {
	return p.Id
}

type Products []Product

/*
func (products Products) Print() {
	for _, p := range products {
		fmt.Println(p)
	}
}
*/

func Print[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

/*
func GetById(products Products, id int) (p Product, err error) {
	for _, product := range products {
		if id == product.Id {
			p = product
			return
		}
	}
	err = errors.New("item not found")
	return
}
*/

/*
type IdType interface {
	GetId() int
}

func GetById[T IdType](items []T, id int) (p T, err error) {
	for _, item := range items {
		if id == item.GetId() {
			p = item
			return
		}
	}
	err = errors.New("item not found")
	return
}
*/

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128
}

type IdType interface {
	GetId() int
}

type CollectionWithId[T IdType, K Numbers] []T

func (items CollectionWithId[T, K]) GetById(id int) (p T, err error) {
	for _, item := range items {
		if id == item.GetId() {
			p = item
			return
		}
	}
	err = errors.New("item not found")
	return
}

func (items CollectionWithId[T, K]) filter(predicate func(T) bool) CollectionWithId[T, K] {
	var result CollectionWithId[T, K]
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func (items CollectionWithId[T, K]) sum(valueSelector func(T) K) K {
	var result K
	for _, item := range items {
		result += valueSelector(item)
	}
	return result
}

var products = CollectionWithId[Product, int]{
	Product{100, "Pen", 10, "Stationary", 10},
	Product{101, "Pencil", 5, "Stationary", 5},
	Product{102, "Marker", 50, "Utencil", 20},
}

var numbers = []int{10, 20, 30}

func main() {
	// Print(products)
	Print(numbers)
	// p, err := GetById(products, 101)
	p, err := products.GetById(101)
	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(p)
	}

	stationaryProducts := products.filter(func(p Product) bool {
		return p.Category == "Stationary"
	})
	fmt.Println(stationaryProducts)

	totalUnits := products.sum(func(p Product) int {
		return p.Units
	})
	fmt.Println(totalUnits)

}
