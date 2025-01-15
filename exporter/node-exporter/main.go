package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"node-exporter/node"
)

func main() {
	prometheus.MustRegister(node.NewNodeStatusController())
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":19990", nil))
}
