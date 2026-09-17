package main

import "fmt"

func main() {
	items := 3
	individualPrice := 12

	if totalPrice := items * individualPrice; totalPrice > 100 {
		fmt.Println("Eligible for shipping")
	} else {
		fmt.Println("Not eligible for shipping")
	}
}
