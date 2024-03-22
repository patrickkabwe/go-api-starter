package main

import (
	"github.com/patrickkabwe/go-api-starter/cmd/api"
	database "github.com/patrickkabwe/go-api-starter/db"
	"log"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":5200", db)

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
