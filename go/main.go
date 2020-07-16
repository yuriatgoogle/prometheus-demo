package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// create gauge metric
var (
	// total requests counter
	goRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_requests",
		Help: "total go requests",
	})
	// total failed requests counter
	goFailedRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_failed_requests",
		Help: "total failed go requests",
	})
	// go successful request latency histogram
	goSuccessLatencyHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "go_success_request_latency",
		Help:    "go successful request latency by path",
		Buckets: []float64{100, 400},
	})
	// go failed request latency histogram
	goFailedLatencyHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "go_failed_request_latency",
		Help:    "go failed request latency by path",
		Buckets: []float64{100, 400},
	})
)

func main() {
	log.Printf("main function")
	http.HandleFunc("/", handle)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// record incoming request time
	requestReceivedTime := time.Now()
	// increment total request counter
	goRequestCounter.Inc()
	rand.Seed(time.Now().Unix())

	// sleep randomly before failing
	if rand.Intn(10) == 1 {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		requestFailedLatency := time.Now().Sub(requestReceivedTime)
		// increment failed request counter
		goFailedRequestCounter.Inc()
		// record failed latency
		goFailedLatencyHistogram.Observe(float64(requestFailedLatency))

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed in %v ms", requestFailedLatency.Milliseconds())
	} else {
		// sleep randomly before being successful
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))

		// collect latency histogram
		requestSuccessLatency := time.Now().Sub(requestReceivedTime)
		goSuccessLatencyHistogram.Observe(float64(requestSuccessLatency))

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "success in %v ms", requestSuccessLatency.Milliseconds())
	}
}
