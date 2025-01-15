package main

import (
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

const htmlBody = `
<html>
<div>亲爱的伙伴，您好!</div>
<div>本月底，您名下将有2位客户的积分到期，建议您提醒客户及时前去微诺亚积分商城进行积分兑换，积分一经逾期无法恢复，请知悉为感</div>
<div>FR端：提醒短信及邮件将于今日陆续发送至FR手机号及邮箱</div>
<div>客户端：提醒短信将于今日下午陆续发送至客户手机</div>
</html>

`

/*
const htmlBody = `
<html>
<div>亲爱的伙伴，您好!</div>
<div>本月底，您名下将有2位客户的积分到期，建议您提醒客户及时前去微诺亚积分商城进行积分兑换，积分一经逾期无法恢复，请知悉为感</div>
<div>FR端：提醒短信及邮件将于今日陆续发送至FR手机号及邮箱</div>
<div>客户端：提醒短信将于今日下午陆续发送至客户手机</div>
<div><div style="padding-left: 2em">到期积分明细如下:</div><div style="padding-left: 2em"><table border="1" cellspacing="0"><tr><th>序号</th><th>会员编号</th><th>集团号</th><th>姓名</th><th>会员等级</th><th>目前积分</th><th>月底即将到期积分</th></tr><tr><td>1</td><td>927787</td><td>000335616</td><td>毛*</td><td>白银</td><td>13</td><td>12</td></tr><tr><td>2</td><td>371856</td><td>000507326</td><td>许*锋</td><td>白银</td><td>2</td><td>2</td></tr></table></div></div>
<div>2024年09月24日Tue</div>
<div>尊敬的投資者：</div>
<div>中秋佳節至，月圓人團圓，我們與您共慶佳節！</div>
<div>福利一：認購任意公募基金（含現金寶），您將有機會獲得費用減免獎賞，最高可達HKD12,688。</div>
<div>福利二：您可享10%的佣金折扣，最高可享50元。</div>
<div>福利二：聯繫您的國際持牌代表獲取818理財節優惠碼，部分產品可專享5折認購費優惠（限美元、港幣份額）。</div>
<div>活動條件與條款適用，具體活動詳情登錄方舟達富APP或點擊https://noahmkt.com/2P7x查看。</div>
<div><img src="./111.png" /> </div>
<div>此邮件由系统发出，请勿回复。如有疑问，请联系区域客经负责人。</div>
</html>

`
*/

//const htmlBody = `尊敬的客户,<br/>
//<br/>
//感谢您的支持和选用供阁下参考。<br/>
//<br/>
//如阁下对此结单有任何疑问，请致电您的理财顾问或<br/>
//<br/>
//感谢您的关注！<br/>
//<br/>
//<br/>
//如无法查阅电邮附件，请按此连接 http://www.adobe.com下载及安装 Adobe Acrobat Reader软件。<br/>
//<br/>
//Dear Customer,<br/>
//<br/>
//Tlease find attached the latest Consolidated Monthly Statement for your information.<br/>
//<br/>
//Should you have any questions about the enclosed Statement, please free feel to contact your Investment Consultant or our <br/>
//<br/>
//Thank you for your kind attention!<br/>
//<br/>
//<br/>
//If you cannot view the email attachment, please click here http://www.adobe.com to download and install Adobe Acrobat Reader.<br/>
//<br/>
//<br/>
//客户如有通过登记的手机号码登录查看 或扫码下载。<br/>
//<br/>
//Log onto our A<br/>`

func main() {
	server := mail.NewSMTPClient()

	//// SMTP Server

	//Office 365
	//server.Host = "smtp-mail.outlook.com"
	//server.Port = 587
	//server.Username = "aws_noah@noahgroup.sg"
	//server.Password = "Noah8u3$4!Q8as"

	//Aliyun
	//server.Host = "smtpdm.aliyun.com"
	//server.Port = 80
	//server.Username = "testxiy@notice.noahgrouptest.sg"
	//server.Password = "QWert12345"
	//sender := "testxiy@notice.noahgrouptest.sg"

	// Tencent
	//server.Host = "smtp.qcloudmail.com"
	//server.Port = 587
	//server.Username = "testxiy@tecent.noahgrouptest.com"
	//server.Password = "QWert12345"

	// Noah
	server.Host = "smtp.i.noahgroup.com"
	server.Port = 25

	// Noah
	//server.Host = "mail.noahgroup.com"
	//server.Port = 25
	server.Username = "DataScience_BI"
	server.Password = "IwC1edDyfb7eWwCn"
	sender := " aaaaa <datascience_bi@noahgroup.com>"

	// Noah test
	//server.Host = "10.1.234.44"
	//server.Port = 25
	//server.Username = "testqq"
	//server.Password = " "
	//sender := "testqq@qq.com"

	// AWS
	//server.Host = "email-smtp.ap-southeast-1.amazonaws.com"
	//server.Port = 587
	//server.Username = "AKIA2AUU3RFNY456FU6Y"
	//server.Password = "BPS692oguP4iDWv+fc6cMEGb327W9s568EbQSyqpCYG/"
	//server.SendTimeout = 30

	//sender := "aws_noah@noahgroup.sg"
	recepient := "test-ou15@noah-fund.com"
	//ccemail := " "
	//server.Username = "aws_noah@noahgroup.sg"
	//server.Password = "Noah8u3$4!Q8as"
	//sender := "aws_noah@noahgroup.sg"
	//server.Password = "yang_maple@outlook.com"
	//server.Password = "iqptpzxrukgciesq"
	server.Encryption = mail.EncryptionNone
	////outlook 必须使用starttls 进行验证
	//server.Encryption = mail.EncryptionTLS
	//server.Host = "smtp-mail.outlook.com"
	//server.Port = 587
	//server.Encryption = mail.EncryptionSTARTTLS

	// You can specify authentication type:
	// - AUTO (default)
	// - PLAIN
	// - LOGIN
	// - CRAM-MD5
	// - None
	//server.Authentication = mail.AuthPlain

	// Variable to keep alive connection
	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	//server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// New email simple html with inline and CC
	email := mail.NewMSG()
	email.AddHeader("X-Notify-Message-ID", "3121639760461820")
	email.SetFrom(sender).
		AddTo(recepient).
		//	AddCc(ccemail).
		SetSubject("No-reply")
	//取消订阅
	//SetListUnsubscribe("<mailto:unsubscribe@example.com?subject=https://example.com/unsubscribe>")

	// 解析模板
	email.SetBody(mail.TextHTML, htmlBody)
	//email.AddHeader("X-Notify-Message-ID", "3121639760461820")
	// also you can add body from []byte with SetBodyData, example:
	// email.SetBodyData(mail.TextHTML, []byte(htmlBody))
	// or alternative part
	//email.AddAlternativeData(mail.TextHTML, []byte(htmlBody))

	// add inline
	//email.Attach(&mail.File{FilePath: "../dzfp_20240801172853.pdf", Name: "dzfp.pdf", Inline: false})
	//email.Attach(&mail.File{FilePath: "../dzfp_20240801172853.pdf", Name: "go.pdf", Inline: false})
	//email.Attach(&mail.File{FilePath: "../dzsw.pdf", Name: "go1.pdf", Inline: false})
	//email.Attach(&mail.File{FilePath: "../dzsw.pdf", Name: "go2.pdf", Inline: false})
	//email.Attach(&mail.File{FilePath: "../dzsw.pdf", Name: "go3.pdf", Inline: false})
	// dsn 状态
	// also you can set Delivery Status Notification (DSN) (only is set when server supports DSN)
	//email.SetDSN([]mail.DSN{mail.SUCCESS, mail.FAILURE}, true)
	// 添加已读回执
	////you can add dkim signature to the email.
	////to add dkim, you need a private key already created one.
	////电子签名
	//if privateKey != "" {
	//	options := dkim.NewSigOptions()
	//	options.PrivateKey = []byte(privateKey)
	//	options.Domain = "example.com"
	//	options.Selector = "default"
	//	options.SignatureExpireIn = 3600
	//	options.Headers = []string{"from", "date", "mime-version", "received", "received"}
	//	options.AddSignatureTimestamp = true
	//	options.Canonicalization = "relaxed/relaxed"
	//	email.SetDkim(options)
	//}

	// always check error after send
	if email.Error != nil {
		log.Fatal(email.Error)
	}

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {

		log.Println("Email Sent")
	}
}
