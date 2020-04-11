package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "go_service_http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
)

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(prometheusMiddleware)
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/slow", slowHandler)
	log.Fatal(http.ListenAndServe(":3000", router))
}
