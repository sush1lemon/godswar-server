package networking

import (
	"github.com/nats-io/nats.go"
	"godswar/pkg/decode"
	"godswar/pkg/logger"
	"net"
)

func NewConnection(conn net.Conn, recvHashPointer *int, sentHashPointer *int, listener *nats.Subscription) *Connection {
	return &Connection{
		n:               conn,
		recvHashPointer: recvHashPointer,
		sentHashPointer: sentHashPointer,
		listener: listener,
	}
}

type Connection struct {
	n net.Conn
	recvHashPointer *int
	sentHashPointer *int
	listener *nats.Subscription
	ServerID *int
}

func (c *Connection) Disconnect() {
	logger.BasicLog("Client disconnected")
	if c.listener != nil {
		c.listener.Unsubscribe()
	}
	c.n.Close()
}

func (c Connection) Send(m []byte) {
	c.n.Write(decode.Crypt(m, c.sentHashPointer))
}

func (c *Connection) AttachGameStateListener(l *nats.Subscription)  {
	c.listener = l
}
