package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
Write the apis for the following

IndexOf => return the index of the given product (return an error if not exists )
	ex:  returns the index of the given product

Includes => return true if the given product is present in the collection else return false
	ex:  returns true if the given product is present in the collection

Filter => return a new collection of products that satisfy the given condition
	some of the use cases:
		1. filter all costly products (cost > 1000)
			OR
		2. filter all stationary products (category = "Stationary")
			OR
		3. filter all costly stationary products
		etc

Any => return true if any of the product in the collections satifies the given criteria
	some of the use cases:
		1. are there any costly product (cost > 1000)?
		OR
		2. are there any stationary product (category = "Stationary")?
		OR
		3. are there any costly stationary product?
		etc

All => return true if all the products in the collections satifies the given criteria
	some of use cases:
		1. are all the products costly products (cost > 1000)?
		OR
		2. are all the products stationary products (category = "Stationary")?
		OR
		3. are all the products costly stationary products?
		etc

Sort => to sort the products by any attribute
	IMPORTANT: use sort.Sort() function

*/

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

var ErrProductNotFound error = errors.New("product not found")

type Products []Product

func (products Products) String() string {
	var sb strings.Builder
	for _, p := range products {
		sb.WriteString(fmt.Sprintf("%s\n", p))
	}
	return sb.String()
}

func (products Products) IndexOf(p Product) (int, error) {
	for idx, prod := range products {
		if p == prod {
			return idx, nil
		}
	}
	return 0, ErrProductNotFound
}

func (products Products) Includes(p Product) bool {
	_, err := products.IndexOf(p)
	return err == nil
}

/*
func (products Products) FilterCostlyProducts() Products {
	var result Products
	for _, p := range products {
		if p.Cost > 1000 {
			result = append(result, p)
		}
	}
	return result
}

func (products Products) FilterStationaryProducts() Products {
	var result Products
	for _, p := range products {
		if p.Category == "Stationary" {
			result = append(result, p)
		}
	}
	return result
}
*/

func (products Products) Filter(predicate func(Product) bool) Products {
	var result Products
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}
func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	// IndexOf
	fmt.Println("IndexOf")
	stove := Product{102, "Stove", 5000, 5, "Utencil"}
	if idx, err := products.IndexOf(stove); err == nil {
		fmt.Println("Index Of Stove : ", idx)
	} else {
		fmt.Println("Stove not found")
	}

	// Includes
	fmt.Println("Is stove in the product list ?:", products.Includes(stove))

	// filter costly products
	fmt.Println("Costly Products")
	// costlyProducts := products.FilterCostlyProducts()
	costlyProductPredicate := func(p Product) bool {
		return p.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println(costlyProducts)

	fmt.Println("Stationary Products")
	// stationaryProducts := products.FilterStationaryProducts()
	stationaryProductPredicate := func(p Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println(stationaryProducts)
}
