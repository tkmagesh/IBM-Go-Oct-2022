package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int
	noPtr = &no
	fmt.Println(noPtr)

	//deferencing
	no2 := *noPtr
	fmt.Println(no2)
	fmt.Println(no == *(&no))

	fmt.Println("Before incrementing, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	n1, n2 := 10, 20
	fmt.Printf("Before swapping n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping n1 = %d and n2 = %d\n", n1, n2)
}

func increment(x *int) {
	*x += 1
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}
