package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	// 绑定环境变量
	viper.AutomaticEnv()
	viper.BindEnv("go.path", "GOPATH")
}
func main() {
	fmt.Println("gopath:", viper.Get("GOPATH"))
	fmt.Println("gopath:", viper.Get("go.path"))
}
