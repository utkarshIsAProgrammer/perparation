package main

// Import the packages we need.
import (
	"encoding/json" // Used to convert Go data into JSON
	"fmt"           // Used for printing messages to the terminal
	"net/http"      // Provides HTTP server functionality
	"time"          // Used to get the current time
)

// successHandler handles requests made to the /ok route.
//
// w -> used to send a response back to the client
// r -> contains information about the incoming HTTP request
func successHandler(w http.ResponseWriter, r *http.Request) {

	// Tell the client that our response will be JSON.
	//
	// HTTP headers contain extra information about the response.
	// "Content-Type: application/json" means:
	// "The data I'm sending you is JSON."
	w.Header().Set("Content-Type", "application/json")

	// Send HTTP status code 200 (OK).
	//
	// 200 means the request was successful.
	w.WriteHeader(http.StatusOK)

	// Create the data that we want to send as a response.
	//
	// map[string]any means:
	// - string = keys must be strings
	// - any    = values can be of any type
	//
	// For example:
	// "ok"       -> boolean
	// "message"  -> string
	// "datetime" -> time.Time
	res := map[string]any{
		"ok":       true,
		"message":  "JSON encoded successfully!",
		"datetime": time.Now().UTC(), // Get the current UTC time
	}

	// Convert our Go map into JSON
	// and send it to the client.
	//
	// json.NewEncoder(w)
	//     -> creates a JSON encoder that writes directly to w
	//
	// Encode(res)
	//     -> converts the res map into JSON
	//
	// w is the HTTP response writer, so the JSON
	// is sent back to whoever made the request.
	//
	// _ means we are intentionally ignoring the error returned by Encode().
	_ = json.NewEncoder(w).Encode(res)
}

func main() {

	// Register a route.
	//
	// When someone visits:
	//
	//     http://localhost:5000/ok
	//
	// Go will call:
	//
	//     successHandler
	//
	http.HandleFunc("/ok", successHandler) // route

	fmt.Println("Server is running on port :5000")
	err := http.ListenAndServe(":5000", nil) // listen on port 5000
	fmt.Println(err)
}
