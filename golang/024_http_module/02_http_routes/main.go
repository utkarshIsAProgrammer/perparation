package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Welcome try to /home?name=indiedev!")) // response
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // get the name from the query string (url)
	if name == "" {
		name = "Guest"
	}

	_, _ = w.Write([]byte(name)) // response
}

func main() {
	// routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/home", homeHandler)

	fmt.Println("Server is running on port :5000")
	err := http.ListenAndServe(":5000", nil) // listening on port 5000
	fmt.Println(err)
}
