package linebot

// Line Bot publisher layer for go

import (
	"fmt"

	"github.com/franela/goreq"
)

const (
	endpoint = "https://api.line.me/v2/bot/message/push"
)

// Messages has type and text.
type Messages struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Payload includes Messages structure.
type Payload struct {
	// LINE ID
	To                   string     `json:"to"`
	Messages             []Messages `json:"messages"`
	NotificationDisabled bool       `json:"notificationDisabled"`
}

// Publish sends a message to LINE.
func Publish(toID string, token string, message string, notifies bool) (string, error) {
	if toID == "" || token == "" {
		return "", fmt.Errorf("toID and token must be needed")
	}

	messages := Payload{
		To: toID,
		Messages: []Messages{
			Messages{
				Type: "text",
				Text: message,
			},
		},
		NotificationDisabled: notifies,
	}

	res, _ := goreq.Request{
		Method:      "POST",
		Uri:         endpoint,
		ContentType: "application/json",
		Body:        messages,
	}.WithHeader("Authorization", "Bearer "+token).Do()

	body, _ := res.Body.ToString()
	return body, nil
}

func main() {
	// Example
	//   resp, _ := Publish(env.ToID, "テスト", true)
	//   fmt.Printf("%v\n", resp)
}
