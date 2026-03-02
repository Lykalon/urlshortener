package database

import (
	"errors"
	"log"
)

var ErrorInvalidArg = errors.New("Invalid arg for storage.")

func InitStorage(s string) (IDatabase, error) {
	log.Println("Storage type -", s)
	switch s {

	case "local":
		storage := new(LocalStorage)
		return storage , nil

	case "postgres":
		storage := new(PgStorage)
		return storage, nil

	default:
		return nil, ErrorInvalidArg
	}
}