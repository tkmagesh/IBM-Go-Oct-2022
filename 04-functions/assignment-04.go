package main

import "fmt"

var counter int

func main() {
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	counter = 1000           //restrict this!!!!
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
}

func increment() int {
	counter += 1
	return counter
}
