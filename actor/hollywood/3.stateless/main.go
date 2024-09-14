package main

import (
	"fmt"
	"github.com/anthdm/hollywood/actor"
)

func main() {
	engine, _ := actor.NewEngine(actor.NewEngineConfig())
	pid := engine.SpawnFunc(func(c *actor.Context) {
		switch msg := c.Message().(type) {
		case actor.Started:
			fmt.Println("started")
		case string:
			fmt.Println("msg:", msg)
		}
	}, "foo")
	engine.Send(pid, "helloWorld")
	engine.Poison(pid).Wait()
}
