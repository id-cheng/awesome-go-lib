package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := []User{
		{"jia", 18},
		{"cheng", 30},
	}
	bytes, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(bytes))
}

func (u *User) MarshalJSON() ([]byte, error) {
	type user User
	return json.Marshal(&struct {
		*user
		Email string `json:"email"`
	}{
		user:  (*user)(u),
		Email: "id.cheng@gamil.com",
	})
}
