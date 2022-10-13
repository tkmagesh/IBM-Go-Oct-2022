package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		fmt.Scanln()
		done <- true
	}()
	go genNos(ch, done)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func genNos(ch chan int, done chan bool) {
	var no int = 1
LOOP:
	for {
		select {
		case ch <- no * 10:
			time.Sleep(500 * time.Millisecond)
			no++
		case <-done:
			break LOOP
		}
	}
	close(ch)
}

/* func timeout(delay time.Duration) <-chan time.Time {
	var timeoutCh chan time.Time = make(chan time.Time)
	go func() {
		time.Sleep(delay)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
} */
