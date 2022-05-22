package main

import (
	"Scoltest/internal/app"
	"Scoltest/internal/config"
	"Scoltest/pkg/loger"
)

const configPath = "config.yml"

func main() {

	l := loger.NewLogger()
	l.Infoln("Load config ", configPath)

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		l.Fatalln("Error load config ", err.Error())
	}

	server, err := app.New(cfg, l)
	if err != nil {
		l.Fatalln("Error create server ", err.Error())
	}

	server.Start()

}
