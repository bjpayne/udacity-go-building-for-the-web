package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var cityPopulations = map[string]uint32{
	"Tokyo":       37435191,
	"Delhi":       29399141,
	"Shanghai":    26317104,
	"Sao Paulo":   21846507,
	"Mexico City": 21671908,
}

func index(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	_error := json.NewEncoder(response).Encode(cityPopulations)

	if _error != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("Server is starting on port 3000...")
	_error := http.ListenAndServe(":3000", nil)

	if _error != nil {
		fmt.Println(_error)
		return
	}
}
