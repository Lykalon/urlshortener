package main

import (
	"log"
	"os"

	"github.com/Lykalon/urlshortener/internal/app"
	"github.com/Lykalon/urlshortener/internal/config"
	"github.com/Lykalon/urlshortener/internal/database"
)

func main() {
	config.Init()
	storage, err := database.InitStorage(config.GetConfig().Storage)
	if err != nil {
		log.Fatal("Storage storage obj creation failed: ", err)
		os.Exit(1)
	}
	log.Println("Storage object ready")
	log.Println("Storage init")
	storage.Init()
	log.Println("Storage init succes")
	config.SetStorage(storage)
	app.Init()
	//graceful shutdown
}