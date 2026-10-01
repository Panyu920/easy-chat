package websocketx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type Server struct {
	Addr     string
	Upgrader websocket.Upgrader
	Patten   string
	logx.Logger

	routes map[string]HandlerFunc
	auth   IAuth

	mtx        sync.RWMutex
	connToUser map[*websocket.Conn]string
	userToConn map[string]*websocket.Conn
	serverOpt  *serverOption
}

func NewServer(addr string, opts ...ServerOptions) Server {
	opt := newServerOptions(opts...)
	return Server{
		Addr:      addr,
		Upgrader:  websocket.Upgrader{},
		Logger:    logx.WithContext(context.Background()),
		Patten:    opt.pattern,
		routes:    make(map[string]HandlerFunc),
		mtx:       sync.RWMutex{},
		auth:      opt.auth,
		serverOpt: opt,

		connToUser: make(map[*websocket.Conn]string),
		userToConn: make(map[string]*websocket.Conn),
	}
}

func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	// catch panic
	defer func() {
		if r := recover(); r != nil {
			s.Errorf("Panic in ServerWs: %v", r)
		}
	}()
	// 验证认证
	if !s.auth.Auth(w, r) {
		s.Errorf("Auth failed")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Auth failed"))
		return
	}

	// 升级为 WebSocket 连接
	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Errorf("Failed to upgrade WebSocket connection: %v", err)
		return
	}
	// 添加连接
	s.AddConn(conn, r)
	// 异步处理 WebSocket 连接
	go s.HandleConn(conn)
}

func (s *Server) HandleConn(conn *websocket.Conn) {
	// 处理 WebSocket 消息
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.Errorf("Failed to read message from WebSocket connection: %v", err)
			// 关闭连接
			s.closeConn(conn)
			return
		}

		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			s.Errorf("Failed to parse message: %v, message: %s", err, string(msg))
			// 关闭连接
			s.closeConn(conn)
			return
		}
		// 调用路由处理函数
		if handler, ok := s.routes[message.Method]; ok {
			handler(s, conn, &message)
		} else {
			s.Errorf("Unknown method: %s", message.Method)
			conn.WriteMessage(websocket.TextMessage, fmt.Appendf(nil, "Unknown method, please check the method %s", message.Method))
		}
	}
}

func (s *Server) RegisterRoutes(route []*Route) {
	for _, r := range route {
		s.routes[r.Method] = r.Handler
	}
}

func (s *Server) AddConn(conn *websocket.Conn, r *http.Request) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	user := s.auth.GetUser(r)
	s.connToUser[conn] = user
	s.userToConn[user] = conn
}

func (s *Server) GetConn(user string) (*websocket.Conn, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	conn, ok := s.userToConn[user]
	if !ok {
		return nil, fmt.Errorf("user %s not found", user)
	}
	return conn, nil
}

func (s *Server) GetUser(conn *websocket.Conn) (string, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	user, ok := s.connToUser[conn]
	if !ok {
		return "", fmt.Errorf("conn %p not found", conn)
	}
	return user, nil
}

func (s *Server) GetConns(users ...string) []*websocket.Conn {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	var conns = make([]*websocket.Conn, 0, len(users))
	for _, user := range users {
		conn, ok := s.userToConn[user]
		if ok {
			conns = append(conns, conn)
		}
	}
	return conns
}

func (s *Server) GetUsers(conns ...*websocket.Conn) []string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if len(conns) == 0 {
		// 如果没有指定连接，返回所有用户
		var users = make([]string, 0, len(s.userToConn))
		for user := range s.userToConn {
			users = append(users, user)
		}
		return users
	}

	var users = make([]string, 0, len(conns))
	for _, conn := range conns {
		user, ok := s.connToUser[conn]
		if ok {
			users = append(users, user)
		}
	}
	return users
}

func (s *Server) Send(msg any, conns ...*websocket.Conn) error {
	if len(conns) == 0 {
		return nil
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return err
		}
	}
	return nil
}
func (s *Server) SendToUser(msg any, user ...string) error {
	conns := s.GetConns(user...)
	return s.Send(msg, conns...)
}

func (s *Server) closeConn(conn *websocket.Conn) {
	conn.Close()
	s.mtx.Lock()
	defer s.mtx.Unlock()
	delete(s.connToUser, conn)
	delete(s.userToConn, s.connToUser[conn])
}

func (s *Server) Start() {
	http.HandleFunc(s.Patten, s.ServerWs)
	s.Info(http.ListenAndServe(s.Addr, nil))
}

func (s *Server) Stop() {
	s.Info("stop server")
}
