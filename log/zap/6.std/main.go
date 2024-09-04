package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

// 适用于项目一开始使用的是标准日志库log，后面想转为zap
func main() {
	logger := zap.NewExample()
	defer logger.Sync()
	//调用zap.NewStdLog(l *Logger)返回一个标准库的log.Logger
	stdLog := zap.NewStdLog(logger)
	stdLog.Println("standard logger")

	//让标准logger以level级别写入内部的*zap.Logger
	warnLog, _ := zap.NewStdLogAt(logger, zapcore.WarnLevel)
	warnLog.Println("standard warn logger")

	//RedirectStdLog(l *Logger) func()。它会返回一个无参函数恢复
	undo := zap.RedirectStdLog(logger)
	log.Print("redirected standard library")

	//恢复标准log输出格式
	undo()
	log.Print("restored standard library")
}
