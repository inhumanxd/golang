package main

import "fmt"

var package_variable int = 5

func main() {
	fmt.Println("Packages Tutorial kek.")

	var variable int = 69

	{
		var another_variable int = 96
		fmt.Println((another_variable))
	}

	fmt.Println(variable)
	notMain()
}

func notMain() {
	fmt.Println("In not main", package_variable)
}
