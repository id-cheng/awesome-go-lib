package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	env, _ := godotenv.Unmarshal("KEY=value")
	content, _ := godotenv.Marshal(env)
	fmt.Println(content)
	if err := godotenv.Write(env, "./.env"); err != nil {
		return
	}
}
