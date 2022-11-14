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
