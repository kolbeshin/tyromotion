package main

import (
	"log"
	"tyromotion/backend/internal/app"
	"tyromotion/backend/internal/config"
)

func main() {
	cfg, err := config.Parse("./config/config.json")
	if err != nil {
		panic("Can`t parse a config file.")
	}
	log.Fatal(app.Run(cfg))
}
