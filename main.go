package main

import (
	"log"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/cmd/app"
	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/config"
)

func main() {
	// create the server config instance
	serverConfig, err := config.LoadServerConfig()

	if err != nil {
		log.Fatal("Something went wrong while creating server config: " + err.Error())
	}

	// create the db config instance
	dbConfig, err := config.LoadDbConfig()

	if err != nil {
		log.Fatal("Something went wrong while creating database config: " + err.Error())
	}

	// create the app instance
	serverApp := &app.App{
		ServerConfig: serverConfig,
		DbConfig:     dbConfig,
	}

	serverApp.Run()
}
