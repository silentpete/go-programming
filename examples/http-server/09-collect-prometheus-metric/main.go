package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	namespace = "metric_exporter"
)

func main() {
	var (
		listenAddress = kingpin.Flag("address", "Address to listen on for web requests.").Default(":9811").String()
		metricsPath   = kingpin.Flag("metrics-path", "Relative path to metrics.").Default("/metrics").String()
		// collectTimeout = kingpin.Flag("collect.timeout", "Timeout for trying to get stats from remove host.").Default("5s").Duration()
		metricUp = prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "up"), "Was the last scrape of the metric successful.", nil, nil)
		// thing    = prometheus.NewDesc("metric_name", "This is the HELP string.", nil, nil)
	)

	log.AddFlags(kingpin.CommandLine)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.Infoln("Starting metric_exporter")

	prometheus.MustRegister(metricUp)

	log.Infoln("Listening on", *listenAddress)
	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/metrics", http.StatusSeeOther)
	})
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
