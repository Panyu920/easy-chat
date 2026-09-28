package constant

type HandlerResult int
const (
	NoHandle HandlerResult = iota
	Pass
	Refuse
	Cancel
)

type RoleLevel int
const (
	NormalMember RoleLevel = iota 
	Admin
	Owner
)
