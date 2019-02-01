package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// create gauge metric
var (
	goRandomValue = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_random_value",
		Help: "randomly generated value in Go",
	})
)

func main() {
	log.Printf("main function")
	http.HandleFunc("/", handle)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	goRandomValue.Set(rand.Float64()) //
	fmt.Fprintf(w, "Hello")
}
