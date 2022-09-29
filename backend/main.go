package main

import (
	"github.com/Fatur525/learninggo/api"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := api.SetupDb()
	if err != nil {
		panic(err)
	}

	server := api.MakeServer(db)
	server.RunServer()
}
