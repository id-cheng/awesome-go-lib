package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
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
