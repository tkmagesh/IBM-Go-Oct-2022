/* anonymous functions */

package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous function invoked")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a good day!\n", userName)
	}("Magesh")

	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println("result =", result)

}
