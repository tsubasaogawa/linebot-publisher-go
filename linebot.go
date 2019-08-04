package linebot

// Line Bot publisher layer for go

import (
	"fmt"

	"github.com/franela/goreq"
	"github.com/kelseyhightower/envconfig"
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

// Env is the object consisted of environment variables used by envconfig
type Env struct {
	AccessToken string `required:"true"`
	ToID        string `required:"true"`
}

// Publish sends a message to LINE.
func Publish(toID string, message string, notifies bool) (string, error) {
	env := Env{}
	if err := envconfig.Process("", &env); err != nil {
		_ = fmt.Errorf("Error: %s", err)
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
	}.WithHeader("Authorization", "Bearer "+env.AccessToken).Do()

	body, _ := res.Body.ToString()
	return body, nil
}

func main() {
	// Example
	//   env := Env{}
	//   envconfig.Process("", &env)
	//   resp, _ := Publish(env.ToID, "テスト", true)
	//   fmt.Printf("%v\n", resp)
}
