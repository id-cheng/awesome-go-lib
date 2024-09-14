package main

import (
	"flag"

	"mynode/apps/myapp"

	"ergo.services/application/observer"

	"ergo.services/ergo"
	"ergo.services/ergo/gen"
	"ergo.services/ergo/lib"
)

var (
	OptionNodeName   string
	OptionNodeCookie string
)

func init() {
	flag.StringVar(&OptionNodeName, "name", "MyNode@localhost", "node name")
	flag.StringVar(&OptionNodeCookie, "cookie", lib.RandomString(16), "a secret cookie for the network messaging")
}

func main() {
	var options gen.NodeOptions

	flag.Parse()

	// create applications that must be started
	apps := []gen.ApplicationBehavior{
		observer.CreateApp(observer.Options{}),
		myapp.CreateMyApp(),
	}
	options.Applications = apps

	// set network options
	options.Network.Cookie = OptionNodeCookie

	// starting node
	node, err := ergo.StartNode(gen.Atom(OptionNodeName), options)
	if err != nil {
		panic(err)
	}

	// starting process MyActor2
	if _, err := node.SpawnRegister("myactor2", factory_MyActor2, gen.ProcessOptions{}); err != nil {
		panic(err)
	}

	// starting process MyWeb
	if _, err := node.SpawnRegister("myweb", factory_MyWeb, gen.ProcessOptions{}); err != nil {
		panic(err)
	}

	node.Wait()
}
