package main

import "fmt"

func main() {
	fmt.Println(("Hello World :)"))

	// var	name type = expression
	var variable int = 69        // Omitting nothing
	var var1, var2, var3 int     // Omitting expression
	var name, age = "Bhavik", 22 //	Omitting Type

	//Short declaration
	//name := expression
	shortVar := 96
	boolVar := true

	fmt.Println(variable)
	fmt.Println(var1, var2, var3)
	fmt.Println(name, age)
	fmt.Println(shortVar, boolVar)
}
