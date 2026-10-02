package websocketx

import (
	"easy-chat/pkg/lock"
	"encoding/json"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Connection struct {
	conn                *websocket.Conn // 连接
	idleLock            sync.Locker     // 空闲锁
	idleAt              time.Time
	maxConnIdleDuration time.Duration
	s                   *Server
	userID              string

	done    chan struct{}
	msgChan chan []byte
	once    sync.Once
}

func NewConnection(s *Server, conn *websocket.Conn, maxConnIdleDuration time.Duration) *Connection {

	res := &Connection{
		s:                   s,
		conn:                conn,
		idleAt:              time.Now(),
		idleLock:            lock.NewSpinLock(),
		maxConnIdleDuration: maxConnIdleDuration,
		done:                make(chan struct{}),
		msgChan:             make(chan []byte, 10),
	}

	go res.KeepAlive()
	return res
}

func (c *Connection) KeepAlive() {
	// 创建定时器
	idleTimer := time.NewTimer(c.maxConnIdleDuration)

	defer idleTimer.Stop()
	for {
		select {
		case <-c.done:
			return
		case <-idleTimer.C:
			c.idleLock.Lock()
			idleAt := c.idleAt
			c.idleLock.Unlock()
			val := c.maxConnIdleDuration - time.Since(idleAt)
			if val <= 0 {
				c.s.closeConn(c)
				return
			}
			idleTimer.Reset(val)
		}
	}
}

func (c *Connection) Close() error {
	var err error
	c.once.Do(func() {
		close(c.done)
		// close(c.msgChan)
		err = c.conn.Close()
	})
	return err
}

func (c *Connection) ReadMessage() error {
	// 处理 WebSocket 消息
	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			c.s.Errorf("Failed to read message from WebSocket connection: %v", err)
			// 关闭连接
			c.s.closeConn(c)
			return err
		}
		c.updateIdleAt()

		var message Message
		if err := json.Unmarshal(data, &message); err != nil {
			c.s.Errorf("Failed to parse message: %v, message: %s", err, string(data))
			// 关闭连接
			c.s.closeConn(c)
			return err
		}
		// 调用路由处理函数
		if handler, ok := c.s.routes[message.Method]; ok {
			handler(c.s, c, &message)
		} else {
			c.s.Errorf("Unknown method: %s", message.Method)
			msg := NewMessage(FrameTypePing, message.Method, message.FromID, "unknown_method")
			data, err := json.Marshal(msg)
			if err != nil {
				c.s.Errorf("Failed to marshal message: %v", err)
				continue
			}
			c.SendMessageToChan(data)
		}
	}
}

func (c *Connection) SendMessageToChan(msg []byte) error {
	c.msgChan <- msg
	return nil
}

func (c *Connection) WriteMessage() error {
	for {
		select {
		case <-c.done:
			return nil
		case msg := <-c.msgChan:
			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				c.s.Errorf("Failed to write message to WebSocket connection: %v", err)
				// 关闭连接
				c.s.closeConn(c)
				return err
			}
			c.updateIdleAt()
		}
	}
}

func (c *Connection) updateIdleAt() {
	c.idleLock.Lock()
	c.idleAt = time.Now()
	c.idleLock.Unlock()
}

func (c *Connection) GetUserID() string {
	return c.userID
}
