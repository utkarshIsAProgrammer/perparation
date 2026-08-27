package main

import "fmt"

func main() {
	// pointer address
	num := 10
	fmt.Println(&num)

	// create pointer
	numPtr := &num
	fmt.Println(numPtr)

	// type of the pointer value
	fmt.Printf("%T\n", numPtr)

	// get value through pointer
	fmt.Println(*numPtr)

	// change value of the original variable
	*numPtr = 55
	fmt.Println(num)

	// nil pointer
	var ptr *int
	fmt.Println(ptr) // <nil>
}
