package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func main() {
	errenv := godotenv.Load()
	if errenv != nil {
		log.Panic("Could not load the .env")
	}

	var port = os.Getenv("PORT")
	if len(port) == 0 {
		log.Panic("Missing PORT environment variable")
	}
	log.Info("* Starting PaaS Auth Service * ")

}
