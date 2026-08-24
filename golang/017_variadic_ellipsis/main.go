package main

import "fmt"

func sumAll(num ...int) int {
	sum := 0
	for _, val := range num {
		sum += val
	}
	return sum
}

func main() {
	res := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(res)

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res = sumAll(values...)
	fmt.Println(res)
}
