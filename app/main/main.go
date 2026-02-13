package main

import (
	"fmt"
	"gitcrawler/app/impl/service"
)

func main() {
	c := service.CloneService{}
	url := "https://github.com/Gabriel-Gerhardt/Webhook-Manager.git"
	path, err := c.CloneRepository(url)
	if err != nil {
		return
	}
	fmt.Println(path + "clone on")

}
