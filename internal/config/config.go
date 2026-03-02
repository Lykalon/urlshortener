package config

import (
	"log"
	"os"

	"github.com/Lykalon/urlshortener/internal/database"
)

type Config struct {
	Env		string
	Port	string
	Storage	string
	DB		database.IDatabase
}

var conf Config

func Init() {
	log.Println("Config init")
	conf.Port		= "8080"
	conf.Storage	= os.Getenv("STORAGE")
	conf.DB			= nil
}

func SetStorage(storage database.IDatabase) {
	conf.DB = storage
}

func GetConfig() (config Config){
	return conf
}