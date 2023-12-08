package dao

import (
	"errors"
	"fmt"
	"kubeops/db"
	"kubeops/model"
)

var Login login

type login struct {
}

func (l *login) VerifyUser(username, password string) (err error) {
	var userInfo model.User
	tx := db.GORM.Where("username = ?", username).First(&userInfo)
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		return errors.New("用户名或密码错误" + err.Error())
	}
	fmt.Println(userInfo.Username, userInfo.Password)
	if userInfo.Password != password {
		return errors.New("用户名或密码错误" + err.Error())
	}
	return nil
}
