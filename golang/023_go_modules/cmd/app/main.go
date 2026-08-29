package main

import (
	"fmt"
	greet "go-modules/internal"
)

func main() {
	msg1 := greet.Hello("Indiedev")
	fmt.Println(msg1)
}
