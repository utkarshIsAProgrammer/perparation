package main

import (
	"fmt"
	"strconv"
)

func run() {}

func parse(s string) (int, error) {
	value, error := strconv.Atoi(s)
	if error != nil {
		return 0, fmt.Errorf("Value must be a number!")
	}

	if value < 1 || value > 5 {
		return 0, fmt.Errorf("Value must be in between 1 and 5")
	}

	return value, nil
}

func main() {
	fmt.Println()
}
