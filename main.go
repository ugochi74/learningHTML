package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serves everything in the current directory (including index.html)
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("Server running! Open your browser to: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
