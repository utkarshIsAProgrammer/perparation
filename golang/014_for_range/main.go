package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	total := 0

	// for index, value := range collection_name{}
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
