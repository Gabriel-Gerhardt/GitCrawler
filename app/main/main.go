package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{}

	server.Addr = ":8080"
	fmt.Println("Running server on " + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
