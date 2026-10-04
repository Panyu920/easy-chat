package websocketx

import (
	"github.com/google/uuid"
)

type FrameType uint8

const (
	FrameTypeData FrameType = iota + 1
	FrameTypePing
	FrameTepeAck
)

type Result struct {
	ClientMsgId uuid.UUID
	Status      string
	Seq         uint32
}

type Message struct {
	FrameType   FrameType `json:"frame_type"`
	Seq         uint32    `json:"seq"`
	Method      string    `json:"method"`
	FromID      string    `json:"from_id"`
	ClientMsgId uuid.UUID `json:"client_msg_id"`
	ServerMsgId uuid.UUID `json:"server_msg_id"`
	Data        any       `json:"data"`
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

func NewAckMessage(results ...Result) Message {
	return Message{
		FrameType: FrameTepeAck,
		Data:      results,
	}
}
