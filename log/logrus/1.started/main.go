package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	//日志级别从上向下依次增加，Trace最大，Panic最小
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("trace msg")
	logrus.Debug("debug msg")

	// 设置在输出日志中添加文件名和方法信息
	logrus.SetReportCaller(true)
	logrus.Info("info msg")

	//logrus支持两种日志格式，文本和 JSON，默认为文本格式
	//通过logrus.SetFormatter设置日志格式
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Warn("warn msg")

	//设置标准日志记录器是否包含调用者的信息
	logrus.SetReportCaller(true)
	logrus.Error("error msg")

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.Fatal("fatal msg")
	logger.Panic("panic msg")
}
