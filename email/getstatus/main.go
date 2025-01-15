package main

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"os"
	"strings"
)

type Info struct {
	NextStart string `json:"NextStart"`
	RequestId string `json:"RequestId"`
	Data      struct {
		MailDetail []MailDetail `json:"mailDetail"`
	} `json:"data"`
}

type MailDetail struct {
	AccountName         string `json:"AccountName"`
	Subject             string `json:"Subject"`
	Status              int    `json:"Status"`
	UtcLastUpdateTime   string `json:"UtcLastUpdateTime"`
	ErrorClassification string `json:"ErrorClassification"`
	Message             string `json:"Message"`
	ToAddress           string `json:"ToAddress"`
	LastUpdateTime      string `json:"LastUpdateTime"`
}

// Description:
//
// 使用AK&SK初始化账号Client
//
// @return Client
//
// @throws Exception
func CreateClient() (_result *dm20151123.Client, _err error) {
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	// 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID:。
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET:。
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dm
	config.Endpoint = tea.String("dm.aliyuncs.com")
	_result = &dm20151123.Client{}
	_result, _err = dm20151123.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	senderStatisticsDetailByParamRequest := &dm20151123.SenderStatisticsDetailByParamRequest{
		//StartTime: tea.String("2024-09-21 08:15"),
		//EndTime:   tea.String("2024-09-29 18:35"),
		//NextStart: tea.String("2e22358848#201#1727540277#nbedk1@163.com.600000116293235049"),
		TagName: tea.String("aa"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		result, _err := client.SenderStatisticsDetailByParamWithOptions(senderStatisticsDetailByParamRequest, runtime)
		if _err != nil {
			return _err
		}
		var info Info

		_err = json.Unmarshal([]byte(result.Body.String()), &info)
		fmt.Println(info)

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}

// data:{
//metail: [
//{
//]
//}
