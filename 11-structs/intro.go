package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{100, "Magesh", "Bangalore"}
	//var emp Employee = Employee{Id: 100, Name: "Magesh"}
	//var emp = Employee{Id: 100, Name: "Magesh"}
	//emp := Employee{Id: 100, Name: "Magesh"}
	//emp := Employee{} // equivalent to "var emp Employee"
	//emp := new(Employee)
	//emp := &Employee{}
	//emp := &Employee{Id: 100, Name: "Magesh"}
	//fmt.Printf("%#v\n", emp)

	fmt.Println("Accessing the fields")
	emp := Employee{Id: 100, Name: "Magesh", City: "Bangalore"}
	fmt.Println(emp.Id, emp.Name, emp.City)

	fmt.Println("Accessing the fields from a struct pointer")
	empPtr := &Employee{Id: 100, Name: "Magesh", City: "Bangalore"}
	fmt.Println((*empPtr).Id, (*empPtr).Name, (*empPtr).City)
	fmt.Println(empPtr.Id, empPtr.Name, empPtr.City)

}
