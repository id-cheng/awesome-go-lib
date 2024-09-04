package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.Set("redis.port", 16379)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read config failed:", err)
	}
	fmt.Println("redis port: ", viper.Get("redis.port"))
}
