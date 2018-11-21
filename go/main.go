package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	projectID = "yuri-next2019"
)

func main() {

	log.Printf("main function")
	http.HandleFunc("/", handle)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[golang-profiler:handle] Entered")
	fmt.Fprintln(w, "hello!")
	log.Printf("[golang-profiler:handle] Exited")
}
