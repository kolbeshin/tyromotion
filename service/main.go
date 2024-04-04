package main

import (
	"github.com/SenyashaGo/tyromotion/app"
	"github.com/SenyashaGo/tyromotion/config"
)

func main() {
	cfg, err := config.Parse("./config/config.json")
	if err != nil {
		panic("nn")
	}
	app.Run(cfg)
}
