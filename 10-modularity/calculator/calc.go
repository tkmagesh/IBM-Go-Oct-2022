package calculator

import "fmt"

var op_count int

var Nos []int

func init() {
	Nos = make([]int, 0, 10)
	fmt.Println("calculator package initialized [calc.go] - 1")
}

func init() {
	fmt.Println("calculator package initialized [calc.go] - 2")
}

func OpCount() int {
	return op_count
}
