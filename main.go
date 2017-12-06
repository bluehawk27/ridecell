package main

import (
	"fmt"
	"net/http"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/bluehawk27/ridecell/httpi"
	"github.com/gorilla/mux"
)

func main() {
	// todo move to config
	addr := "127.0.0.1:8080"

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		r := mux.NewRouter()
		log.Info("Running HTTP Server on: " + addr)
		r.HandleFunc("/list/", httpi.List).Methods("GET")
		http.Handle("/", r)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(fmt.Sprintf("Fatal error server.Serve: %s \n", err))
		}
	}()
	wg.Wait()
}
