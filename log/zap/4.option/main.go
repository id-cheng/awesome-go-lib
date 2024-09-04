package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}
func main() {
	// 调用zap.AddCaller()返回的选项设置输出文件名和行号
	// 希望输出的文件名和行号是调用封装函数的位置
	// 这时可以使用zap.AddCallerSkip(skip int)向上跳 1 层
	logger, _ := zap.NewProduction(zap.AddCaller())
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	info("helloWorld!!!")

	//zap.AddStackTrace(zapcore.LevelEnabler) 输出调用堆栈
	logger2, _ := zap.NewProduction(zap.AddStacktrace(zapcore.WarnLevel))
	defer logger2.Sync()
	zap.ReplaceGlobals(logger2)
	warn("helloWorld!!!")
}
