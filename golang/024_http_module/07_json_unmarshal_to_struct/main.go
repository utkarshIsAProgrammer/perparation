package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	url := "https://catfact.ninja/fact"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("status code:", res.StatusCode)
		return
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read body failed:", err)
		return
	}

	var data CatFact

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		fmt.Println("JSON unmarshal failed:", err)
		return
	}

	fmt.Println(data.Fact)
	fmt.Println(data.Length)
}
