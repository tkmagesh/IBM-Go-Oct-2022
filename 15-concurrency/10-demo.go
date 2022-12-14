package main

import (
	"fmt"
	"sync"
)

//communicate by sharing memory
var result int

func main() {
	//add => should add the number
	//main => print the added result
	//execute the add function as a goroutine
	//feel free to modify the add function
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	result = x + y
}
