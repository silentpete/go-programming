package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	totalRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "helloworld_total_requests",
		Help: "The total count of page requests to /* sinse server started.",
	})
	iconRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "helloworld_total_icon_requests",
		Help: "The total count of page requests to /* sinse server started.",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(iconRequests)

}

func handler(w http.ResponseWriter, r *http.Request) {
	// increment the page hits when someone requests the /
	totalRequests.Inc()
	// print out to the browser Hello World!
	fmt.Fprintf(w, "Hello World!")
	// log the request URI to the server stdout
	log.Println(r.RequestURI)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	totalRequests.Inc()
	iconRequests.Inc()
	log.Println("icon was requested by:", r.RemoteAddr)
	fmt.Fprintf(w, "")
}

func main() {
	log.Println("start of main...")
	// add a handler for / to the default http server
	http.HandleFunc("/", handler)
	// because browsers request /favicon.ico, we add a handler so our metrics don't get false calls
	http.HandleFunc("/favicon.ico", faviconHandler)
	// add the /metrics handler, notice we use prometheus's http server
	http.Handle("/metrics", promhttp.Handler())
	// start the server
	err := http.ListenAndServe("0.0.0.0:6060", nil)
	if err != nil {
		fmt.Println(err)
	}
}
