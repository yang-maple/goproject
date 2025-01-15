package main

import (
	"context"
	"fmt"
	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	graphusers "github.com/microsoftgraph/msgraph-sdk-go/users"
)

func main() {
	//基本信息
	//tenantID := ""     //租户ID
	//clientID := ""     //程序ID
	//clientSecret := "" //程序密码
	// 创建客户端
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		panic(err)
	}
	//验证是否获取到token
	//token, err := cred.GetToken(context.Background(), policy.TokenRequestOptions{
	//	Scopes: []string{"https://graph.microsoft.com/.default"}, // https://docs.microsoft.com/zh-cn/azure
	//})
	//fmt.Println(token.Token)
	// 创建客户端
	// scopes 字段在sdk 默认使用 https://graph.microsoft.com/.default 即可
	graphClient, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"https://graph.microsoft.com/.default"})
	//users, err := graphClient.Users().Get(context.Background(), nil)
	//for _, v := range users.GetValue() {
	//	fmt.Println(*v.GetDisplayName())
	//}
	////列出所有邮件主题， top 指定数量
	top := int32(32)
	requestParameters := &graphusers.ItemMessagesRequestBuilderGetQueryParameters{
		Top:    &top,
		Select: []string{"sender", "subject"},
	}
	configuration := &graphusers.ItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	messages, err := graphClient.Users().ByUserId("").Messages().Get(context.Background(), configuration)
	if err != nil {
		panic(err.Error())
	}
	for _, v := range messages.GetValue() {
		fmt.Println(*v.GetSubject(), *v.GetId())
		err = graphClient.Users().ByUserId("").Messages().ByMessageId(*v.GetId()).Delete(context.Background(), nil)
		if err != nil {
			panic(err)
		}
	}

	//// 创建邮件
	//requestBody := graphusers.NewItemSendMailPostRequestBody() //定义邮件请求体
	//message := graphmodels.NewMessage()                        //创建邮件
	//subject := "Meet for lunch?"                               //主题
	//message.SetSubject(&subject)
	//body := graphmodels.NewItemBody()        //内容
	//contentType := graphmodels.TEXT_BODYTYPE //类型
	//body.SetContentType(&contentType)
	//content := "The new cafeteria is open." //内容
	//body.SetContent(&content)
	//message.SetBody(body) //
	//
	////收件人
	//recipient := graphmodels.NewRecipient()
	//emailAddress := graphmodels.NewEmailAddress()
	//address := "xiongyang@noahgroup.com"
	//emailAddress.SetAddress(&address)
	//recipient.SetEmailAddress(emailAddress)
	//toRecipients := []graphmodels.Recipientable{
	//	recipient,
	//}
	//message.SetToRecipients(toRecipients)
	//
	////抄送
	//recipients := graphmodels.NewRecipient()
	//emailAddresss := graphmodels.NewEmailAddress()
	//addresss := "danas@contoso.com"
	//emailAddresss.SetAddress(&addresss)
	//recipients.SetEmailAddress(emailAddresss)
	//
	//ccRecipients := []graphmodels.Recipientable{
	//	recipients,
	//}
	//// 邮件头
	////internetMessageHeader := graphmodels.NewInternetMessageHeader()
	////name := "X-Return-Receipt-To"
	////internetMessageHeader.SetName(&name)
	////value := "aws_noah@noahgroup.sg"
	////internetMessageHeader.SetValue(&value)
	//
	////internetMessageHeaders := []graphmodels.InternetMessageHeaderable{
	////	internetMessageHeader,
	////}
	////message.SetInternetMessageHeaders(internetMessageHeaders)
	//message.SetCcRecipients(ccRecipients)
	//requestBody.SetMessage(message)
	//// 是否保存到发送文件夹 默认true
	////saveToSentItems := false
	////requestBody.SetSaveToSentItems(&saveToSentItems)
	//
	//err = graphClient.Users().ByUserId("").SendMail().Post(context.Background(), requestBody, nil)
	//if err != nil {
	//	panic(err)
	//}

	//删除邮件
	//err = graphClient.Users().ByUserId("").Messages().ByMessageId("").Delete(context.Background(), nil)
	//if err != nil {
	//	panic(err)
	//}
}
