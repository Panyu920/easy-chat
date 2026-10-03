package websocketx

type FrameType uint8

const (
	FrameTypeData FrameType = iota + 1
	FrameTypePing
)

type Message struct {
	FrameType FrameType `json:"frame_type"`
	Method    string    `json:"method"`
	FromID    string    `json:"from_id"`
	Data      any       `json:"data"`
}

func NewMessage(frameType FrameType, method string, fromID string, data any) *Message {
	return &Message{
		FrameType: frameType,
		Method:    method,
		FromID:    fromID,
		Data:      data,
	}
}

func NewErrMessage(data string) Message {
	return Message{
		FrameType: FrameTypeData,
		Data:      data,
	}
}
