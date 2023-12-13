package linuxinfo

import (
	"os/exec"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type processInfo struct {
	desc *prometheus.Desc
}

func NewProcesscontroller() *processInfo {
	desc := prometheus.NewDesc(prometheus.BuildFQName("host", "process", "capability"), "capability", []string{"port"}, nil)
	return &processInfo{desc}
}

func (c *processInfo) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *processInfo) Collect(metrics chan<- prometheus.Metric) {
	value := getportstatus()
	for _, port := range value {
		metrics <- prometheus.MustNewConstMetric(
			c.desc,
			prometheus.GaugeValue,
			1,
			port,
		)
	}

}

func getportstatus() []string {
	cmd := exec.Command("/bin/bash", "-c", `netstat -lntp | egrep -v "tcp6" |egrep -v "127.0.0.1" |awk -F ' '  '{print $4}' |awk -F ':' '{print $2}' |egrep -v '^$'`)
	out, _ := cmd.CombinedOutput()
	lines := strings.Split(string(out), "\n")
	port := lines[0:(len(lines) - 1)]
	return port
}
