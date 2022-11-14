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
