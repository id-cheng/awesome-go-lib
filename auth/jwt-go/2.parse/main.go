package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func main() {
	jwtKey := make([]byte, 32) // 生成256位的密钥
	if _, err := rand.Read(jwtKey); err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "cheng",
		"sub": "com",
		"exp": time.Now().Add(time.Second * 60).UnixMilli(),
	})
	jwtStr, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}
	// 解析 jwt
	claims, err := ParseJwt(jwtKey, jwtStr, jwt.WithExpirationRequired())
	if err != nil {
		panic(err)
	}
	claims2, err := ParseJwtWithClaims(jwtKey, jwtStr, jwt.WithExpirationRequired())
	if err != nil {
		panic(err)
	}
	fmt.Println(claims.GetIssuer())
	fmt.Println(claims2.GetIssuer())
}

func ParseJwt(key any, jwtStr string, options ...jwt.ParserOption) (jwt.Claims, error) {
	token, err := jwt.Parse(jwtStr,
		func(token *jwt.Token) (any, error) {
			return key, nil
		}, options...)
	if err != nil {
		return nil, err
	}
	// 校验 Claims 对象是否有效
	// 基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims, nil
}

func ParseJwtWithClaims(key any, jwtStr string, options ...jwt.ParserOption) (jwt.Claims, error) {
	mc := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtStr, mc, func(token *jwt.Token) (any, error) {
		return key, nil
	}, options...)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	fmt.Println(mc)
	return token.Claims, nil
}
