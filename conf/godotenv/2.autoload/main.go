package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

// 自动加载
func main() {
	secretKey := os.Getenv("SECRET_KEY")
	fmt.Println(secretKey)
}
