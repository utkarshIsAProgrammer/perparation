package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func fetchCatFact() (CatFact, error) {
	url := "https://catfact.ninja/fact"

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		return CatFact{}, fmt.Errorf("failed to call external API: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return CatFact{}, fmt.Errorf(
			"external API returned status: %s",
			res.Status,
		)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return CatFact{}, fmt.Errorf("failed to read response: %w", err)
	}

	var data CatFact

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return CatFact{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil
}

func externalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Only GET requests are allowed",
		})
		return
	}

	data, err := fetchCatFact()
	if err != nil {
		fmt.Println("FETCH ERROR:", err)

		writeJSON(w, http.StatusBadGateway, map[string]any{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"timestamp": time.Now().UTC(),
		"external": map[string]any{
			"source": "Catfact.ninja",
			"fact":   data.Fact,
			"length": data.Length,
		},
	})
}

func main() {
	http.HandleFunc("/external", externalHandler)

	fmt.Println("Server is running on http://localhost:5000")

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("SERVER ERROR:", err)
	}
}
