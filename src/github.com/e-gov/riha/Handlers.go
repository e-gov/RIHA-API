package riha

import (
	"encoding/json"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
	"github.com/e-gov/riha/util"
	"github.com/gorilla/mux"
)

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	l, err := GetList()

	for i := range *l {
		u, _ := url.Parse((*l)[i].Shortname)
		(*l)[i].URI = util.GetAPIBase(Port) + "/" + u.String()
	}

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
