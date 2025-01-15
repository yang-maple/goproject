package main

import (
	"fmt"
	"log"
)

// 主函数
func main() {

	conn, err := MyClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal("连接失败", err)
	}
	ret := 0
	err2 := conn.Perimeter(Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal("err2:", err2)
	}
	fmt.Println("周长：", ret)

	err3 := conn.Area(Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal("err3:", err3)
	}
	fmt.Println("面积：", ret)

}
