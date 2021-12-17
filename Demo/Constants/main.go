package main

import "fmt"

func main() {
	fmt.Println("Contatns - Here we go.")

	const five int = 5
	fmt.Println(five)

	const (
		one = 1
		one1
		one2
		two = 2
		two1
		two2
		sixtyNine = 69
	)

	fmt.Println(one, one1, one2, two, two1, two2, sixtyNine)

	fmt.Println("\nAutomatically increasing values of constants")

	const (
		oneI = iota + 1
		twoI
		threeI
		fourI
	)
	fmt.Println(oneI, twoI, threeI, fourI)

	const (
		_ = 1 << (10 * iota)
		KB
		MB
		GB
		TB
		PB
		EB
		ZB
		YB // Value is super big that int will overflow
	)

	fmt.Println(KB, MB, GB, TB, PB, YB/ZB)

}
