package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
	Write the apis for the following

	IndexOf => return the index of the given product
		ex:  returns the index of the given product

	Includes => return true if the given product is present in the collection else return false
		ex:  returns true if the given product is present in the collection

	Filter => return a new collection of products that satisfy the given condition
		use cases:
			1. filter all costly products (cost > 1000)
				OR
			2. filter all stationary products (category = "Stationary")
				OR
			3. filter all costly stationary products
			etc

	Any => return true if any of the product in the collection satifies the given criteria
		use cases:
			1. are there any costly product (cost > 1000)?
			2. are there any stationary product (category = "Stationary")?
			3. are there any costly stationary product?
			etc

	All => return true if all the products in the collection satifies the given criteria
		use cases:
			1. are all the products costly products (cost > 1000)?
			2. are all the products stationary products (category = "Stationary")?
			3. are all the products costly stationary products?
			etc


	Sort => Sort the products collection by any attribute
		IMPORTANT : MUST Use sort.Sort() function
		use cases:
			1. sort the products collection by cost
			OR
			2. sort the products collection by name
			OR
			3. sort the products collection by units
			OR
			4. sort the products collection by cost in descending order
			OR
			5. sort the products collection by name in descending order
			OR
			6. sort the products collection by units in descending order
			OR
			etc



*/

type Products []Product

func (products Products) Print() {
	for _, p := range products {
		fmt.Println(p.Format())
	}
}

func (products Products) IndexOf(p Product) int {
	for idx, product := range products {
		if product == p {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(p Product) bool {
	for _, product := range products {
		if product == p {
			return true
		}
	}
	return false
}

/*
func (products Products) FilterCostlyProducts() Products {
	result := Products{}
	for _, product := range products {
		if product.Cost > 1000 {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) FilterStationaryProducts() Products {
	result := Products{}
	for _, product := range products {
		if product.Category == "Stationary" {
			result = append(result, product)
		}
	}
	return result
}
*/

//2 days ago
func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
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
	products.Print()

	kettle := Product{101, "Kettle", 2500, 10, "Utencil"}
	fmt.Println("Index of kettle = ", products.IndexOf(kettle))

	fmt.Println("Filter")
	fmt.Println("Costly Products")
	//costlyProducts := products.FilterCostlyProducts()
	costlyProductPredicate := func(p Product) bool {
		return p.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	costlyProducts.Print()

	fmt.Println("Stationary Products")
	stationaryProductsPredicate := func(p Product) bool {
		return p.Category == "Stationary"
	}
	//stationaryProducts := products.FilterStationaryProducts()
	stationaryProducts := products.Filter(stationaryProductsPredicate)
	stationaryProducts.Print()

	fmt.Println("Costly Stationary Products")
	costlyStationaryProductPredicate := func(p Product) bool {
		return costlyProductPredicate(p) && stationaryProductsPredicate(p)
	}
	costlyStationaryProducts := products.Filter(costlyStationaryProductPredicate)
	costlyStationaryProducts.Print()

	//today
	utencilProductsPredicate := func(p Product) bool {
		return p.Category == "Utencil"
	}
	utencilProducts := products.Filter(utencilProductsPredicate)
	utencilProducts.Print()

}
