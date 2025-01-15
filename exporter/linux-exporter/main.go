package main

import (
<<<<<<< HEAD
	"goprometheus/linuxinfo"
=======
	"linux-exporter/linuxinfo"
>>>>>>> 4622a0a (2025-1-15)
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheus.MustRegister(linuxinfo.NewCpucontroller())
	prometheus.MustRegister(linuxinfo.NewMemorycontroller())
	prometheus.MustRegister(linuxinfo.NewDiskcontroller())
	prometheus.MustRegister(linuxinfo.NewProcesscontroller())
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":19998", nil))
}
