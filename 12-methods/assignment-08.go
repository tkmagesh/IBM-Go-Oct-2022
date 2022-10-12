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
	fmt.Println(FormatProduct(pen))
	fmt.Println("After applying 10% discount")
	ApplyDiscount(&pen, 10)
	fmt.Println(FormatProduct(pen))

	grapes := NewPerishableProduct(102, "Grapes", 20, 10, "Fruits", "2 Days")
	fmt.Println(FormatPerishableProduct(*grapes))
	fmt.Println("After applying 10% discount")
	ApplyDiscountPP(grapes, 10)
	fmt.Println(FormatPerishableProduct(*grapes))
}

func FormatProduct(p Product) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%v, Units=%d, Category=%q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(p *Product, discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func FormatPerishableProduct(pp PerishableProduct) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%v, Units=%d, Category=%q, Expiry=%q", pp.Id, pp.Name, pp.Cost, pp.Units, pp.Category, pp.Expiry)
}

func ApplyDiscountPP(pp *PerishableProduct, discountPercentage float32) {
	pp.Cost = pp.Cost * ((100 - discountPercentage) / 100)
}
