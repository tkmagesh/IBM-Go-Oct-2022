package main

import (
	"fmt"
	"sync"
)

//share memory by communicating
func main() {
	//add => should add the number
	//main => print the added result
	//execute the add function as a goroutine
	//feel free to modify the add function
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, ch)
	wg.Wait()
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
