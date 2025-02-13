package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting server at port 8090...")

	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Println("Error startign server:", err)
	}
}
