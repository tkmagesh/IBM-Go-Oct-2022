package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func (e Employee) WhoAmI() {
	fmt.Println("I am the employee - ", e.Name)
}

func (e *Employee) ChangeName(newName string) {
	fmt.Printf("[ChangeName] %p\n", e)
	e.Name = newName
}
func main() {
	emp := Employee{Name: "Magesh"}
	emp.WhoAmI()
	fmt.Printf("[main] %p\n", &emp)
	//(&emp).ChangeName("Suresh")
	emp.ChangeName("Suresh")
	emp.WhoAmI()
}
