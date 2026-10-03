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
	connToUser map[*Connection]string
	userToConn map[string]*Connection
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

		connToUser: make(map[*Connection]string),
		userToConn: make(map[string]*Connection),
	}
}

func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	// catch panic
	defer func() {
		if r := recover(); r != nil {
			s.Errorf("Panic in ServerWs: %v", r)
		}
	}()

	// 升级为 WebSocket 连接
	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Errorf("Failed to upgrade WebSocket connection: %v", err)
		return
	}
	connection := NewConnection(s, conn, s.serverOpt.maxConnIdleDuration)
	// 验证认证
	if !s.auth.Auth(w, r) {
		s.Errorf("Auth failed")
		s.Send("Auth failed", connection)
		s.closeConn(connection)
		return
	}
	// 添加连接
	s.AddConn(connection, r)
	// 设置用户ID
	connection.userID = s.auth.GetUserID(r)

	// 异步处理 WebSocket 连接
	go s.HandleConn(connection)
}

func (s *Server) HandleConn(conn *Connection) {
	// 处理 WebSocket 消息
	go conn.WriteMessage()
	conn.ReadMessage()
}

func (s *Server) RegisterRoutes(route []*Route) {
	for _, r := range route {
		s.routes[r.Method] = r.Handler
	}
}

func (s *Server) AddConn(conn *Connection, r *http.Request) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	user := s.auth.GetUserID(r)
	if conn, ok := s.userToConn[user]; ok {
		conn.Close()
	}
	s.connToUser[conn] = user
	s.userToConn[user] = conn
}

func (s *Server) GetConn(user string) (*Connection, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	conn, ok := s.userToConn[user]
	if !ok {
		s.Logger.Errorf("user %s not found", user)
		return nil, fmt.Errorf("user %s not found", user)
	}
	return conn, nil
}

func (s *Server) GetUser(conn *Connection) (string, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	user, ok := s.connToUser[conn]
	if !ok {
		return "", fmt.Errorf("conn %p not found", conn)
	}
	return user, nil
}

func (s *Server) GetConns(users ...string) []*Connection {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	var conns = make([]*Connection, 0, len(users))
	for _, user := range users {
		conn, ok := s.userToConn[user]
		if ok {
			conns = append(conns, conn)
		}
	}
	return conns
}

func (s *Server) GetUsers(conns ...*Connection) []string {
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

func (s *Server) Send(msg any, conns ...*Connection) error {
	if len(conns) == 0 {
		return nil
	}

	data, err := json.Marshal(msg)
	if err != nil {
		s.Errorf("Failed to marshal message: %v", err)
		return err
	}
	for _, conn := range conns {
		if err := conn.SendMessageToChan(data); err != nil {
			return err
		}
	}
	return nil
}
func (s *Server) SendToUser(msg any, user ...string) error {
	conns := s.GetConns(user...)
	return s.Send(msg, conns...)
}

func (s *Server) closeConn(conn *Connection) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	delete(s.userToConn, s.connToUser[conn])
	delete(s.connToUser, conn)
	conn.Close()
}

func (s *Server) Start() {
	http.HandleFunc(s.Patten, s.ServerWs)
	s.Info(http.ListenAndServe(s.Addr, nil))
}

func (s *Server) Stop() {
	s.Info("stop server")
}
