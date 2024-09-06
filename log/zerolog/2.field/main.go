package main

import "github.com/rs/zerolog/log"

func main() {
	log.Debug().
		Str("Name", "cheng").
		Float64("Interval", 888.88).
		Msg("HelloWorld")
	//调用Msg()或Send()之后，日志会被输出
	log.Debug().
		Str("Name", "cheng").
		Send()
}
