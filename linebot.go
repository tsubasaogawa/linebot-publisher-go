package linebot

import (
	"fmt"

	"github.com/franela/goreq"
	"github.com/kelseyhightower/envconfig"
)

const (
	endpoint = "https://api.line.me/v2/bot/message/push"
)

// Messages is
type Messages struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Payload is
type Payload struct {
	To                   string     `json:"to"`
	Messages             []Messages `json:"messages"`
	NotificationDisabled bool       `json:"notificationDisabled"`
}

// Env is
type Env struct {
	AccessToken string `required:"true"`
	ToID        string
}

// Publish is ...
func Publish(toID string, message string, notifies bool) (string, error) {
	env := Env{}
	if err := envconfig.Process("", &env); err != nil {
		fmt.Errorf("Error: %s", err)
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
	env := Env{}
	envconfig.Process("", &env)
	resp, _ := Publish(env.ToID, "テスト", true)
	fmt.Printf("%v\n", resp)
}
