package calculator

import "fmt"

func init() {
	fmt.Println("calculator package initialized [subtract.go]")
}

func Subtract(x, y int) int {
	op_count++
	return x - y
}
