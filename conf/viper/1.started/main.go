package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("redis.port", 16379)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("read config failed:", err)
	}

	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))
}
