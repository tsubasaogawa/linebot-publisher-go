# linebot-publisher-layer-go

The simplest line bot publisher module for golang. It supports sending only a text message.

## Usage

<!--
### Layer side

1. `make`
1. Upload zip includes binary to AWS Lambda Layer
-->

### Client side

1. Write the code using the publisher (*).
1. Upload function to AWS lambda

(*) Example

```go
import "github.com/tsubasaogawa/linebot-publisher-layer-go"
   :
// To LINE ID, LINE Access Token, Message, Notification Flag
linebot.Publish("abc123...", "def456...", "Test Message", false)
```

Information of needed arguments of `Publish()` are written in linebot.go.

## Image

![Image](https://raw.githubusercontent.com/tsubasaogawa/linebot-publisher-layer-go/images/image.png)
