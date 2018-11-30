package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// create gauge metric
var (
	go_random_value = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_random_value",
		Help: "randomly generated value in Go",
	})
)

// create counter metric
var (
	pageHits = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_main_page_hits",
		Help: "The total number of times the main page was hit",
	})
)

func main() {
	log.Printf("main function")
	pageHits.Inc()
	http.HandleFunc("/", handle)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	go_random_value.Set(rand.Float64()) //

}
