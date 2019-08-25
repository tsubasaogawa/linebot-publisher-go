# linebot-publisher-go

The simplest line bot publisher module for golang. It supports sending only a text message.

## Usage

```go
import "github.com/tsubasaogawa/linebot-publisher-go"
   :
// To LINE ID, LINE Access Token, Message, Notification Flag
linebot.Publish("abc123...", "def456...", "Test Message", false)
```

Information of needed arguments of `Publish()` are written in linebot.go.

## Image

![Image](https://raw.githubusercontent.com/tsubasaogawa/linebot-publisher-layer-go/images/image.png)
