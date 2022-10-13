package main

import (
	"fmt"
)

//share memory by communicating
func main() {
	//add => should add the number
	//main => print the added result
	//execute the add function as a goroutine
	//feel free to modify the add function

	ch := add(100, 200)
	result := <-ch
	fmt.Println(result)
}

func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		result := x + y
		ch <- result
	}()
	return ch
}
