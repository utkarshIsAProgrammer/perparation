package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("status code:", res.StatusCode)
		return
	}

	// read body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert body to string
	bodyText := string(bodyBytes)
	max := 250

	if len(bodyText) < max {
		max = len(bodyText)
	}

	fmt.Println(bodyText[:max])
}
