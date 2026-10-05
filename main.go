package main

import (
	"fmt"
	"net/http"
)

func Greet(name string) string {
	if name == "" {
		name = "мир"
	}
	return "Привет, " + name + "!!!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, Greet(r.URL.Query().Get("name")))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server is running on http://localhost:8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
