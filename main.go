package main

import (
	"github.com/rakesh740/csv_reader/matrix"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", matrix.Echo)
	http.HandleFunc("/invert", matrix.Invert)
	http.HandleFunc("/flatten", matrix.Flatten)
	http.HandleFunc("/sum", matrix.Sum)
	http.HandleFunc("/multiply", matrix.Multiply)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
