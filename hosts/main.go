package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type node struct {
	ID      string        `json:"ID"`
	Service string        `json:"Service"`
	Tags    []interface{} `json:"Tags"`
	Meta    struct {
		Area     string `json:"area"`
		Hostname string `json:"hostname"`
	} `json:"Meta"`
	Port            int    `json:"Port"`
	Address         string `json:"Address"`
	TaggedAddresses struct {
		LanIpv4 struct {
			Address string `json:"Address"`
			Port    int    `json:"Port"`
		} `json:"lan_ipv4"`
		WanIpv4 struct {
			Address string `json:"Address"`
			Port    int    `json:"Port"`
		} `json:"wan_ipv4"`
	} `json:"TaggedAddresses"`
	Weights struct {
		Passing int `json:"Passing"`
		Warning int `json:"Warning"`
	} `json:"Weights"`
	EnableTagOverride bool   `json:"EnableTagOverride"`
	Datacenter        string `json:"Datacenter"`
}

<<<<<<< HEAD
func Readall(lfile *os.File, wfile *os.File) {
=======
func Readall(lfile *os.File, wfile *os.File, l6file *os.File) {
>>>>>>> 4622a0a (2025-1-15)
	f1, _ := os.ReadFile("./services.json")
	var R1 []map[string]node
	err := json.Unmarshal(f1, &R1)
	if err != nil {
		fmt.Println("读取失败")
		return
	}
	for _, v := range R1[0] {
<<<<<<< HEAD
		if v.Port == 19999 || v.Port == 9100 {
			lfile.Write([]byte(v.Address + "\n"))
=======
		if v.Port == 9100 {
			l6file.Write([]byte(v.Address + "\n"))
>>>>>>> 4622a0a (2025-1-15)
		} else if v.Port == 9182 {
			wfile.Write([]byte(v.Address + "\n"))
		}
	}
}

func main() {
	lfile, err := os.OpenFile("./all_linux.json", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	wfile, err := os.OpenFile("./all_windows.json", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
<<<<<<< HEAD
	if err != nil {
		fmt.Println("读取文件失败")
	}
	Readall(lfile, wfile)
=======
	l6file, err := os.OpenFile("./all_linux6.json", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("读取文件失败")
	}
	Readall(lfile, wfile, l6file)
>>>>>>> 4622a0a (2025-1-15)
}
