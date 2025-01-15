package linuxinfo

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

type cpuInfo struct {
	desc *prometheus.Desc
}

func NewCpucontroller() *cpuInfo {
	desc := prometheus.NewDesc(prometheus.BuildFQName("host", "cpu", "capability"), "capability", []string{"cpu"}, nil)
	return &cpuInfo{desc}
}

func (c *cpuInfo) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *cpuInfo) Collect(metrics chan<- prometheus.Metric) {
	cmds := []string{"cores", "used", "load1", "load5", "load15"}
	for _, cmd := range cmds {
		metrics <- prometheus.MustNewConstMetric(
			c.desc,
			prometheus.GaugeValue,
			getcpuinfo(cmd),
			cmd,
		)
	}
}

func getcpuinfo(name string) float64 {
	switch {
	case name == "cores":
		physicalCnt, _ := cpu.Counts(false)
		return float64(physicalCnt)
	case name == "used":
		cpu_used, _ := cpu.Percent(time.Millisecond*200, false)
		return float64(cpu_used[0])
	case name == "load1":
		loads, _ := load.Avg()
		return loads.Load1
	case name == "load2":
		loads, _ := load.Avg()
		return loads.Load5
	default:
		loads, _ := load.Avg()
		return loads.Load15
	}

}
