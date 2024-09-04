package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	done := make(chan bool)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return
	}

	viper.WatchConfig() // 监听文件
	fmt.Println("redis port before: ", viper.Get("redis.port"))

	//文件修改时会执行这个回调
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("file:%s,op:%s\n", e.Name, e.Op)
		fmt.Println("redis port after: ", viper.Get("redis.port"))
		done <- true
	})
	<-done
}
