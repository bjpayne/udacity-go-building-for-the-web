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
