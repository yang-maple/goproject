package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"os"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("无法创建 Docker 客户端：", err)
		os.Exit(1)
	}

	// 设置认证信息
	authConfig := registry.AuthConfig{
		Username:      "admin",
		Password:      "Harbor12345",
		ServerAddress: "http://10.1.131.31",
	}

	//创建一个docker secret
	// 登录 Docker 仓库
	secr, err := cli.SecretList(context.Background(), types.SecretListOptions{})
	if err != nil {
		fmt.Println("无法获取 secr 仓库列表：", err)
		os.Exit(1)
	}
	fmt.Println(secr)
	resp, err := cli.RegistryLogin(context.Background(), authConfig)
	if err != nil {
		fmt.Println("无法登录 Docker 仓库：", err)
		os.Exit(1)
	}
	// 打印登录结果
	fmt.Println("登录结果：", resp.Status)
}
