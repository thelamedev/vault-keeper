package main

import (
	"fmt"
	"log"

	"github.com/thelamedev/vault-keeper/internal"
	"github.com/thelamedev/vault-keeper/internal/config"
	"github.com/thelamedev/vault-keeper/internal/database"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Connect(config)
	if err != nil {
		log.Fatal(err)
	}

	app := internal.LoadServer(config)

	err = app.Listen(fmt.Sprintf("%s:%s", config.Server.Address, config.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
