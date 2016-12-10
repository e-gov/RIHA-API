package riha

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// Login returns a token
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	l, err := GetList()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(l); err != nil {
		log.WithFields(log.Fields{
			"payload": l,
			"error":   err,
		})
		log.Error("Error encoding system list")
		panic(err)
	}

}

func GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)

	d, err := GetDetails(vars["shortname"])
	if err != nil {
		log.Error(err.Message)
		w.WriteHeader(err.Code)
		return
	}
	w.WriteHeader(http.StatusOK)

	if e := json.NewEncoder(w).Encode(d); e != nil {
		log.WithFields(log.Fields{
			"payload": d,
			"error":   e,
		}).Error("Error encoding system details")
	}
}
