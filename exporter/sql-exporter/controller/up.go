package controller

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type Upcontroller struct {
	*Basecontroller
	desc *prometheus.Desc
}

func NewUpcontroller(db *sql.DB) *Upcontroller {
	desc := prometheus.NewDesc("mysql_up", "mysql health", nil, nil)
	return &Upcontroller{NewBasecontroller(db), desc}
}

func (c *Upcontroller) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *Upcontroller) Collect(metrics chan<- prometheus.Metric) {
	up := 1
	if err := c.db.Ping(); err != nil {
		up = 0
	}
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue, float64(up))
}
