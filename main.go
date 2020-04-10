package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":3000", router))
}
