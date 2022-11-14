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
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "./static/index.html")
	})

	// Handle with an anonymous function
	http.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "./static/hello.html")
	})

	fmt.Println("Server is starting on port 3000...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)

		return
	}
}
```

Routing

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	dictionary := map[string]string{"Go": "A language"}

	router := mux.NewRouter()

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "./static/index.html")
	}).Methods("GET")

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		// 1. Set content type to JSON
		response.Header().Set("Content-Type", "application/json")

		// 2. Keep track of new entry
		var newEntry map[string]string

		// 3. Read the request
		requestBody, _ := io.ReadAll(request.Body)

		// 4. Parse JSON body
		err := json.Unmarshal(requestBody, &newEntry)

		if err != nil {
			log.Println(err)
		}

		// 5. Add new entry into the dictionary
		for field, value := range newEntry {
			if _, ok := dictionary[field]; ok {
				// Field exists...
				_, fieldExistsMessageError := fmt.Fprintf(response, "%s exists", field)

				if fieldExistsMessageError != nil {
					log.Println(fieldExistsMessageError)
				}
			} else {
				// Add field to the dictionary
				dictionary[field] = value
				response.WriteHeader(http.StatusCreated)
			}
		}
		// 6. Return dictionary
		responseEncodeError := json.NewEncoder(response).Encode(dictionary)

		if responseEncodeError != nil {
			log.Println(responseEncodeError)
		}
	}).Methods("POST")

	router.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "./static/hello.html")
	}).Methods("GET")

	fmt.Println("Server is starting on port 3000...")
	err := http.ListenAndServe(":3000", router)
	log.Fatal(err)
}
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	dictionary := map[string]string{"Go": "A language"}

	router := mux.NewRouter()

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "./static/index.html")
	}).Methods("GET")

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		// 1. Set content type to JSON
		response.Header().Set("Content-Type", "application/json")

		status := map[string]string{"message": "success"}

		responseMessage := map[string]map[string]string{"_status": status, "data": dictionary}

		// 2. Keep track of new entry
		var newEntry map[string]string

		// 3. Read the request
		requestBody, _ := io.ReadAll(request.Body)

		// 4. Parse JSON body
		err := json.Unmarshal(requestBody, &newEntry)

		if err != nil {
			log.Println(err)
		}

		// 5. Add new entry into the dictionary
		for field, value := range newEntry {
			if _, ok := dictionary[field]; ok {
				// Field exists...
				responseMessage["_status"]["message"] = fmt.Sprintf("%s exists", field)
				response.WriteHeader(http.StatusConflict)
			} else {
				responseMessage["_status"]["message"] = "success"
				// Add field to the dictionary
				dictionary[field] = value
				response.WriteHeader(http.StatusCreated)
			}
		}

		// 6. Return dictionary
		responseEncodeError := json.NewEncoder(response).Encode(responseMessage)

		if responseEncodeError != nil {
			log.Println(responseEncodeError)
		}
	}).Methods("POST")

	router.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "./static/hello.html")
	}).Methods("GET")

	fmt.Println("Server is starting on port 3000...")
	err := http.ListenAndServe(":3000", router)
	log.Fatal(err)
}
```