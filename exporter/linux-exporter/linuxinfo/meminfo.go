package linuxinfo

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/mem"
)

type memInfo struct {
	desc *prometheus.Desc
}

func NewMemorycontroller() *memInfo {
	desc := prometheus.NewDesc(prometheus.BuildFQName("host", "memory", "capability"), "capability", []string{"memory"}, nil)
	return &memInfo{desc}
}

func (c *memInfo) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *memInfo) Collect(metrics chan<- prometheus.Metric) {
	cmds := []string{"total", "used", "free", "buffers", "cached"}
	for _, cmd := range cmds {
		metrics <- prometheus.MustNewConstMetric(
			c.desc,
			prometheus.GaugeValue,
			getmeminfo(cmd),
			cmd,
		)
	}
}

func getmeminfo(name string) float64 {
	v, _ := mem.VirtualMemory()
	gb := float64(1024 * 1024 * 1024)
	switch {
	case name == "total":
		return float64(v.Total) / gb
	case name == "used":
		return float64(v.Used) / gb
	case name == "free":
		return float64(v.Free) / gb
	case name == "buffers":
		return float64(v.Buffers) / gb
	default:
		return float64(v.Cached) / gb
	}

}
