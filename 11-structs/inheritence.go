package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	Product
	//Dummy
	Expiry string
	Cost   float32
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
	grapes := PerishableProduct{
		Product: Product{
			Id:       100,
			Name:     "Grapes",
			Cost:     50,
			Units:    10,
			Category: "Food",
		},
		Expiry: "2 Days",
		Cost:   100,
	}
	fmt.Printf("%#v\n", grapes)
	fmt.Println(grapes.Expiry)
	fmt.Println(grapes.Product.Cost)
	fmt.Println(grapes.Cost)

	fmt.Println(grapes.Name)

	bread := NewPerishableProduct(102, "Bread", 20, 2, "Confectionary", "1 week")
	fmt.Println(bread)
}
