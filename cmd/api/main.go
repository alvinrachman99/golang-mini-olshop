package main

import (
	"log"
	"mini-online-shop/external/database"
	"mini-online-shop/internal/config"
)

func main() {
	filename := "cmd/api/config.yaml"

	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

}
