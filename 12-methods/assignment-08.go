/*

Refactor the "Format" & "ApplyDiscount" functions as methods to the respective structs

*/
package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Units:    units,
			Category: category,
		},
		Expiry: expiry,
	}
}

func main() {
	pen := Product{100, "Pen", 10, 20, "Stationary"}
	fmt.Println(pen.Format())
	fmt.Println("After applying 10% discount")
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())

	grapes := NewPerishableProduct(102, "Grapes", 20, 10, "Fruits", "2 Days")
	fmt.Println(grapes.Format())
	fmt.Println("After applying 10% discount")
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())
}

func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%v, Units=%d, Category=%q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry=%q", pp.Product.Format(), pp.Expiry)
}
