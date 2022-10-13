/* buffered channel operations */

package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 100
	fmt.Println(<-ch)
	fmt.Println("Done")
}
