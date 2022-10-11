package main

import "fmt"

func main() {
	increment := getIncrement()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	//counter = 1000           //restrict this!!!!
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
}

func getIncrement() func() int {
	var counter int
	return func() int {
		counter += 1
		return counter
	}
}
