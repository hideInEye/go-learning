package main

import (
	"fmt"
	"reflect"
)

type UserInfo struct {
	Name    string
	Age     int
	Sex     string
	Address string
}

func DoUserInfo(user interface{}) map[string]interface{} {
	userType := reflect.TypeOf(user)
	fmt.Println("type", userType)
	userValue := reflect.ValueOf(user)
	fmt.Println("userValue", userValue)

	a := make(map[string]interface{})
	for i := 0; i < userType.NumField(); i++ {
		a[userType.Field(i).Name] = userValue.Field(i).Interface()
	}
	return a
}

func main() {
	value := DoUserInfo(UserInfo{Name: "test", Age: 22, Sex: "男", Address: "北京"})
	fmt.Println("value", value)
}
