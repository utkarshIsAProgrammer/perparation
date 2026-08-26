package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./index.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Closed with defer!")

	defer file.Close()
}
