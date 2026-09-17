package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName:="Indie"
	lastName:="Dev"
	fullName:=firstName+lastName
	fmt.Println("Full Name is:",fullName)
	fmt.Println(strings.ToUpper(firstName), strings.ToUpper(lastName))
}
