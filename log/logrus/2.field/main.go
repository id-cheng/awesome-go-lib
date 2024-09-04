package main

import "github.com/sirupsen/logrus"

/*
调用logrus.WithField
接受一个logrus.Fields类型的参数
*/
func main() {
	logrus.WithFields(logrus.Fields{
		"name": "cheng",
		"age":  20,
	}).Info("info msg")

	logger := logrus.WithFields(logrus.Fields{
		"user_id": 8888,
		"ip":      "127.0.0.1",
	})
	logger.Warn("warn msg")
	logger.Error("err msg")
}
