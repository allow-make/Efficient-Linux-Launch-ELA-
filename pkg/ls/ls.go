package ls

import "ela/pkg/bridge"

const (
TypeText       = 0
TypeFramebuffer = 1
TypeUnknown    = 2
)

func Init(channelID string) int {
return bridge.LSInit(channelID)
}

func Capture() string {
return bridge.LSCapture()
}

func DetectType(data []byte) int {
return bridge.LSDetectType(data)
}

func Send(channelID string, data []byte) int {
return bridge.LSSend(channelID, data)
}
