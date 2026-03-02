package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"

	"github.com/Lykalon/urlshortener/internal/config"
	"github.com/Lykalon/urlshortener/internal/lib"
)

func GetFullLink(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed. Use GET.")
		return
	}

	var bodyReq BodyGetRequest
	
	err := json.NewDecoder(r.Body).Decode(&bodyReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error() + ` ¯\_(ツ)_/¯`)
		return
	}
	if bodyReq.Data == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `Required field "data" missed. ¯\_(ツ)_/¯`)
		return
	}
	if strings.Count(bodyReq.Data, "") != 11 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `Wrong length for field "data". ¯\_(ツ)_/¯`)
		return
	}

	var bodyRes BodyGetResponse

	fullLink, ok := conf.DB.FindFull(lib.Encode(bodyReq.Data))
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `Full link for short link not found`)
		return
	}

	bodyRes.Url = fullLink
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