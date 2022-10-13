package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func genNos(ch chan int) {
	var no int = 1
	timeoutCh := timeout(3 * time.Second)
LOOP:
	for {
		select {
		case ch <- no * 10:
			time.Sleep(500 * time.Millisecond)
			no++
		case <-timeoutCh:
			break LOOP
		}
	}
	close(ch)
}

func timeout(delay time.Duration) <-chan time.Time {
	var timeoutCh chan time.Time = make(chan time.Time)
	go func() {
		time.Sleep(delay)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
