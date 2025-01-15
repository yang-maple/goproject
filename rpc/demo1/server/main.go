package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Copyright 2019-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");

// Params 定义一个结构体
type Params struct {
	Width, Height int
}

type Rect struct{}

// Area 计算面积的方法
func (*Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

// Perimeter 计算周长的方法
func (*Rect) Perimeter(p Params, ret *int) error {
	*ret = 2 * (p.Height + p.Width)
	return nil
}

func main() {
	// 1.注册服务
	rect := new(Rect)
	// 注册一个rect的服务
	rpc.RegisterName("Rect", rect)

	//rpc.HandleHTTP()
	//// 2.启动服务
	//err := http.ListenAndServe(":1234", nil)
	//if err != nil {
	//	log.Panicln(err)
	//}

	// 2. 创建监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error Starting Server:", err)
	}
	fmt.Println("Listening on 1234...")
	for {
		// 3. 建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting:", err)
		}
		fmt.Println("连接成功", conn.RemoteAddr())
		// 绑定服务
		go rpc.ServeConn(conn)
	}
}
