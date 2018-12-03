package main

import (
	"fmt"
	"net/http"

	"./collector"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.HandleFunc("/", handleMain)            // handle main logic
	http.Handle("/metrics", promhttp.Handler()) // handle /metrics request

	log.Info("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) //start server
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page")
	collector.fooMetric.inc()

}
