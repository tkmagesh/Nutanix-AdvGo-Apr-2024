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

// fmt.Stringer interface implementation
func (p Product) String() string {
	return p.Format()
}

func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

// composition
type PerishableProduct struct {
	Product
	Expiry string
}

// overriding the Format() method\
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, expiry : %q", pp.Product.Format(), pp.Expiry)
}

// fmt.Stringer interface implementation
func (pp PerishableProduct) String() string {
	return pp.Format()
}

func main() {
	// pen := Product{Id: 100, Name: "Pen", Cost: 10}
	pen := &Product{Id: 100, Name: "Pen", Cost: 10}

	// fmt.Println(Format(pen))
	fmt.Println("Before applying discount")
	// fmt.Println(pen.Format())
	fmt.Println(pen)

	fmt.Println("After applying 10% discount")
	// ApplyDiscount(&pen, 10)
	pen.ApplyDiscount(10)
	// fmt.Println(pen.Format())
	fmt.Println(pen)

	fmt.Printf("\nComposition\n")
	grapes := PerishableProduct{
		Product: Product{
			Id:   200,
			Name: "Grapes",
			Cost: 50,
		},
		Expiry: "2 Days",
	}
	// fmt.Println(grapes.Format())
	fmt.Println(grapes)
	fmt.Println("After applying 10% discount")
	grapes.ApplyDiscount(10)
	// fmt.Println(grapes.Format())
	fmt.Println(grapes)

}

func ApplyDiscount(p *Product, discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
