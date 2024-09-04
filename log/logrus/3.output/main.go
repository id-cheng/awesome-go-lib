package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func main() {
	buf := &bytes.Buffer{}
	stout := os.Stdout
	file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	logrus.SetOutput(io.MultiWriter(buf, stout, file))
	logrus.Info("info msg")
}
