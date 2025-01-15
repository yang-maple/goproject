package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
<<<<<<< HEAD
	corpid  = "wwb038eade9f4772c4"                          //企业ID
	agentId = "1000003"                                     //应用ID
	secret  = "Oj3RnbV8SaJCAgadFfL8NAEADigOOi-qT4NqzoT_X8U" //Secret
	webhook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=d9956c6a-1f7d-45bc-9667-2d7f42241443"
=======
	corpid  = "" //企业ID
	agentId = "" //应用ID
	secret  = "" //Secret
	webhook = ""
>>>>>>> 4622a0a (2025-1-15)
)

// 企业微信应用消息提醒方法如下
func SendCardMsg(ToUsers interface{}, title, description, url string) (map[string]interface{}, error) {
	btntxt := "详情" //可自定义卡片底下的文字

	qyurl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpid, secret)
	data, err := httpGetJson(qyurl)
	if err != nil {
		log.Println(err)
		return data, err
	}
	errcode := data["errcode"].(float64)
	if errcode != 0 {
		log.Println(errcode)
		return make(map[string]interface{}), nil
	}
	access_token := data["access_token"]
	//卡片内容，不同类型消息通知以下内容需修改(可参考企业微信开发文档)
	req := map[string]interface{}{
		"touser":  ToUsers,
		"msgtype": "textcard",
		"agentid": agentId,
		"textcard": map[string]interface{}{
			"title":       title,
			"description": description,
			"url":         url,
			"btntext":     btntxt,
		},
	}

	sendurl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", access_token)
	data, err = httpPostJson(sendurl, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data, nil
}

// 封装了http请求的GET和POST方法，其返回值都是map[string]interface{}，方便我们使用。
func httpGetJson(url string) (map[string]interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func httpPostJson(url string, data map[string]interface{}) (map[string]interface{}, error) {
	res, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(res))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data2 map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data2)
	if err != nil {
		return nil, err
	}
	return data2, nil
}

func werebot() {
	//touser := []string{"xiongyang", "@all"}
	var (
		id        = 1
		alertname = "test1"
		//time      = "2022-12-1"
		//ip        = "10.0.0.0"
		//job       = "test1"
		//describe  = "this is test"
	)
	content := fmt.Sprintf("title<font color=\\\"warning\\\">%s</font> \n >id:<font color=\\\"comment\\\">%d</font> \n >alertname:<font color=\\\"comment\\\">%s</font>", alertname, id, alertname)

	req := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": content,
		},
	}

	res, err := json.Marshal(req)
	if err != nil {
		fmt.Println("err", err)
	}
	resp, err := http.Post(webhook, "application/json", bytes.NewReader(res))
	var data2 map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data2)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(data2)

}

func main() {
	//data, err := SendCardMsg("熊洋", "test", "this test", "www.baidu.com")
	//if err == nil {
	//	fmt.Println("data:", data)
	//} else {
	//	fmt.Println("err:", err)
	//}

	werebot()
}
