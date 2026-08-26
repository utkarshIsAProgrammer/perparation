package main

import (
	"fmt"
)

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}

func ageVerification(age int) error {
	if age < 18 {
		return fmt.Errorf("Age must be 18 or above")
	}

	return nil
}

func main() {
	res, err := divide(7, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	err = ageVerification(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
