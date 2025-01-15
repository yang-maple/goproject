package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/xuri/excelize/v2"
	"strings"
	"time"
)

type Data struct {
	HostName string `json:"hostname"`
	IP       string `json:"ip"`
	Product  string `json:"product"`
}

func main() {
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
	query := "count by (ip,hostname) (windows_cs_hostname)\n  *\non (ip)\ngroup_left(product)\ncount by (ip, product) (windows_os_info)"

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
	datas := []Data{}
	for _, sample := range vector {
		modified := strings.Replace(sample.Metric.String(), "=", ":", -1)
		modified = strings.Replace(modified, "hostname", `"hostname"`, -1)
		modified = strings.Replace(modified, "ip", `"ip"`, -1)
		modified = strings.Replace(modified, "product", `"product"`, -1)
		fmt.Println(modified)
		var data Data
		json.Unmarshal([]byte(modified), &data)
		datas = append(datas, data)
	}
	err = createexcel(datas)
	if err != nil {
		panic(err.Error())
	}

	//aaa := `{hostname:"AWS-AD-Server01","ip":"10.221.9.119","product":"Microsoft Windows Server 2016 Datacenter"}`
	//var bbb Data
	//json.Unmarshal([]byte(aaa), &bbb)
	//fmt.Println(bbb)
}

func createexcel(data []Data) error {
	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	// 将第一个工作表的名称定为"sheet"
	sheet := f.GetSheetName(0)

	// 设置表头
	f.SetCellValue(sheet, "A1", "HostName")
	f.SetCellValue(sheet, "B1", "IP")
	f.SetCellValue(sheet, "C1", "Product")

	// 创建一个map，key 类型为 string (对应 group 字段),value 类型为 int (对应 Excel 表格中的行号)
	groupOperatorMap := make(map[string]int)

	// 设置行数，从2开始，因为第一行是表头
	i := 2

	// 遍历原数据
	for _, value := range data {
		//判断当前 group 是否已经存在于 groupOperatorMap 中，并将 group 对应的行号存储在 idx 变量中
		if idx, ok := groupOperatorMap[value.IP]; ok {
			f.SetCellValue(sheet, fmt.Sprintf("C%d", idx), value.Product) //如果存在，则更新对应行的 Operator 值
		} else { //如果不存在，则插入新行
			f.SetCellValue(sheet, fmt.Sprintf("A%d", i), value.HostName) //fmt.Sprintf("A%d", i) 根据i值生成excel单元格的值，如"A2"
			f.SetCellValue(sheet, fmt.Sprintf("B%d", i), value.IP)
			f.SetCellValue(sheet, fmt.Sprintf("C%d", i), value.Product)

			//将当前 group 和对应的行号存储到 groupOperatorMap 中,并将行号递增
			groupOperatorMap[value.IP] = i
			i++
		}
	}

	// 保存 Excel 文件,默认存储当前目录下
	return f.SaveAs("output.xlsx")
}
