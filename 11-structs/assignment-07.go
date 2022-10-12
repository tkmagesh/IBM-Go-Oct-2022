/*
	Write the following functions for both "Product" and "PerishableProduct"
		1. Format
			returns the formatted string of the object
			ex: "Id=100, Name="Grapes", Cost=10, Units=10, Category="Fruits""
		2. ApplyDiscount
			update the given object's cost after applying the given percentage (product & perishable product)

*/
package main

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

}
