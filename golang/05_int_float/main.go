package main

import (
	"fmt"
)

func main() {
	view1 := 1000
	view2 := 2000
	total := view1 + view2
	fmt.Println(total)

	likes := 10

	// increment
	likes++
	fmt.Println(likes)

	// decrement
	likes--
	fmt.Println(likes)

	// average
	avgLikes := likes / 2
	fmt.Println(avgLikes)

	rating1 := 4.5
	rating2 := 3.5
	fmt.Println((rating1 + rating2) / 2)
}
