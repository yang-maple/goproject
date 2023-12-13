package powersh

import (
	"fmt"
	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
	"strings"
)

// cmd命令 定义为常量
const cmd = "Get-DhcpServerv4Scope | ForEach-Object { Get-DhcpServerv4ScopeStatistics -ScopeId $_.ScopeId } | Format-Table"

// 匹配获取的数据生产结构体
type Datadhcp struct {
	ScopeId         string
	Free            float64
	InUse           float64
	PercentageInUse float64
}

func Read(field string) []string {
	//选择一个后台
	back := &backend.Local{}
	// 生成一个new shell
	shell, err := ps.New(back)
	if err != nil {
		fmt.Println("shell err is: ", err)
	}
	//交互命令
	ScopeIdstdout, _, err := shell.Execute(cmd + " " + field + " " + "-AutoSize")

	// 测试 完成的执行命令
	//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	//fmt.Println(cmd + " " + field + " " + "-AutoSize")
	//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	if err != nil {
		fmt.Println("cmd err is: ", err)
	}
	// 数据处理 去除空格并分离字符串
	data := strings.Split(strings.ReplaceAll(ScopeIdstdout, " ", ""), "\r\n")
	// 处理完毕延迟关闭后台
	defer shell.Exit()

	return data
}