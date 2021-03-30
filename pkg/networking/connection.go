package networking

import (
	"encoding/hex"
	"fmt"
	"godswar/pkg/decode"
	"godswar/pkg/logger"
	"net"
)

func NewConnection(conn net.Conn, recvHashPointer *int, sentHashPointer *int) Connection {
	return Connection{
		n:               conn,
		recvHashPointer: recvHashPointer,
		sentHashPointer: sentHashPointer,
	}
}

type Connection struct {
	n net.Conn
	recvHashPointer *int
	sentHashPointer *int
}

func (c Connection) Disconnect() {
	logger.BasicLog("Client disconnected")
	c.n.Close()
}

func (c Connection) Send(m []byte) {
	logger.BasicLog("Sending data to client")
	fmt.Println(hex.Dump(m))
	c.n.Write(decode.Crypt(m, c.sentHashPointer))
}
