package main

import (
	"log"
)

func main() {
	log.SetPrefix("prefix:")            // 设置log的前缀
	log.Println("print:", log.Prefix()) // 正常输出日志
	go func() {
		log.Fatalln("fatal") // 输出日志后，调用os.Exit(1)退出程序
	}()
	log.Panicln("panic") // 输出日志后调用panic
}
