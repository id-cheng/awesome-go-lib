package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	mapClaims := jwt.MapClaims{
		"iss": "cheng",
		"sub": "com",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	token2 := jwt.New(jwt.SigningMethodHS256)
	fmt.Println(token.SigningString())
	fmt.Println(token2.SignedString([]byte("123")))
}
