package main

import "log"

// log库定义了一个Lstdflag，为Ldate | Ltime，这就是log默认的选项
// 默认情况下，每条日志前会自动加上日期和时间
func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
	log.Println("test")
}
