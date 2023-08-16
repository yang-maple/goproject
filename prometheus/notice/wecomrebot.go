package notice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"prometheus/pstruct"

	"github.com/spf13/viper"
)

func readwebhook() string {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("notice")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到文件")
		} else {
			fmt.Println("解析出错", err)
		}
	}

	return viper.GetString("webhook.url")
}

func WeRebot(alert *pstruct.Alerts) {
	content := fmt.Sprintf("#事件<font color=\\\"warning\\\">%s</font> \n >ip:<font color=\\\"comment\\\">%s</font> \n >Description:<font color=\\\"comment\\\">%s</font>",
		alert.Alertname,
		alert.Ip,
		alert.Description,
	)
	touser := []string{"xiongyang", "@all"}
	req := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content":        content,
			"mentioned_list": touser,
		},
	}

	res, err := json.Marshal(req)
	if err != nil {
		fmt.Println("err", err)
	}
	resp, err := http.Post(readwebhook(), "application/json", bytes.NewReader(res))
	var data2 map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data2)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(data2)

}
