package main

import (
	"fmt"
	"gitcrawler/app/impl/register"
	"net/http"
)

func main() {
	server := http.Server{}
	register.GetHandlers()
	server.Addr = ":8080"
	fmt.Println("Running server on " + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
