package api

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/Lykalon/urlshortener/internal/config"
	"github.com/Lykalon/urlshortener/internal/lib"
)

func CreateShortLink(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed. Use POST.")
		return
	}
	
	var bodyReq BodyCreateRequest

	err := json.NewDecoder(r.Body).Decode(&bodyReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error() + ` ¯\_(ツ)_/¯`)
		return
	}
	if bodyReq.Url == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `Required field "url" missed. ¯\_(ツ)_/¯`)
		return
	}
	
	var bodyRes BodyCreateResponce

	shortLink, ok := conf.DB.FindShort(bodyReq.Url)
	if !ok {
		shortLink = lib.Generate()
		conf.DB.Save(shortLink, bodyReq.Url)
	}

	bodyRes.Data = lib.Decode(int64(shortLink))
	data, err := json.Marshal(bodyRes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `Internal server error ¯\_(ツ)_/¯`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}