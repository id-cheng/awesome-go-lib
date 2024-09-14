package main

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

type message struct {
	data string
}

type foo struct{}

func newFoo() actor.Receiver {
	return &foo{}
}

func (f *foo) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("actor started")
	case actor.Stopped:
		fmt.Println("actor stopped")
	case *message:
		fmt.Println("actor has received", msg.data)
	}
}

func main() {
	engine, _ := actor.NewEngine(actor.NewEngineConfig())
	pid := engine.Spawn(newFoo, "my_actor",
		actor.WithMaxRestarts(4),    // 最大重启次数
		actor.WithInboxSize(1024*2), //size of the inbox
		actor.WithID("bar"),
	)
	for i := 0; i < 3; i++ {
		engine.Send(pid, &message{data: fmt.Sprintf("hello world %d!", i)})
	}
	engine.Poison(pid).Wait()
}
