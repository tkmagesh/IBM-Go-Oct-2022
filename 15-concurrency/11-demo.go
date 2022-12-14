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
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
