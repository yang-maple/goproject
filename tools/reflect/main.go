package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Sex   string `json:"sex"`
	Email string `json:"email"`
}

func main() {
	user := &User{
		Id:    1,
		Name:  "Tom",
		Age:   20,
		Sex:   "male",
		Email: "tom@gmail.com",
	}
	newinfo := &User{
		Name: "Tom",
		Age:  21,
		Sex:  "male",
	}
	fmt.Println(buildUpdates(newinfo, user))
}

func buildUpdates(user, newInfo interface{}) (map[string]interface{}, error) {
	updates := make(map[string]interface{})
	userValue := reflect.ValueOf(user).Elem()
	newInfoValue := reflect.ValueOf(newInfo).Elem()

	for i := 0; i < userValue.NumField(); i++ {
		fieldName := userValue.Type().Field(i).Name
		userFieldValue := userValue.Field(i).Interface()
		newInfoFieldValue := newInfoValue.FieldByName(fieldName).Interface()

		// 比较字段值是否不同
		if !reflect.DeepEqual(userFieldValue, newInfoFieldValue) {
			updates[strings.ToLower(fieldName)] = newInfoFieldValue
		}
	}

	return updates, nil
}
