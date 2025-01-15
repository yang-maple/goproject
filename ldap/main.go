package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

func main() {
	BindUserDn := "CN=熊洋管理员,OU=SSGroup,OU=NOAH AdminGroup,DC=noahgrouptest,DC=com,DC=local"
	BindUserPassword := "20200725.zyj"
	username := "nd_test1"
	address := "10.1.234.41"
	port := "636"

	l, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%s", address, port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	// 开启 TLS
	err = l.StartTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	// 绑定用户
	err = l.Bind(BindUserDn, BindUserPassword)
	if err != nil {
		log.Fatal(err)
	}

	// 查询用户
	searchRequest := ldap.NewSearchRequest(
		"DC=noahgrouptest,DC=com,DC=local",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(sAMAccountName=%s)", ldap.EscapeFilter(username)),
		[]string{"distinguishedName", "lockoutTime"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) != 1 {
		log.Fatal("User does not exist or too many entries returned")
	}

	//userdn := sr.Entries[0]
	//dn := ""
	for _, entry := range sr.Entries {
		for _, v := range entry.Attributes {
			fmt.Println(v.Name, v.Values)
			//dn = v.Values[0]
		}
	}

	//modifyRequest := ldap.NewModifyRequest(dn, nil)
	//modifyRequest.Replace("lockoutTime", []string{"0"})
	//err = l.Modify(modifyRequest)
	//if err != nil {
	//	fmt.Errorf("无法解锁用户: %v", err)
	//}

	// 验证用户登录
	//err = l.Bind(sr.Entries[0].DN, password)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("login success")
}

//func convertLockoutTime(lockoutTime int64) string {
//	// 将 lockoutTime 转换为以 100 纳秒为单位的时间间隔
//	nanoseconds := lockoutTime * 100
//	const windowsToUnix = 116444736000000000
//	unixNanos := nanoseconds - windowsToUnix
//	fmt.Println("unixNanos", time.Unix(0, unixNanos))
//	// 计算时间偏移量
//	epoch := time.Date(1601, time.January, 1, 0, 0, 0, 0, time.UTC)
//
//	// 添加时间偏移量
//	lockoutDateTime := epoch.Add(time.Duration(nanoseconds))
//
//	// 将结果转换为可读的日期和时间格式
//	formattedTime := lockoutDateTime.Format("2006-01-02 15:04:05")
//
//	date := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
//
//	// Windows 时间戳起点
//	windowsEpoch := time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC)
//
//	// 计算指定日期与起点的时间差
//	timestampDifference := date.Sub(windowsEpoch)
//
//	// 将时间差转换为以 100 纳秒为单位的数值
//	timestamp := int64(timestampDifference.Nanoseconds() / 100)
//
//	fmt.Println("timestamp", timestamp)
//
//	return formattedTime
//}
