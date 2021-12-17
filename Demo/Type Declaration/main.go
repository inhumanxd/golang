package main

import "fmt"

func main() {
	fmt.Println("Type Declarations uwu~~")

	// Fahrenheit and Celcius
	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius

	fmt.Println(f, c)
}
