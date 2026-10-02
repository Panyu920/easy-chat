package websocketx

type MessageType uint8

const (
	MessageTypeData MessageType = iota + 1
	MessageTypePing
)

type Message struct {
	MsgType MessageType `json:"msg_type"`
	Method  string      `json:"method"`
	FromID  string      `json:"from_id"`
	Data    any         `json:"data"`
}

func NewMessage(msgType MessageType, method string, fromID string, data any) *Message {
	return &Message{
		MsgType: msgType,
		Method:  method,
		FromID:  fromID,
		Data:    data,
	}
}
