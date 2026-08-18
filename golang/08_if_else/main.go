package main

import "fmt"

func main() {
	score := 78
	if score >= 90 {
		fmt.Println("You got A grade")
	} else if score >= 75 {
		fmt.Println("You got B grade")
	} else if score >= 45 {
		fmt.Println("You got C grade")
	} else {
		fmt.Println("You failed the exam")
	}
}
