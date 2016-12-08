package main

import (
	"flag"
	"fmt"
	"net/http"

	riha "github.com/e-gov/riha"
	util "github.com/e-gov/riha/util"

	log "github.com/Sirupsen/logrus"
)

func main() {
	var port = flag.Int("port", 8090, "Port to bind to on the localhost interface")

	flag.Parse()

	util.SetupSvcLogging()

	router := riha.NewRouter()
	log.Infof("Starting a server on localhost:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
