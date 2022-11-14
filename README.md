# UDACITY - Building for the Web

HTTP Fundamentals, net/HTTP Package, Handlers, Serving HTML, Routing

HTTP Fundamentals 1 and net/HTTP Package

```go
package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func index(response http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(response, "Hello World :]")

	if err != nil {
		return
	}
}

func cityList(response http.ResponseWriter, request *http.Request) {
	html := "<h1>List of most populous cities</h1>"

	for _, city := range cities {
		html += fmt.Sprintf("<p>%s</p>", city)
	}

	_, err := fmt.Fprintf(response, html)

	if err != nil {
		return
	}
}

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/", index)
	http.HandleFunc("/city-list", cityList)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)

		return
	}
}
```

Handlers

```go
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
```

Serving HTML

```go
```