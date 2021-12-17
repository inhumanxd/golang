package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Math libraries and numbers")

	var x int = 5
	var x32 int32 = 5
	var x64 int64 = 5
	var ux uint = 5

	fmt.Printf("%T, %T, %T, %T \n", x, x32, x64, ux)

	fmt.Println("\nMath libraries and numbers - Floating points")

	var pi32 float32 = 3.141
	var pi64 float64 = 3.141
	pi := 3.141

	fmt.Printf("%T, %T, %T \n", pi, pi32, pi64)

	fmt.Println("\nMath libraries and numbers - Scientific Notations")

	var sn float64 = 1e140

	fmt.Println(sn + 9e200)

	fmt.Println("\nMath libraries and numbers - Type Conversion")

	x = int(pi)
	pi = float64(ux)

	fmt.Println(x, pi)
	fmt.Printf("%d - %T, %F - %T \n", x, x, pi, pi)

	fmt.Println("\nMath libraries and numbers - Library")

	piCeil := 3.00000000001
	fmt.Println(math.Ceil(piCeil))
	fmt.Println(math.Min(69, 69.01))
	fmt.Println(math.Max(69, 69.01))
	fmt.Println(math.Abs(-69))
	fmt.Println(math.Sqrt(256))
	fmt.Println(math.Cbrt(512))
	fmt.Println(math.Pow(10, 69))

	fmt.Println("\nMath libraries and numbers - Complex numbers")
	fmt.Println(complex(1, 2))

}
