package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	CountryCode string `json:"country_code"`
}

func main() {
	str := `{
		"name": "cheng",
		"age":18,
		"country_code": "CN"
    }`
	arrStr := `[
     { 
		"name": "cheng",
		"age":18,
		"country_code": "CN"
     },
     { 
		"name": "jia",
		"age":18,
		"country_code": "CN"
     }
    ]`
	var user User
	if err := json.Unmarshal([]byte(str), &user); err != nil {
		return
	}
	fmt.Println("user:", user)

	var m map[string]any
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		return
	}
	fmt.Println("map:", m)

	var users []User
	if err := json.Unmarshal([]byte(arrStr), &users); err != nil {
		return
	}
	fmt.Println("users:", users)

	var maps []map[string]any
	if err := json.Unmarshal([]byte(arrStr), &maps); err != nil {
		return
	}
	fmt.Println("maps:", maps)
}
