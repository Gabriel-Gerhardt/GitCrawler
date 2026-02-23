package main

import (
	"fmt"
	"gitcrawler/app/impl/adapters/register"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	server := http.Server{}
	register.GetHandlers()
	server.Addr = ":8080"
	fmt.Println("Running server on " + server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
