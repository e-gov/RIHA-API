package riha

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	util "github.com/e-gov/riha/util"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		handler = util.LoggingHandler(handler)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		log.WithFields(log.Fields{
			"path":   route.Pattern,
			"method": route.Method,
		}).Infof("Added route %s", route.String())
	}
	return router
}
