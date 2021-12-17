package main

import "fmt"

func main() {
	fmt.Println("Hewlo to Pointers Turorial ^_^")
	p := 10
	_p := &p

	fmt.Printf("\n%d is stored at %p!.\n", p, _p)
	fmt.Printf("%p holds the value %d.\n", _p, *_p)

	var p1 int = 20
	var _p1 *int = &p1

	fmt.Printf("\n%d is stored at %p!.\n", p1, _p1)
	fmt.Printf("%p holds the value %d.\n", _p1, *_p1)
}
