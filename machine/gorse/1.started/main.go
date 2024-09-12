package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var json = []byte(`
[
    {
        "FeedbackType": "star",
        "UserId": "bob",
        "ItemId": "vuejs:vue",
        "Timestamp": "2022-02-24"
    },
    {
        "FeedbackType": "star",
        "UserId": "bob",
        "ItemId": "d3:d3",
        "Timestamp": "2022-02-25"
    },
    {
        "FeedbackType": "star",
        "UserId": "bob",
        "ItemId": "dogfalo:materialize",
        "Timestamp": "2022-02-26"
    },
    {
        "FeedbackType": "star",
        "UserId": "bob",
        "ItemId": "mozilla:pdf.js",
        "Timestamp": "2022-02-27"
    },
    {
        "FeedbackType": "star",
        "UserId": "bob",
        "ItemId": "moment:moment",
        "Timestamp": "2022-02-28"
    }
]`)

func main() {
	//假设 Bob 是一名前端开发人员
	//他在 GitHub 中点赞过为多个前端仓库
	//我们通过 API 将他的点赞插入 Gorse
	resp, _ := http.Post("http://127.0.0.1:8088/api/feedback", "application/json", bytes.NewBuffer(json))
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	// Gorse 获取 10 个推荐项。我们可以发现 Gorse 为 cheng 推荐了前端相关的仓库。
	resp, _ = http.Get("http://127.0.0.1:8088/api/recommend/cheng?n=10")
	body, _ = io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
