package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// 创建带 house 和 room 标签的 gauge 指标对象
	temp := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "home_temperature_celsius",
			Help: "The current temperature in degrees Celsius.",
		},
		// 指定标签名称
		[]string{"house", "room"},
	)

	// 注册到全局默认注册表中
	prometheus.MustRegister(temp)

	// 针对不同标签值设置不同的指标值
	temp.WithLabelValues("cnych", "living-room").Set(27)
	temp.WithLabelValues("cnych", "bedroom").Set(25.3)
	temp.WithLabelValues("ydzs", "living-room").Set(24.5)
	temp.WithLabelValues("ydzs", "bedroom").Set(27.7)

	// 暴露自定义的指标
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
