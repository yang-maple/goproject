package main

import (
	"dhcpexporter/powersh"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 自定义接口 注册信息 包括监听的地址和端口
var addr = flag.String("listen-address", ":19998", "The address to listen on for HTTP requests.")

var (
	// 创建带标签的 gauge 模板
	opsid = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "windows_dhcppool",                       // 监控指标的Name
			Help: "dhcppool in Free,InUse,PercentageInUse", // 注释字段
		},
		// 指定标签名称
		[]string{"ip", "chart"},
	)
)

func init() {
	// 注册到全局默认注册表中
	prometheus.MustRegister(opsid)
}

// main 方法
func main() {
	// 解析自定义接口 注册信息  即 addr
	flag.Parse()
	go func() {
		for {
			// 定义结构体
			dhcp := powersh.Datadhcp{}
			// 调用powersh 包中的read 函数 获取的 数据值保存在 变量中
			dataScopeId := powersh.Read("ScopeId")
			dataFree := powersh.Read("Free")
			dataInUse := powersh.Read("InUse")
			dataPercentageInUse := powersh.Read("PercentageInUse")
			// 前三行为无用数据过滤
			for i := 3; i < len(dataScopeId)-1; i++ {
				//将切片的值赋值给结构体
				// 数据类型转化为 float64
				dhcp.Free, _ = strconv.ParseFloat(dataFree[i], 64)
				dhcp.InUse, _ = strconv.ParseFloat(dataInUse[i], 64)
				dhcp.PercentageInUse, _ = strconv.ParseFloat(dataPercentageInUse[i], 64)
				// 数据渲染到gauge模板上
				opsid.WithLabelValues(dataScopeId[i], "Free").Set(dhcp.Free)
				opsid.WithLabelValues(dataScopeId[i], "InUse").Set(dhcp.InUse)
				opsid.WithLabelValues(dataScopeId[i], "PercentageInUse").Set(dhcp.PercentageInUse)
			}
			// 十分钟获取一次数据
			time.Sleep(time.Minute * 10)
		}
	}()
	//  暴露自定义的指标 访问/api地址
	http.Handle("/metrics", promhttp.Handler())
	// 启动监听服务，当有错误时终止程序输出日志
	log.Fatal(http.ListenAndServe(*addr, nil))

}
