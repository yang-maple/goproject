package notice

import (
	"fmt"
	"github.com/spf13/viper"
	mail "github.com/xhit/go-simple-mail/v2"
	"html/template"
	"log"
	"path/filepath"
	"prometheus/pstruct"
	"strings"
	"time"
)

//var htmlBody = `<html>
//	<head>
//		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
//		<title>告警</title>
//	</head>
//	<body>
//		<p>description: "10.1.131.32 监控程序异常"</p>
//		<p><img src="cid:Gopher.png" alt="Go gopher" /></p>
//		<p>Image created by Renee French</p>
//	</body>
//</html>`

func FormatEmailBody(path string, data interface{}) string {
	builder := &strings.Builder{}
	funcs := map[string]interface{}{
		"dateformat": func(t *time.Time) string {
			if t == nil {
				return ""
			}
			return t.Format("2006-01-02 15:04:05")
		},
	}
	tpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles(path))
	tpl.ExecuteTemplate(builder, filepath.Base(path), data)
	return builder.String()
}

type sendinfo struct {
	Host     string
	Username string
	Password string
	Address  string
	Port     int
	Raddress string
}

//var receive []string

func readyaml() sendinfo {
	viper.AddConfigPath(".\\conf")
	viper.SetConfigName("notice")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到文件")
		} else {
			fmt.Println("解析出错", err)
		}
	}

	return sendinfo{viper.GetString("smtp.Host"), viper.GetString("smtp.Username"), viper.GetString("smtp.Password"), viper.GetString("smtp.Address"), viper.GetInt("smtp.Port"), viper.GetString("notice.raddress")}

}

func Emails(v pstruct.Alerts, content string) {

	sinfo := readyaml()
	fmt.Println(sinfo)
	server := mail.NewSMTPClient()

	//SMTP Server
	server.Host = sinfo.Host
	server.Port = sinfo.Port
	server.Username = sinfo.Username
	server.Password = sinfo.Password

	server.KeepAlive = false

	server.ConnectTimeout = 10 * time.Second

	server.SendTimeout = 10 * time.Second

	//SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	//New email simple html with inline and CC
	email := mail.NewMSG()

	email.SetFrom(sinfo.Address).
		AddTo(sinfo.Raddress).
		//AddCc("otherto@example.com").
		SetSubject(v.Alertname)

	email.SetBody(mail.TextHTML, content)
	err = email.Send(smtpClient)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}
}
