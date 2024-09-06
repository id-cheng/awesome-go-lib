package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello, world! ")
}

// air会自动编译，启动程序，并监听当前目录中文件
// 直接执行air命令，使用的就是默认的配置
// 使用air init初始化配置文件
// 如果想查看air更详细的执行流程，可以使用-d选项
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}
