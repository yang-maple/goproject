package ldap

import (
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/wonderivan/logger"
)

var (
	BindUserDn       = "CN=熊洋管理员,OU=SSGroup,OU=NOAH AdminGroup,DC=noahgrouptest,DC=com,DC=local"
	BindUserPassword = "20200725.zyj"
	address          = "ldap.i.noahgrouptest.com"
	port             = "389"
)

func Authentication(username, password string) error {
	l, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%s", address, port))
	if err != nil {
		logger.Error("ldap dial error: " + err.Error())
		return errors.New("系统繁忙")
	}
	defer l.Close()
	// 绑定用户
	err = l.Bind(BindUserDn, BindUserPassword)
	if err != nil {
		logger.Error("ldap bind error: " + err.Error())
		return errors.New("系统繁忙 ")
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
		[]string{""},
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil {
		logger.Error("ldap search error: " + err.Error())
		return errors.New("用户名或密码错误")
	}
	if len(sr.Entries) != 1 {
		logger.Error("User does not exist or too many entries returned")
		return errors.New("用户名或密码错误")
	}
	// 验证用户密码
	err = l.Bind(sr.Entries[0].DN, password)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("用户名或密码错误")
	}

	return nil
}
