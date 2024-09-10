package main

import "github.com/tidwall/gjson"

const json = `{"name":{"first":"jia","last":"cheng"},"age":20}`

func main() {
	value := gjson.Get(json, "name.last")
	println(value.String())
	age := gjson.Get(json, "age")
	println(age.Int())
	// 返回相同的结果
	gjson.Parse(json).Get("name").Get("last")
	gjson.Get(json, "name").Get("last")
	gjson.Get(json, "name.last")
}
