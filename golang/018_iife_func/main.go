package main

import "fmt"

func main() {

	// IIFE function
	func() {
		fmt.Println("IIFE - Immediately Invoked Function Expression")
	}()

	ans := func(name string) string {
		return fmt.Sprintf("Hey! %s", name)
	}("IndieDev")
	fmt.Println(ans)
}
