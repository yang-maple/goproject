package main

import (
	"bufio"
	"fmt"
	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
	"github.com/bhendo/go-powershell/middleware"
	"io"
	"os"
	"powershell_remote/prometheus_info"
	"strings"
)

func main() {
	list := prometheus_info.GetInfo()
	newList := make(map[string]string)
	//打开并按行读取文件
	files, _ := os.Open("./aaa.txt")
	reader := bufio.NewReader(files)
	for {
		strs, err := reader.ReadString('\n') // 读取到换行符为止，读取内容包括换行符

		//判断字典中存在的主机名
		if value, ok := list[strings.Trim(strs, "\r\n")]; ok {
			// 如果 key 存在，更新到新的字典
			newList[strings.Trim(strs, "\r\n")] = value
		}
		if err == io.EOF { //io.EOF 读取到了文件的末尾
			break
		}
	}
	defer files.Close()
	//挑选一个后台
	back := &backend.Local{}

	// 开启一个本地 powershell 进程
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	// 准备本地文件
	lfile, err := os.OpenFile("./info.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	// 准备远程会话配置
	config := middleware.NewSessionConfig()
	var session middleware.Middleware
	i := 0
	for k, v := range newList {
		config.ComputerName = v
		session, err = middleware.NewSession(shell, config)
		if err != nil {
			fmt.Println("connect failed: ", k)
			continue
		}
		stdout, _, err := session.Execute("Get-LocalGroupMember Administrators | Select-Object -ExpandProperty Name")
		if err != nil {
			fmt.Println("cmd failed: ", k)
			continue
		}
		_, _ = lfile.Write([]byte(k + "\r\n" + stdout + "\r\n"))
		_, _, _ = session.Execute("exit")
		i++
	}
	fmt.Println(i)
	_ = lfile.Close()
	defer session.Exit()
	////config.ComputerName = "HQ-T-NGSCCM01"
	//
	//// 通过包装已存在的会话中间件来创建一个新的 shell
	////session, err := middleware.NewSession(shell, config)
	////if err != nil {
	////	panic(err)
	////}
	////defer session.Exit() //也将关闭底层的 ps shell！
	//
	//// everything run via the session is run on the remote machine
	//// 会话相关的事情都是运行在远程及其上的
	////stdout, _, err := session.Execute("$env:COMPUTERNAME")
	////if err != nil {
	////	panic(err)
	////}
	////fmt.Println(stdout)
}
