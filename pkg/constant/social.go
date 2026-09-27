package constant

type HandlerResult int

const (
	NoHandle HandlerResult = iota
	Pass
	Refuse
	Cancel
)
