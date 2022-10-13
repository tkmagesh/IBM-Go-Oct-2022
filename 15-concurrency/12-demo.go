//deadlock simulation

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go fn()
	wg.Wait()
}

func fn() {
	fmt.Println("fn invoked")
	wg.Done()
}
