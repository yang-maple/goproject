package main

import (
	"encoding/base64"
	"encoding/json"
	"gopkg.in/gomail.v2"
	"log"
)

//const htmlBody = `
//<html>
//<div>亲爱的伙伴，您好!</div>
//<div>本月底，您名下将有2位客户的积分到期，建议您提醒客户及时前去微诺亚积分商城进行积分兑换，积分一经逾期无法恢复，请知悉为感</div>
//<div>FR端：提醒短信及邮件将于今日陆续发送至FR手机号及邮箱</div>
//<div>客户端：提醒短信将于今日下午陆续发送至客户手机</div>
//<div><div style="padding-left: 2em">到期积分明细如下:</div><div style="padding-left: 2em"><table border="1" cellspacing="0"><tr><th>序号</th><th>会员编号</th><th>集团号</th><th>姓名</th><th>会员等级</th><th>目前积分</th><th>月底即将到期积分</th></tr><tr><td>1</td><td>927787</td><td>000335616</td><td>毛*</td><td>白银</td><td>13</td><td>12</td></tr><tr><td>2</td><td>371856</td><td>000507326</td><td>许*锋</td><td>白银</td><td>2</td><td>2</td></tr></table></div></div>
//<div>2024年09月24日Tue</div>
//<div>尊敬的投資者：</div>
//<div>中秋佳節至，月圓人團圓，我們與您共慶佳節！</div>
//<div>福利一：認購任意公募基金（含現金寶），您將有機會獲得費用減免獎賞，最高可達HKD12,688。</div>
//<div>福利二：您可享10%的佣金折扣，最高可享50元。</div>
//<div>福利二：聯繫您的國際持牌代表獲取818理財節優惠碼，部分產品可專享5折認購費優惠（限美元、港幣份額）。</div>
//<div>活動條件與條款適用，具體活動詳情登錄方舟達富APP或點擊https://noahmkt.com/2P7x查看。</div>
//<div></div>
//<div>此邮件由系统发出，请勿回复。如有疑问，请联系区域客经负责人。</div>
//</html>
//
//`

const htmlBody = `
<html>
<div>no-reply no-reply</div>
</html>
`

func main() {
	//ALI YUN HZ EMAIL SERVER
	//smtpAddress := "smtpdm.aliyun.com"
	//port := 80
	//username := "noahfundtest@notice.noahgrouptest.com"
	//password := "1nW4UqVx66grzbEs"
	//sender := "cccccc<noahfundtest@notice.noahgrouptest.com>"
	//username := "testxiy@notice.noahgrouptest.com"
	//password := "QWert12345"
	//sender := "ddd<testxiy@notice.noahgrouptest.com>"
	//username := "testxiy@notice.noahgrouptest.sg"
	//password := "QWert12345"
	//sender := "qqqqqq<testxiy@notice.noahgrouptest.sg>"

	//ALI YUN SG EMAIL SERVER
	smtpAddress := "smtpdm-ap-southeast-1.aliyun.com"
	port := 80
	username := "testxiy@noticesg.noahgrouptest.sg"
	password := "QWert12345"
	sender := "cccccc<testxiy@noticesg.noahgrouptest.sg>"

	// Noah
	//smtpAddress := "mail.noahgroup.com"
	//port := 25
	//username := "mail04@noahgroup.com"
	//password := "vx!Z2FQSa@zW"
	//sender := "mail04@noahgroup.com"

	recepient := "xmaple725@gmail.com"
	tagName := "testsg"
	trach := map[string]string{}
	trach["OpenTrace"] = "1"
	//trach["LinkTrace"] = "1"
	trach["TagName"] = tagName
	jsontrach, err := json.Marshal(trach)
	if err != nil {
		log.Println(err)
	}
	base64Trace := base64.StdEncoding.EncodeToString(jsontrach)
	//ccemail := ""
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recepient)
	//m.SetAddressHeader("Cc", ccemail, " ")
	m.SetHeader("Subject", "NO REPLY1111")
	m.SetHeader("X-Notify-Message-ID", "3121639760461820")
	m.SetHeader("X-AliDM-Trace", base64Trace)
	//m.SetHeader("Reply-To", "xiongyang@noahgroup.com")
	m.SetBody("text/html", htmlBody)
	//m.Attach("../dzjenkins.pdf", gomail.Rename("goalng1.pdf"))

	d := gomail.NewDialer(smtpAddress, port, username, password)

	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
	println("ok")

}
