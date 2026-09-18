package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	// Tell the client that our response is JSON.
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code before writing the response body.
	w.WriteHeader(status)

	// Convert the data to JSON and write it to the response.
	// Encode() automatically adds a newline at the end.
	_ = json.NewEncoder(w).Encode(data)
}

// TestRequest represents the JSON data we expect from the client.
type TestRequest struct {
	// The `json:"name"` tag tells Go to map the JSON
	// field "name" to this struct field.
	Name string `json:"name"`
}

// testHandler handles requests coming to /test.
func testHandler(w http.ResponseWriter, r *http.Request) {

	// Check whether the client used the POST method.
	if r.Method != http.MethodPost {

		// Return HTTP 405 Method Not Allowed.
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Only POST method is allowed",
		})

		return
	}

	// Close the request body after we're finished reading it.
	// defer means this will execute when the function returns.
	defer r.Body.Close()

	// Create a variable to store the decoded JSON request.
	var req TestRequest

	// Create a JSON decoder that reads JSON from the request body.
	dec := json.NewDecoder(r.Body)

	// Decode the JSON body into our TestRequest struct.
	if err := dec.Decode(&req); err != nil {

		// If the JSON is invalid, return HTTP 400 Bad Request.
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid JSON format!",
		})

		return
	}

	// If everything is successful, return the received name.
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":   true,
		"name": req.Name,
	})
}

func main() {
	http.HandleFunc("/test", testHandler)

	fmt.Println("Server is running on port :5000")
	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)
}
