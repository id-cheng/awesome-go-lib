package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string `json:"name,omitempty"` //相关字段的值为空，则不会被用于编码
	Age         int    `json:"-"`              //忽略字段
	CountryCode string `json:"country_code"`   //下划线
}

func main() {
	user := User{"cheng", 20, "CN"}
	bytes, _ := json.Marshal(user)
	fmt.Println(string(bytes))

	//生成具有缩进功能的JSON(pretty-print)
	bytesIndent, _ := json.MarshalIndent(user, "", "\t")
	fmt.Println(string(bytesIndent))
}
