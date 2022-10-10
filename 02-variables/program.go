package main

import "fmt"

var x = 100

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/
	/*
		var msg string = "Hello World!"
	*/

	/*
		//type inference
		var msg = "Hello World!"
	*/

	msg := "Hello World!" //=> := is allowed ONLY in a function (not in package level)
	fmt.Println(msg)

	/*
		var myVar int
		myVar = 100
		fmt.Println(myVar)
	*/

	//declaring multiple variables & initializing them
	/*
		var x int
		var y int
		var str string
		var result int
		str = "Sum of x and y is :"
		x = 100
		y = 200
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		str = "Sum of x and y is :"
		x = 100
		y = 200
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		str = "Sum of x and y is :"
		x = 100
		y = 200
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of x and y is :"
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			result int    = x + y
			str    string = "Sum of x and y is :"
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   = 100, 200
			result = x + y
			str    = "Sum of x and y is :"
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of x and y is :"
			result    = x + y
		)
		fmt.Println(str, result)
	*/

	x, y, str := 100, 200, "Sum of x and y is :"
	result := x + y
	fmt.Println(str, result)
}
