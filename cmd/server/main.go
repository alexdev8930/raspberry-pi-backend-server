package main

import (
	"log"
	"rpi-backend.local/internal/app"
	"rpi-backend.local/internal/config"
)

func main() {
	cfg := config.Load()
	a := app.New(cfg)
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
