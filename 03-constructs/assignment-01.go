/*
Print all the prime numbers between 3 and 100
*/

package main

import "fmt"

func main() {
LOOP:
	for i := 3; i <= 100; i++ {
		for j := 2; j <= (i / 2); j++ {
			if i%j == 0 {
				continue LOOP
			}
		}
		fmt.Printf("%d is a prime\n", i)
	}
}
