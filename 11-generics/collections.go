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

type IdType interface {
	GetId() int
}

type CollectionWithId[T IdType] []T

func (items CollectionWithId[T]) GetById(id int) (p T, err error) {
	for _, item := range items {
		if id == item.GetId() {
			p = item
			return
		}
	}
	err = errors.New("item not found")
	return
}

func (items CollectionWithId[T]) filter(predicate func(T) bool) CollectionWithId[T] {
	var result CollectionWithId[T]
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

var products = CollectionWithId[Product]{
	Product{100, "Pen", 10, "Stationary"},
	Product{101, "Pencil", 5, "Stationary"},
	Product{102, "Marker", 50, "Utencil"},
}

var numbers = []int{10, 20, 30}

func main() {
	Print(products)
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
}
