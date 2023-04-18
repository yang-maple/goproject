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

func cmds(conn net.Conn) (string, []string) {
	reader := bufio.NewReader(conn)
	op, err := reader.ReadString('|')
	if err != nil {
		return "", nil
	}
	op = strings.Trim(op, "|")
	fmt.Println("op", op)
	count, _ := reader.ReadString('|')

	counts, _ := strconv.Atoi(strings.Trim(count, "|"))
	fmt.Println("countes:", counts)
	agrs := make([]string, 0, counts)
	for counts > 0 {
		param, err := reader.ReadString('|')
		if err != nil {
			fmt.Println("param err ", err)
		}
		agrs = append(agrs, strings.Trim(param, "|"))
		counts--
	}
	fmt.Println("args:", agrs)
	return op, agrs
}
func ls(conn net.Conn) {

	dir, err := os.ReadDir(".")
	if err != nil {
		log.Println(err)
	}
	dirname := make([]string, 0, len(dir))
	for i := 0; i < len(dir); i++ {
		dirname = append(dirname, dir[i].Name())
	}
	//fmt.Println(dirname)
	names := strings.Join(dirname, ":")
	//fmt.Println(names)
	suffix := ""
	if len(dir) > 0 {
		suffix = ":"
	}
	fmt.Fprintf(conn, "%d|%s%s", len(dir), names, suffix)

}

func cat(conn net.Conn, filename []string) {

	file, err := os.Open(filename[0])
	if err != nil {
		fmt.Fprint(conn, "0|")
	}
	ctx := make([]byte, 2099)
	n, _ := file.Read(ctx)
	fmt.Println("n:", n)
	fmt.Println("ctx:", ctx)
	fmt.Fprintf(conn, "%d|%s", n, string(ctx[:n]))
	defer file.Close()
}

func HandleConn(conn net.Conn) {
	//读取消息
	defer conn.Close()
END:
	for {
		op, agrs := cmds(conn)
		if op == "" && agrs == nil {
			fmt.Printf("客户%s 已退出\n", conn.RemoteAddr())
			break END
		}
		switch op {
		case "ls":
			ls(conn)
		case "cat":
			cat(conn, agrs)
		case "quit":
			fmt.Printf("客户%s 已退出\n", conn.RemoteAddr())
			break END
		case "":
		default:
			fmt.Println("指令错误")
		}
	}

}

func main() {

	logfile, err := os.OpenFile("filesever.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	addr := "0.0.0.0:9999"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	log.Printf("listening on %s\n", addr)
	defer listen.Close()
	func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				log.Println(err)
			}
			log.Printf("connected on %s\n", conn.RemoteAddr())
			go HandleConn(conn)
		}
	}()

}
