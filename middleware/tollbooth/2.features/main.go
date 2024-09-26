package main

import (
	"time"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
)

func main() {
	lmt := tollbooth.NewLimiter(1, nil)

	// This setting means:
	// create a 1 request/second limiter and
	// every token bucket in it will expire 1 hour after it was initially set.
	// 创建一个限制器，允许每秒 1 个请求
	// 指定令牌桶在初始设置后 1 小时过期
	lmt = tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	// Configure list of places to look for IP address.
	// By default it's: "RemoteAddr", "X-Forwarded-For", "X-Real-IP"
	// If your application is behind a proxy, set "X-Forwarded-For" first.
	// 设置查找 IP 地址的顺序，适用于代理环境
	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})

	// Limit only GET and POST requests.
	// 只允许 GET 和 POST 请求通过限制器
	lmt.SetMethods([]string{"GET", "POST"})

	// Limit based on basic auth usernames.
	// You add them on-load, or later as you handle requests.
	// 添加了一些基本认证用户名（如 bob, jane 等），并移除了其中一个用户
	lmt.SetBasicAuthUsers([]string{"bob", "jane", "vip"})
	// You can remove them later as well.
	lmt.RemoveBasicAuthUsers([]string{"vip"})

	// Limit request headers containing certain values.
	// You add them on-load, or later as you handle requests.
	// 设置了一个请求头 X-Access-Token 的限制值
	lmt.SetHeader("X-Access-Token", []string{"abc123", "xyz098"})
	// You can remove all entries at once.
	// 可以移除整个请求头或特定的值
	lmt.RemoveHeader("X-Access-Token")
	// Or remove specific ones.
	lmt.RemoveHeaderEntries("X-Access-Token", []string{"limitless-token"})

	// By the way, the setters are chainable. Example:
	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}).
		SetMethods([]string{"GET", "POST"}).
		SetBasicAuthUsers([]string{"sansa"}).
		SetBasicAuthUsers([]string{"tyrion"})
}
