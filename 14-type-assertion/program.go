package main

import "fmt"

func main() {
	//var x interface{}
	var x any
	x = 100
	x = "This is a string"
	x = true
	x = 100.32
	x = []int{10, 20, 30}
	fmt.Println(x)

	//x = 100
	x = "This is a string"
	/*
		y := x + 200
		fmt.Println(y)
	*/
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println("y =", y)
	} else {
		fmt.Println("Non integer value. cannot add.")
	}

	//x = "THis is a long string"
	//x = true
	x = 100.23
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 = ", val+200)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is bool, !x = ", !val)
	default:
		fmt.Println("unknown type")
	}
}
