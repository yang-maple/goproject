package linuxinfo

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/disk"
)

type diskInfo struct {
	desc *prometheus.Desc
}

func NewDiskcontroller() *diskInfo {
	desc := prometheus.NewDesc(prometheus.BuildFQName("host", "disk", "dir"), "df_h", []string{"device", "status"}, nil)
	return &diskInfo{desc}
}

func (c *diskInfo) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *diskInfo) Collect(metrics chan<- prometheus.Metric) {
	part, _ := disk.Partitions(true)
	diskdir := uniqueParts(part)
	status := []string{"total", "used", "free"}
	for _, dir := range diskdir {
		for _, dirstatus := range status {
			value := getdiskinfo(dir.Mountpoint, dirstatus)
			metrics <- prometheus.MustNewConstMetric(
				c.desc,
				prometheus.GaugeValue,
				value,
				dir.Mountpoint,
				dirstatus,
			)
		}
	}
}

func getdiskinfo(name string, status string) float64 {
	d, _ := disk.Usage(name)
	gb := float64(1024 * 1024 * 1024)
	switch {
	case status == "total":
		return float64(d.Total) / gb
	case status == "used":
		return float64(d.Used) / gb
	default:
		return float64(d.Free) / gb
	}
}

func uniqueParts(s []disk.PartitionStat) []disk.PartitionStat {
	result := make([]disk.PartitionStat, 0)
	m := make(map[string]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v.Mountpoint]; !ok {
			result = append(result, v)
			m[v.Mountpoint] = true
		}
	}
	return result
}
