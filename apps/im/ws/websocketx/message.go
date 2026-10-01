package websocketx

type Message struct {
	Method string `json:"method"`
	FromID string `json:"from_id"`
	Data   any    `json:"data"`
}

func NewMessage(fromID string, data any) *Message {
	return &Message{
		FromID: fromID,
		Data:   data,
	}
}
