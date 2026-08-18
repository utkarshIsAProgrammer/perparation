package main

import (
	"fmt"
)

func main() {
	isLoggedIn := true
	isAdmin := false
	hasPremium := false

	fmt.Println("Is logged in:", isLoggedIn)
	fmt.Println("is admin:", isAdmin)
	fmt.Println("has premium:", hasPremium)

	// AND &&
	canOpenDashboard := isLoggedIn && hasPremium
	fmt.Println("Can open dashboard:", canOpenDashboard)

	// OR ||
	canDeletePost := isAdmin || (isLoggedIn && hasPremium)
	fmt.Println("Can delete:", canDeletePost)
}
