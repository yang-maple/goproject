package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	adjustedTime := time.Now().Add(time.Hour * -8).Format("2006-01-02T15:04:05Z")
	fmt.Println("当前时间：", now)
	fmt.Println("减去 8 小时后的时间：", adjustedTime)
	adjustedTime2 := time.Now().Add(time.Hour * -9).Format("2006-01-02T15:04:05Z")
	fmt.Println("减去 9 小时后的时间：", adjustedTime2)
	adjustedTime3 := time.Now().Add(time.Hour * -10).Format("2006-01-02T15:04:05Z")
	fmt.Println("减去 10 小时后的时间：", adjustedTime3)

}
