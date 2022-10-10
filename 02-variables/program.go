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

	//using complex numbers
	c1 := 4 + 5i
	c2 := 3 + 4i

	fmt.Println(c1 + c2)
	fmt.Println(real(c1))
	fmt.Println(imag(c1))

	//type conversion
	var i int = 100
	var f float32
	f = float32(i)
	fmt.Println(f)

	//constants
	//const pi = 3.14

	const (
		PI float32 = 31.4
		NO int     = 100
	)
	fmt.Println(PI, NO)

	//iota
	/*
		const (
			RED   = 0
			GREEN = 1
			BLUE  = 2
		)
	*/

	/*
		const (
			RED   = iota
			GREEN = iota
			BLUE  = iota
		)
	*/

	/*
		const (
			RED = iota
			GREEN
			BLUE
		)
	*/
	/*
		const (
			RED = iota
			GREEN
			_
			BLUE
		)
	*/

	/*
		const (
			RED   = iota + 2 //0 + 2
			GREEN            //1 + 2
			BLUE             //2 + 2
		)
	*/

	const (
		RED = iota * 2
		GREEN
		BLUE
	)

	fmt.Printf("RED = %d, GREEN = %d, BLUE = %d\n", RED, GREEN, BLUE)

	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
}
