package websocketx

import (
	"encoding/json"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type Client interface {
	Send(v any) error
	Close() error
	Read(v any) error
}

type client struct {
	conn *websocket.Conn
	host string
	opt  *dialoOption
}

func NewClient(host string, opts ...Dialoptions) Client {
	o := newDialoOption(opts...)
	c := &client{
		host: host,
		opt:  o,
	}

	err := c.Dial()
	if err != nil {
		panic(err)
	}

	return c
}
func (c *client) Dial() error {
	u := url.URL{Scheme: "ws", Host: c.host, Path: c.opt.pattern}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), c.opt.header)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *client) Send(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = c.conn.WriteMessage(websocket.TextMessage, data)
	if err == nil {
		logx.Infof("Send message to %s success", data)
		return nil
	}

	err = c.Dial()
	if err != nil {
		return err
	}

	return c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Read(v any) error {
	_, data, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
