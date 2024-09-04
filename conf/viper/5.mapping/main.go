package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppName  string
	LogLevel string
	MySQL    MySQLConfig
	Redis    RedisConfig
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	IP   string
	Port int
}

// viper 支持将配置Unmarshal到一个结构体中，为结构体中的对应字段赋值
func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return
	}

	fmt.Println(c.MySQL)
}
