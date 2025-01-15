package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func main() {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-shanghai.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
<<<<<<< HEAD
	client, err := oss.New("https://oss-cn-shanghai.aliyuncs.com", "LTAI5tRwj9HgcUdwHbNbswfu", "qDAGdxiIobXodHyQWwIyZ2aglNFt2t")
=======
	client, err := oss.New("https://oss-cn-shanghai.aliyuncs.com", "xxx", "ddd")
>>>>>>> 4622a0a (2025-1-15)
	if err != nil {
		HandleError(err)
	}

	// 指定图片所在Bucket的名称，例如examplebucket。
	bucketName := "yang-photo"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	// 指定图片名称。如果图片不在Bucket根目录，需携带文件完整路径，例如exampledir/example.jpg。
	ossImageName := "bizhi/1316292.jpeg"
	// 生成带签名的URL，并指定过期时间为600s。
	signedURL, err := bucket.SignURL(ossImageName, oss.HTTPGet, 600, oss.Process("image/format,png"))
	if err != nil {
		HandleError(err)
	} else {
		fmt.Println(signedURL)
	}
}
