package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func multiReturn(a int, b int) (txt string, sum int) {
	sum = a + b
	txt = fmt.Sprintf("Sum of %d and %d is %d", a, b, sum)
	return txt, sum
}

func main() {
	res := add(5, 7)
	fmt.Println(res)

	txt, _ := multiReturn(7, 9)
	fmt.Println(txt)
}
