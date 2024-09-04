package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.Set("app_name", "awesome")
	viper.Set("log_level", "DEBUG")
	viper.Set("mysql.ip", "127.0.0.1")
	viper.Set("mysql.port", 3306)
	viper.Set("mysql.user", "root")
	viper.Set("mysql.password", "123456")
	viper.Set("mysql.database", "awesome")

	viper.Set("redis.ip", "127.0.0.1")
	viper.Set("redis.port", 6379)

	err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatal("write config failed: ", err)
	}
}
