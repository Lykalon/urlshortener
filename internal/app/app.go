package app

import (
	"log"
	"net/http"

	"github.com/Lykalon/urlshortener/internal/api"
	"github.com/Lykalon/urlshortener/internal/config"
)

func Init() {
	http.HandleFunc("/api/create", api.CreateShortLink)
	http.HandleFunc("/api/get", api.GetFullLink)
	http.ListenAndServe(":" + config.GetConfig().Port, nil)
	log.Println("Listening for port ", config.GetConfig().Port)
}

func Stop() {

}