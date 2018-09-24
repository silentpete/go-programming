// Reference: https://github.com/prometheus/client_golang/blob/master/examples/simple/main.go
// A minimal example of how to include Prometheus instrumentation.
package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("starting server...")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
