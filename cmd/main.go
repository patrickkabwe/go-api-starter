package main

import (
	"github.com/patrickkabwe/go-api-starter/cmd/api"
	"log"
)

func main() {
	apiServer := api.NewAPIServer()

	err := apiServer.Start(":5200")
	if err != nil {
		log.Fatal(err)
	}
}
