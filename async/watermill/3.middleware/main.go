package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

var (
	logger = watermill.NewStdLogger(false, false)
)

func main() {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	go publishMessages(pubSub)

	//调用两次处理函数，两次返回的消息都重新发布出去
	router.AddMiddleware(middleware.Duplicator)
	//捕获处理函数中的panic，包装成错误返回
	router.AddMiddleware(middleware.Recoverer)
	//如果消息处理时间超过给定的时间，直接失败
	router.AddMiddleware(middleware.Timeout(3 * time.Second))
	// 设置最大重试次数为 3，重试初始时间间隔为 1s。
	router.AddMiddleware(middleware.Retry{
		MaxRetries:      3,
		InitialInterval: time.Second,
		Logger:          logger,
	}.Middleware)
	// 自定义中间件
	router.AddMiddleware(MyMiddleware)

	router.AddNoPublisherHandler("print_out_messages", "out_topic", pubSub, printMessages)

	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
		if err := publisher.Publish("in_topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func printMessages(msg *message.Message) error {
	fmt.Printf("Received message: %s\n %s\n \n", msg.UUID, string(msg.Payload))
	return nil
}

func MyMiddleware(h message.HandlerFunc) message.HandlerFunc {
	return func(message *message.Message) ([]*message.Message, error) {
		ms, err := h(message)
		return ms, err
	}
}
