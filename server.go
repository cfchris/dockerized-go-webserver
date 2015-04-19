package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Starting static web server.\n")

	http.ListenAndServe(":8080", http.FileServer(http.Dir("/var/www/hello")))
}
