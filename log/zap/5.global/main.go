package main

import "go.uber.org/zap"

// zap提供了两个全局的Logger
// 一个是*zap.Logger 可调用zap.L()获得
// 另一个是*zap.SugaredLogger 可调用zap.S()获得
func main() {
	zap.L().Info("global Logger before")
	zap.S().Info("global SugaredLogger before")

	logger := zap.NewExample()
	defer logger.Sync()

	// 在调用ReplaceGlobals之前记录的日志并没有输出
	zap.ReplaceGlobals(logger)
	zap.L().Info("global Logger after")
	zap.S().Info("global SugaredLogger after")
}
