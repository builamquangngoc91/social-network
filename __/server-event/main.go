package main

import (
	"fmt"
	"time"

	"github.com/r3labs/sse"
	"o.o/backend/com/fabo/pkg/webhook"
	"o.o/common/jsonx"
)

func main() {
	client := sse.NewClient("https://streaming-graph.facebook.com/1141048249747549/live_comments?access_token=EAADth8IZBqFMBAIdQx9AHDBByZBmE6DLwzCjlHM01PpmZBqmLprQIlEDJHZCZCr0rTgecQYCeZAjgGgbA6SWcIUWM10muRcomCUTZAKgbIM9UAZBqtOoS1nAgXvK473UUggKZAslyZCTCsKXMpqa5TjgHFPEAqJqoTMnZANtQPFpmr4FaeJwtums24HVhmQslBOZAFrx1ZBQOCYA8PAZDZD&comment_rate=one_per_two_seconds&fields=from{name,id,email,first_name,last_name,picture},message,attachment,created_time&date_format=U")

	client.SubscribeRaw(func(msg *sse.Event) {
		fmt.Println("--------")
		// Got some data!
		var comment webhook.LiveVideoComment
		if err := jsonx.Unmarshal(msg.Data, &comment); err != nil {
			return
		}

		t, _ := time.Parse("2006-01-02T15:04:05-0700", comment.CreatedTime)
		fmt.Println(t.String())
	})
}
