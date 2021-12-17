package main

import "fmt"

func main() {
	fmt.Println("Loops in GO language.")

	fmt.Println("\nFor as while loop.")

	var i uint = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\nFor loop.")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\nFor loop a little dynamic.")

	for i, j := 1, 2; i < 10; i, j = j+1, j+2 {
		fmt.Println(i, j)
	}

	fmt.Println("\nFor loop to print a string.")
	s := "Hello world"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	println()

	fmt.Println("\nFor loop to print a string. - For range of a string")

	for i, c := range s {
		fmt.Printf("%d %c\n", i, c)
		if i > 3 {
			break
		}
	}

	fmt.Println("\nLabelled For Loop")

outerForLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j > 6 {
				break outerForLoop
			}
			fmt.Printf("%d%d ", i, j)
		}
		fmt.Println()
	}
}
