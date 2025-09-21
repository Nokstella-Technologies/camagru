package main

import (
	"log"

	"github.com/Nokstella-Technologies/camagru/application"
)

func main() {
	app := application.New()
	err := app.Run("8080")
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
