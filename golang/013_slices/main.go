package main

import "fmt"

func main() {
	// can grow and shrink
	res := []int{1, 2, 3, 4, 5, 6}
	res = append(res, 7, 8, 9)

	fmt.Println(res, len(res))
	fmt.Println(res[0:2], res[2:len(res)-1])

	// capacity
	nums := make([]int, 0, 5)
	fmt.Println(nums, len(nums), cap(nums))

	nums = append(nums, 10)
	nums = append(nums, 20, 30, 40, 50)
	fmt.Println(nums, len(nums), cap(nums))

	// if we exceeding the capacity, it will double the capacity of backing array
	nums = append(nums, 75)
	fmt.Println(nums, len(nums), cap(nums))

	// variadic/ellipsis operator (add elements of one slice to another)
	todos := []string{"learn golang", "do calisthenics", "eat good"}
	more := []string{"sleep good"}
	new := append(todos, more...)
	fmt.Println(new)
}
