package websocketx

import "github.com/gorilla/websocket"

type HandlerFunc func(server *Server, conn *websocket.Conn, msg *Message)
type Route struct {
	Method  string
	Handler HandlerFunc
}
