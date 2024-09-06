package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

// 与标准库100%兼容
// 性能比标准库好
var json = jsoniter.ConfigCompatibleWithStandardLibrary

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user1 := map[string]any{
		"name": "cheng",
		"age":  18,
	}
	bytes, _ := json.Marshal(user1)
	fmt.Println(string(bytes))

	user2 := User{"jia", 20}
	bytes, _ = json.Marshal(user2)
	fmt.Println(string(bytes))
}
