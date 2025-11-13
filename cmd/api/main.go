package main

import (
	"log"

	"github.com/kevalsabhani/splitify/internal/bootstrap"
)

func main() {
	app, err := bootstrap.NewApp()
	if err != nil {
		log.Fatalf("init failed: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
