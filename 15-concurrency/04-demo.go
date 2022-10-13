/* WaitGroup */

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10) //set the wg counter
	for i := 0; i < 10; i++ {
		go f1()
	}
	f2()
	wg.Wait() // blocks the execution of this function until the counter becomes 0
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() //decrement the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
