package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func cat(conn net.Conn, filename string) {
	fmt.Fprintf(conn, "cat|1|%s|", filename)

	reader := bufio.NewReader(conn)
	lines, err := reader.ReadString('|')
	if err != nil {
		fmt.Println(err)
	}
	size, err := strconv.Atoi(strings.Trim(lines, "|"))
	fmt.Println("size:", size)
	if err != nil {
		log.Println(err)
	}

	if size > 0 {
		ctx := make([]byte, size)
		n, err := reader.Read(ctx)
		fmt.Println("n:", n)
		if err != nil {
			fmt.Println("read err", err)
		}
		fmt.Printf("文件内容为:%s\n", string(ctx[:n]))
		size--
	} else {

		fmt.Println("文件不存在")

	}

}

func ls(conn net.Conn) {
	//发送消息
	fmt.Fprintf(conn, "ls|0|")

	//读取消息
	reader := bufio.NewReader(conn)
	lines, err := reader.ReadString('|')
	if err != nil {
		fmt.Println(err)
	}
	size, err := strconv.Atoi(strings.Trim(lines, "|"))
	fmt.Println("size:", size)
	if err != nil {
		fmt.Println(err)
	}
	for size > 0 {
		name, _ := reader.ReadString(':')
		fmt.Println(strings.Trim(name, ":"))
		size--
	}

}

func main() {

	logfile, err := os.OpenFile("filecuster.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	log.Printf("conntected on %s\n", addr)
quit:
	for {
		fmt.Print("请输入指令：")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		cmds := strings.Split(input, " ")
		switch cmds[0] {
		case "ls":
			ls(conn)
		case "cat":
			cat(conn, cmds[1])
		case "quit":
			fmt.Fprintf(conn, "quit|0|")
			break quit
		default:
			fmt.Println("输入格式错误")
		}
	}
	fmt.Println("成功退出")

}
