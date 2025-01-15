package prometheus_info

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"time"
)

/*
查询主机信息
*/

func GetInfo() map[string]string {
	// 创建一个Prometheus API客户端
	client, err := api.NewClient(api.Config{
		Address: "http://monitor.i.noahgroup.com/",
		//Address: "http://monitor.test.noahgrouptest.com/",
	})
	if err != nil {
		panic(err)
	}
	api := v1.NewAPI(client)

	// 构建查询参数
	query := "count(windows_cs_hostname) by (hostname,ip)"

	// 执行查询
	result, warnings, err := api.Query(context.Background(), query, time.Now())
	if err != nil {
		panic(err)
	}
	if len(warnings) > 0 {
		fmt.Println("Warnings:", warnings)
	}

	// 处理查询结果
	vector, ok := result.(model.Vector)
	if !ok {
		panic("query result is not a vector")
	}
	host := make(map[string]string)
	for _, sample := range vector {
		host[string(sample.Metric["ip"])] = string(sample.Metric["hostname"])
	}
	return host
}
