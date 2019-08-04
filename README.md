# linebot-publisher-layer-go

The simplest line bot publisher module for golang. It supports sending only a text message.

## Usage

### Layer side

1. `make`
1. Upload zip includes binary to AWS Lambda Layer

### Client side

1. Write the code using the publisher. For example:

```go
import "github.com/tsubasaogawa/linebot-publisher-layer-go"
   :
linebot.Publish("abc123...", "def456...", "Hello", false)
```

1. Upload function to AWS lambda

## Image

TBA
