package riha

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

// Login returns a token
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	l, err := GetList()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(l); err != nil {
		log.Error("Error encoding system list")
		panic(err)
	}

}

// Reissue re-issues a new token based on an existing valid one
func GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
