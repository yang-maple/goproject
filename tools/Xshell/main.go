package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"golang.org/x/crypto/ssh"
)

type Data struct {
	IP string `json:"ip"`
	OS string `json:"os"`
}

func main() {
	start := time.Now()
	// 建立 SSH 客户端配置
	config := &ssh.ClientConfig{
		User: "noahss", // 替换为你的用户名
		Auth: []ssh.AuthMethod{
			ssh.Password("Noah8u3$4!Q8as"), // 替换为你的密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 不检查 host key
	}

	prometheusClient, err := api.NewClient(api.Config{
		//Address: "http://monitor.i.noahgroup.com/",
		Address: "http://monitor.test.noahgrouptest.com/",
	})
	if err != nil {
		panic(err)
	}
	prometheusApi := v1.NewAPI(prometheusClient)

	// 构建查询参数
	query := "count by (ip) (netdata_info{job=\"linux7-node-exporter\"})"
	query2 := "count by (ip)(node_exporter_build_info{job=\"linux6-node-exporter\"})"
	datas := []Data{}
	os := []Data{}
	// 执行查询
	result, warnings, err := prometheusApi.Query(context.Background(), query, time.Now())
	if err != nil {
		panic(err)
	}
	if len(warnings) > 0 {
		fmt.Println("Warnings:", warnings)
	}
	vector, ok := result.(model.Vector)
	if !ok {
		panic("query result is not a vector")
	}

	// 拿出所有 ip
	for _, sample := range vector {
		modified := strings.Replace(sample.Metric.String(), "=", ":", -1)
		modified = strings.Replace(modified, "ip", `"ip"`, -1)
		var data Data
		json.Unmarshal([]byte(modified), &data)
		datas = append(datas, data)
	}

	result2, warnings, err := prometheusApi.Query(context.Background(), query2, time.Now())
	if err != nil {
		panic(err)
	}
	if len(warnings) > 0 {
		fmt.Println("Warnings:", warnings)
	}
	vector2, ok := result2.(model.Vector)
	if !ok {
		panic("query2 result is not a vector")
	}

	// 拿出所有 ip
	for _, sample := range vector2 {
		modified := strings.Replace(sample.Metric.String(), "=", ":", -1)
		modified = strings.Replace(modified, "ip", `"ip"`, -1)
		var data Data
		json.Unmarshal([]byte(modified), &data)
		datas = append(datas, data)
	}

	var wg sync.WaitGroup
	// 循环建立 SSH 连接
	for _, data := range datas {
		wg.Add(1)
		go func(data Data) {
			defer wg.Done()
			// 建立 SSH 连接
			sshClient, err := ssh.Dial("tcp", data.IP+":22", config)
			if err != nil {
				os = append(os, Data{
					IP: data.IP,
					OS: "dial tcp error",
				})
				fmt.Println(err)
				return
			}
			session, err := sshClient.NewSession()
			if err != nil {
				os = append(os, Data{
					IP: data.IP,
					OS: "session error",
				})
				fmt.Println(err)
				return
			}
			defer sshClient.Close()
			defer session.Close()
			var b bytes.Buffer
			session.Stdout = &b
			if err := session.Run("cat /etc/redhat-release"); err != nil {
				os = append(os, Data{
					IP: data.IP,
					OS: "cmd error",
				})
				return
				//log.Fatalf("Failed to run command: %s", err)
			}
			os = append(os, Data{
				IP: data.IP,
				OS: b.String(),
			})
		}(data)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(len(os))
	fmt.Println(len(datas))
	createexcel(os)
}

func createexcel(data []Data) error {
	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	// 将第一个工作表的名称定为"sheet"
	sheet := f.GetSheetName(0)

	// 设置表头
	f.SetCellValue(sheet, "A1", "IP")
	f.SetCellValue(sheet, "B1", "OS")

	// 创建一个map，key 类型为 string (对应 group 字段),value 类型为 int (对应 Excel 表格中的行号)
	groupOperatorMap := make(map[string]int)

	// 设置行数，从2开始，因为第一行是表头
	i := 2

	// 遍历原数据
	for _, value := range data {
		//判断当前 group 是否已经存在于 groupOperatorMap 中，并将 group 对应的行号存储在 idx 变量中
		if idx, ok := groupOperatorMap[value.IP]; ok {
			f.SetCellValue(sheet, fmt.Sprintf("C%d", idx), value.OS) //如果存在，则更新对应行的 Operator 值
		} else { //如果不存在，则插入新行
			f.SetCellValue(sheet, fmt.Sprintf("A%d", i), value.IP) //fmt.Sprintf("A%d", i) 根据i值生成excel单元格的值，如"A2"
			f.SetCellValue(sheet, fmt.Sprintf("B%d", i), value.OS)
			//将当前 group 和对应的行号存储到 groupOperatorMap 中,并将行号递增
			groupOperatorMap[value.IP] = i
			i++
		}
	}

	// 保存 Excel 文件,默认存储当前目录下
	return f.SaveAs("output.xlsx")
}
