/* concurrency basics */

package main

import (
	"fmt"
)

func main() {
	go f1()
	f2()
	fmt.Scanln() //blocking the execution of the main function so that the scheduler executes the scheduled goroutines. DO NOT DO THIS!!!
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
