package main

import (
	"net/http"

	"github.com/didip/tollbooth/v7"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// Create a request limiter per handler.
	// 每秒只能有一个请求
	http.Handle("/", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil), HelloHandler))
	http.ListenAndServe(":12345", nil)
}
