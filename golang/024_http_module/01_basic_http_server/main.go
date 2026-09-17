package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// if the request method is not GET, return a 405 Method Not Allowed error
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	_, _ = w.Write([]byte("Hello from Go net/http server!")) // response
}

func main() {
	http.HandleFunc("/home", helloHandler) // route

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil) // listening on port 8080
	fmt.Println(err)
}
