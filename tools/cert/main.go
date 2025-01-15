package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type KubeConfig struct {
	ApiVersion string   `json:"apiVersion"`
	Clusters   clusters `json:"clusters"`
	Users      users    `json:"users"`
}

type clusters struct {
	Cluster cluster `json:"cluster"`
}
type cluster struct {
	Certificate string `json:"certificate-authority-data"`
	Server      string `json:"server"`
}

type users struct {
	Name string `json:"name"`
	User user   `json:"user"`
}

type user struct {
	ClientCertificate string `json:"client-certificate-data"`
	ClientKey         string `json:"client-key-data"`
}

func main() {

	files, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("读取文件失败")
	}
	config := KubeConfig{}
	err = yaml.Unmarshal(files, config)
	if err != nil {
		fmt.Println("解析失败")
	}
	fmt.Println(config)
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./")
	//err := viper.ReadInConfig()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//data := viper.Get("users.name")
	//fmt.Println(reflect.TypeOf(data))
}
