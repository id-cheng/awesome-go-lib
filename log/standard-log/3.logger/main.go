package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

// 实现一个简单的logger
func main() {
	buf := new(bytes.Buffer)
	stdout := os.Stdout
	file, err := os.OpenFile("./log.txt",
		os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	logger := log.New(io.MultiWriter(buf, stdout, file),
		"logger:", log.Lshortfile|log.LstdFlags)
	logger.Println("test")
}
