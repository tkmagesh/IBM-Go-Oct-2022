package main

import "fmt"

func main() {
	/*
		var name string
		fmt.Println("Enter the name :")
		fmt.Scanln(&name)
	*/
	var fname, lname string
	fmt.Println("Enter firstname and lastname :")
	//fmt.Scanln(&fname, &lname)
	fmt.Scanf("%s , %s\n", &fname, &lname)
	fmt.Printf("Hi %s %s, Have a nice day!\n", fname, lname)
}
