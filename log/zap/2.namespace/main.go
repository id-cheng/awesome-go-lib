package main

import "go.uber.org/zap"

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	//zap.Namespace(key string) Field构建一个命名空间
	//后续的Field都记录在此命名空间中
	logger.Info("info msg",
		zap.Namespace("space1"),
		zap.Int("counter", 1),
	)

	//调用With()创建一个新的Logger
	//新的Logger记录日志时总是带上预设的字段
	logger2 := logger.With(
		zap.Namespace("space2"),
		zap.Int("counter", 10),
	)
	logger2.Info("info msg")
}
