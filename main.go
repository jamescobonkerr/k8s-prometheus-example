package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/slow", slowHandler)
	log.Fatal(http.ListenAndServe(":3000", router))
}
