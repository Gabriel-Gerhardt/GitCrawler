package main

import (
	"fmt"
	"gitcrawler/app/impl/facade"
)

func main() {
	f := facade.RepositoryFacade{}
	url := "https://github.com/Gabriel-Gerhardt/a.git"
	err := f.GetAllRepositoryFiles(url)
	if err != nil {
		return
	}
	fmt.Println("clone on")

}
