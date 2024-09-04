package main

import "time"
import "go.uber.org/zap"

func main() {
	// zap提供了几个快速创建logger的方法
	// zap.NewExample()、zap.NewDevelopment()、zap.NewProduction()
	// Example适合用在测试代码中，Development在开发环境中使用，Production用在生成环境
	logger := zap.NewExample()
	logger, _ = zap.NewDevelopment()
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	url := "http://example.org/api"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("time", time.Second),
	)

	// zap.Fields 预设日志字段
	logger2 := zap.NewExample(zap.Fields(
		zap.Int("id", 1),
		zap.String("name", "cheng"),
	))
	logger2.Info("hello world")

	//zap也提供了便捷的方法SugarLogger，可以使用printf格式符的方式。
	//调用logger.Sugar()即可创建SugaredLogger。
	//SugaredLogger的使用比Logger简单，只是性能比Logger低50%左右
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"time", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
