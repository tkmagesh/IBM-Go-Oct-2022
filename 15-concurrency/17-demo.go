package main

import "fmt"

func main() {
	ch := make(chan int)
	go fn(ch)
	for val := range ch {
		fmt.Println(val)
	}
}

func fn(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i * 10
	}
	close(ch)
}
