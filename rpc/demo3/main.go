package main

import (
	"demo3/protoc"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	// 引入并创建对象
	var user = &protoc.UserInfo{
		UserName: "张三",
		Age:      18,
		Hobby:    []string{"足球", "篮球"},
		Sex:      protoc.Sex_Man,
		Food: &protoc.Food{
			Name: "大餐",
		},
		Address: map[string]*protoc.UserInfo_Address{
			"家": {
				Province: "北京",
				City:     "昌平区",
			},
		},
		Score: map[string]int32{
			"语文": 98,
		},
	}
	// 输出对象
	fmt.Println(user)

	// 序列化
	data, _ := proto.Marshal(user)
	fmt.Println(data)

	// 反序列化
	users := &protoc.UserInfo{}
	_ = proto.Unmarshal(data, users)
	fmt.Println(user)
}
