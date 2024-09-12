package main

import (
	"context"
	"fmt"
	"github.com/zhenghaoz/gorse/client"
)

var ctx = context.Background()

func main() {
	// Create a client
	gorse := client.NewGorseClient("http://127.0.0.1:8087", "api_key")

	// Insert feedback
	feedback, _ := gorse.InsertFeedback(ctx, []client.Feedback{
		{FeedbackType: "star", UserId: "bob", ItemId: "vuejs:vue", Timestamp: "2022-02-24"},
		{FeedbackType: "star", UserId: "bob", ItemId: "d3:d3", Timestamp: "2022-02-25"},
		{FeedbackType: "star", UserId: "bob", ItemId: "dogfalo:materialize", Timestamp: "2022-02-26"},
		{FeedbackType: "star", UserId: "bob", ItemId: "mozilla:pdf.js", Timestamp: "2022-02-27"},
		{FeedbackType: "star", UserId: "bob", ItemId: "moment:moment", Timestamp: "2022-02-28"},
	})
	fmt.Println(feedback)

	// Get recommendation.
	recommend, _ := gorse.GetRecommend(ctx, "bob", "", 10)
	fmt.Println(recommend)
}
