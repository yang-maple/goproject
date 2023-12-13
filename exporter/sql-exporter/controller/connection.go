package controller

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type Concontroller struct {
	*Basecontroller
	Maxconnet *prometheus.Desc
	Connetion *prometheus.Desc
}

func NewConcontroller(db *sql.DB) *Concontroller {
	maxdesc := prometheus.NewDesc("max_connection", "max", nil, nil)
	connecting := prometheus.NewDesc("conntecting", "ing", nil, nil)
	return &Concontroller{NewBasecontroller(db), maxdesc, connecting}
}

func (c *Concontroller) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.Maxconnet
	descs <- c.Connetion
}

func (c *Concontroller) Collect(metrics chan<- prometheus.Metric) {
	metrics <- prometheus.MustNewConstMetric(c.Maxconnet, prometheus.GaugeValue, c.Variables("max_connections"))
	metrics <- prometheus.MustNewConstMetric(c.Connetion, prometheus.GaugeValue, c.Status("threads_connected"))
}
