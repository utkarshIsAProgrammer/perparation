package greet

import "strings"

// exported functions start with a capital letter
func Hello(name string) string {
	clean := normal(name)
	return "Hello " + clean

}

func normal(name string) string {
	n := strings.TrimSpace(name)
	if n == "" {
		return "Guest"
	} else {
		return strings.ToUpper(name)
	}
}
