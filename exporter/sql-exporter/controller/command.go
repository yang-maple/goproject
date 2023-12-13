package controller

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type Command struct {
	*Basecontroller
	desc *prometheus.Desc
}

func NewCommand(db *sql.DB) *Command {
	desc := prometheus.NewDesc(prometheus.BuildFQName("mysql", "cmd", "totals"), "counts", []string{"cmd"}, nil)
	return &Command{NewBasecontroller(db), desc}
}

func (c *Command) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *Command) Collect(metric chan<- prometheus.Metric) {
	cmds := []string{"insert", "select", "delete", "update", "replace"}
	for _, cmd := range cmds {
		metric <- prometheus.MustNewConstMetric(
			c.desc,
			prometheus.CounterValue,
			c.Status(fmt.Sprintf("com_%s", cmd)),
			cmd,
		)
	}
}
