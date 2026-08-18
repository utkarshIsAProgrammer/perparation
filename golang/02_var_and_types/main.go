package main

import "fmt"

func main(){
	// Normal declaration
	var name string = "IndieDev"
	var age int = 20
	
	fmt.Println("Name:",name)
	fmt.Println("Age:",age)
	
	// Shorthand declaration
	role:="Golang Developer"
	fmt.Println("Role:",role)	

	interest, hobby:="Pentesting", "Coding"
	fmt.Println("Interest:",interest)
	fmt.Println("Hobby:",hobby)
}
