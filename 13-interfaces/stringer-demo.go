package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

//fmt.Stringer implementation
func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func main() {
	kettle := Product{101, "Kettle", 2500, 10, "Utencil"}
	fmt.Println(kettle)
}
