package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
	Org  *Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	ibm := Organization{
		Name: "IBM",
		City: "Bangalore",
	}
	emp1 := Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bangalore",
		Org:  &ibm,
	}
	fmt.Printf("%#v\n", emp1)

	emp2 := Employee{
		Id:   101,
		Name: "Suresh",
		City: "Bangalore",
		Org:  &ibm,
	}
	fmt.Printf("%#v\n", emp2)
	fmt.Println("After changing Org City to Pune")
	ibm.City = "Pune"
	fmt.Printf("%p\n", &ibm)
	fmt.Printf("%#v\n", emp1.Org.City)
	fmt.Printf("%#v\n", emp2.Org.City)
}
