package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter, r *http.Request) {

	// Tell the client that our response will be JSON.
	w.Header().Set("Content-Type", "application/json")

	// Send HTTP status code 200 (OK).
	w.WriteHeader(http.StatusOK)

	// Create the data that we want to send as a response.
	res := map[string]any{
		"ok":       true,
		"message":  "JSON encoded successfully!",
		"datetime": time.Now().UTC(),
	}

	// Convert our Go map into JSON and send it to the client.
	_ = json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/ok", successHandler) // route

	fmt.Println("Server is running on port :5000")
	err := http.ListenAndServe(":5000", nil) // listen on port 5000
	fmt.Println(err)
}
