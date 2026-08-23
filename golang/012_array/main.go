package main

import "fmt"

func main() {
	// fixed size and ordered collection of elements of the same type
	var marks [3]int

	marks[0] = 10
	marks[1] = 20
	marks[2] = 40
	fmt.Println(marks)

	// array literal and length
	res := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(res, len(res))
}
