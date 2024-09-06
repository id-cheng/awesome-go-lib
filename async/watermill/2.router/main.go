package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
	"time"
)

var (
	logger = watermill.NewStdLogger(false, false)
)

// 使用路由还有个好处，处理器返回时，若无错误，路由会自动调用消息的Ack()方法
// 若发生错误，路由会调用消息的Nack()方法通知管理器重发这条消息。
func main() {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return
	}
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	// 发布消息
	go publishMessages(pubSub)

	//创建一个名为handler的处理器
	//监听subscriber中主题为subscribeTopic(in_topic)的消息
	//收到消息后调用handlerFunc(myHandler{}.Handler)处理
	//将返回的消息以主题publishTopic(out_topic)发布到publisher中
	router.AddHandler("handler", "in_topic", pubSub, "out_topic", pubSub, myHandler{}.Handler)

	//只处理接收到的消息，不发布新消息
	router.AddNoPublisherHandler("in_messages", "in_topic", pubSub, printMessages)
	router.AddNoPublisherHandler("out_messages", "out_topic", pubSub, printMessages)

	//router.Run()运行这个路由
	if err := router.Run(context.Background()); err != nil {
		return
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
	fmt.Printf("Received message: %v\n %v\n\n", msg.UUID, string(msg.Payload))
	return nil
}

type myHandler struct{}

func (m myHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	log.Println("myHandler received message", msg.UUID)
	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by myHandler"))
	return message.Messages{msg}, nil
}
