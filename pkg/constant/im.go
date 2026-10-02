package constant

type MsgType uint8

const (
	MsgTypeText MsgType = iota + 1
	MsgTypePix
)

type ChatType uint8

const (
	ChatTypeSingle ChatType = iota + 1
	ChatTypeGroup
)
