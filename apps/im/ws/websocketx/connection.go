package websocketx

import (
	"easy-chat/pkg/lock"
	"encoding/json"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Connection struct {
	conn                *websocket.Conn // 连接
	idleLock            sync.Locker     // 空闲锁
	idleAt              time.Time
	maxConnIdleDuration time.Duration
	CurrentAckSqe       uint32
	s                   *Server
	userID              string

	needHanldeMsgMap map[uuid.UUID]*Message
	msgLock          sync.Locker

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
			if errClose, ok := err.(*websocket.CloseError); ok {
				// 客户端关闭连接
				c.s.Errorf("client userId %v close: %v, err: %v", c.userID, errClose.Code, errClose.Text)
				// 关闭连接
				c.s.closeConn(c)
				return errClose
			}
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
		switch message.FrameType {
		case FrameTypePing:
		case FrameTypeData:
			{

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

func (c *Connection) AddNeedHandleMsg(msg *Message) bool {
	c.msgLock.Lock()
	defer c.msgLock.Unlock()
	currentMsg, ok := c.needHanldeMsgMap[msg.ClientMsgId]
	// 1.msg 有server_msg_id, map 中不存在 ,消息已处理
	if msg.ServerMsgId != uuid.Nil && !ok {
		return false
	}
	// 2.msg 有server_msg_id, map 中存在, server_msg_id 相同, 更新, 不同, 重新发送
	// 3.msg 没有server_msg_id,map 中不存在, 直接添加
	// 4.msg 没有server_msg_id,map 中存在,重新发送
	if ok {
		// 已存在, 只更新
		msg.ServerMsgId = currentMsg.ServerMsgId
		c.needHanldeMsgMap[msg.ClientMsgId] = msg
		return true
	} else {
		return false
	}
}
