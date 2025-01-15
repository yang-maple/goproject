package webmodity

import (
	"encoding/json"
	"fmt"
	"os"
)

// 用户信息的结构体
type Person struct {
	Name  string
	Age   int
	Phone int
	Emil  string
	Addr
}

// 用户信息的结构体
type Addr struct {
	City string
}

// 保存用户信息的切片
var custerslince = make([]*Person, 0, 1024)

// 查询功能
func Findtask() []*Person {
	return custerslince
}

// 新增功能
func Addtask(p Person) {

	//custerslince = append(custerslince, &Person{p.Name, p.Age, p.Phone, p.Emil, Addr{p.City}})
	custerslince = append(custerslince, &Person{p.Name, p.Age, p.Phone, p.Emil, p.Addr})

}

// 删除功能
func Deltask(name string) {
	for i, v := range custerslince {
		if name == v.Name {
			custerslince = append(custerslince[:i], custerslince[i+1:]...)
		}
	}
}

// 更新功能
func Updatatask(p Person) {
	for _, v := range custerslince {
		if p.Name == v.Name {
			v.Name = p.Name
			v.Age = p.Age
			v.Phone = p.Phone
			v.Emil = p.Emil
			v.Addr = p.Addr
		}
	}
}

// 保存功能
func Save(file os.File) {
	// 序列化
	data, _ := json.Marshal(custerslince)
	// 写入数据
	err := os.WriteFile("/root/go/goproject/webcuster/custer.json", data, 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}

}

// 初始化 读取功能
func Read(file os.File) {
	// 读取数据
	data, _ := os.ReadFile("/root/go/goproject/webcuster/custer.json")
	// 反序列化 存入缓存中
	json.Unmarshal(data, &custerslince)

}
