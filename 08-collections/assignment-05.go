/*
Write a program that will print all the prime numbers between 3 to 100
genPrimeNos() function should generate the prime numbers
main() function should print the prime numbers
*/

package main

import "fmt"

func main() {
	//get the prime numbers from the function
	primeNos := generatePrimes(3, 100)
	//print the prime numbers
	fmt.Println(primeNos)
}

func generatePrimes(start, end int) []int {
	primeNos := make([]int, 0)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primeNos = append(primeNos, no)
		}
	}
	return primeNos
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
