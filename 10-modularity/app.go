package main

import (
	"fmt"
	/* "modularity-demo/calculator" */
	calc "modularity-demo/calculator"
	"modularity-demo/calculator/utils"
)

func main() {
	fmt.Println("module executed")
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("Op Count = ", calculator.OpCount())
	*/

	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Op Count = ", calc.OpCount())
	fmt.Println("is 21 Even no? :", utils.IsEven(21))

}
