package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Strings, Strings Builder ig.")
	fmt.Println()
	var s string = "Hello World"

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0]) //Prints ascii code when alone
	fmt.Printf("%c is at - %d\n", s[0], s[0])
	fmt.Println(s[0:4])
	fmt.Println(s[:5])
	fmt.Println(s[6:])

	hello := s[0:6]
	world := s[6:]
	hello_world := hello + world
	fmt.Println(hello_world)

	fmt.Println("Hello\nWorld")
	fmt.Println("Hello\tWorld")
	fmt.Println("Hello\bWorld")

	var abc string = "ABC"
	var b []byte = []byte(abc)
	var abcFromb = string(b)
	fmt.Printf("abc variable - %s\nb variable - array of bytes(string) - %s\nabcFromb - %s\n", abc, b, abcFromb)

	fmt.Println(("\nStrings Library with Functions :)"))
	fmt.Println(strings.ToUpper(hello))
	fmt.Println(strings.ToLower(hello))
	fmt.Println(strings.Count(hello, "l"))
	fmt.Println(strings.Split(hello, "lo"))
	fmt.Println(strings.ReplaceAll(hello_world, "l", "w"))
	fmt.Println(hello, strings.HasPrefix(hello, "H"))
	fmt.Println(hello, strings.HasSuffix(hello_world, "World"), strings.TrimSuffix(hello_world, "World"))

	fmt.Println(("\nStrings Builder ^_^"))

	var sb strings.Builder
	fmt.Println("Capacity", sb.Cap())
	sb.WriteString("Hello ")
	fmt.Println("Capacity", sb.Cap())
	sb.WriteString("Kid")
	fmt.Println(sb.String())
	fmt.Println("Length", sb.Len())
	fmt.Println("Capacity", sb.Cap())
	sb.Grow(101)
	fmt.Println("Capacity", sb.Cap())
	fmt.Println(sb.String())
	fmt.Println("Length", sb.Len())
	sb.Reset()
	fmt.Println("Capacity", sb.Cap())
	fmt.Println(sb.String())
	fmt.Println("Length", sb.Len())

	sb.Write([]byte{77, 87})
	fmt.Println(sb.String())

	x := 132
	y := strconv.Itoa(x)
	fmt.Println(x, y)
	fmt.Printf("%T %T\n", x, y)
	y2 := "499"
	z, a2 := strconv.Atoi(y2) // Returns integer and error
	fmt.Println(a2)
	fmt.Printf("%T %T\n", y, z)
}
