package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go lang.")
	var arr [10]int
	fmt.Println(arr)

	fmt.Println(arr[0])
	arr[0] = 1
	fmt.Println(arr[0])

	var arr2 = arr
	fmt.Println(arr2 == arr)
	fmt.Println(arr2)

	arr[1] = 2
	fmt.Println(arr)
	fmt.Println(arr2)

	var arr3 [4]int = [4]int{11, 22, 33}
	fmt.Println(arr3)

	var arr4 = [4]int{11, 22, 33}
	fmt.Println(arr4)

	var arr5 = [...]int{11, 22, 33}
	fmt.Println(arr5)

	arr6 := [...]int{10, 11, 12, 13, 14, 15}
	fmt.Println(arr6)

	for _, v := range arr6 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	arr7 := [...]int{0: 1}
	fmt.Println(arr7)

	arr8 := [...]int{0: 1, 10: 10, 20: 20}
	fmt.Println(arr8)

	strArr := [...]string{"string", "String", "sTRING"}
	fmt.Println(strArr)

	arr2D := [...][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(arr2D)

	arr3D := [...][2][2]int{{{1, 2}, {1, 2}}, {{2, 3}, {2, 3}}}
	fmt.Println(arr3D)

	// While declaring multi-dimensional arrays only the length of outer most array can be dynamic.
}
