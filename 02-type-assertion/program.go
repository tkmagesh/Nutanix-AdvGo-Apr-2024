package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) Format() string {
	return fmt.Sprintf("id : %d, name : %q, cost : %0.2f", p.Id, p.Name, p.Cost)
}

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Ad nulla ea Lorem Lorem deserunt in ullamco voluptate ullamco voluptate elit consequat amet."
	x = true
	x = 99.999
	fmt.Println(x)

	x = 200
	// x = "In tempor aliquip nulla fugiat."
	// y := x + 200
	// y := x.(int) + 200
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	x = Product{200, "Marker", 50}
	// x = 100
	// x = "Ad nulla ea Lorem Lorem deserunt in ullamco voluptate ullamco voluptate elit consequat amet."
	// x = true
	// x = 99.999
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case float64:
		fmt.Println("x is a float64, x * 99% =", val*0.99)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case Product:
		fmt.Println(val.Format())
	default:
		fmt.Println("x is of unknown type")
	}

}
