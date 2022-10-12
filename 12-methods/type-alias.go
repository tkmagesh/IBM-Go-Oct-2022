package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

type Employee struct {
	Id   int
	Name string
	City string
}

func (e Employee) WhoAmI() {
	fmt.Println("I am the employee - ", e.Name)
}

//type aliasing
/*
type Associate Employee

func (a Associate) WhoAmI() {
	Employee(a).WhoAmI()
}
*/

//composition
type Associate struct {
	Employee
}

//whoAmI method is inherited from Employee

func main() {
	s := MyStr("This is a string")
	fmt.Println(s.Length())

	//type alias
	/*
		emp := Associate{Id: 100, Name: "John"}
		emp.WhoAmI()
	*/

	emp := Associate{Employee: Employee{Id: 100, Name: "John"}}
	emp.WhoAmI()
}
