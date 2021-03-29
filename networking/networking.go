package networking

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"godswar/decode"
	"godswar/logger"
	"godswar/opcodes"
	"net"
)

const MaxBuffer = 8192
var hashPointer = 0

type network struct {
	conn net.Conn
}

type Network interface {
	HandleConnection()
	Disconnect()
	Send([]byte, *int)
}

func NewNetwork(conn net.Conn) Network  {
	return &network{conn}
}

func (n network) Disconnect() {
	logger.BasicLog("Client Disconnect")
	n.conn.Close()
}

func (n network) Send(m []byte, hashPointer *int) {
	logger.BasicLog("Sending data to client")
	fmt.Println(hex.Dump(m))
	n.conn.Write(decode.Crypt(m, hashPointer))
}

func (n network) HandleConnection()  {
	// Make a buffer to hold incoming data.
	logger.BasicLog("Client Connected")
	defer n.Disconnect()
	buffer := make([]byte, MaxBuffer)

	hashPointer := 0
	sentHashPointer := 0
	for {
		/*
			BUFF COUNT = PACKET LEN = RECV DATA
		 */
		reader := bufio.NewReaderSize(n.conn, MaxBuffer)
		packet, err := decode.NewDecoder(reader, buffer, &hashPointer)
		if err != nil {
			n.Disconnect()
			break
		}

		if packet.Len == 0 || packet.Len > MaxBuffer {
			n.Disconnect()
			break
		}

		logger.BasicLog("OPCode:", packet.OPCode)
		logger.BasicLog("Buff Count:", packet.Len)
		logger.BasicLog("Data Size:", packet.DataSize)
		logger.BasicLog("Buffer")
		fmt.Println(hex.Dump(packet.Buffer))

		if packet.Len > 4 {
			opch := opcodes.NewOPCodeHandler(packet)
			res, err := opch.HandleOPCode()
			if err != nil {
				n.Disconnect()
			}

			if res != nil {
				n.Send(res, &sentHashPointer)
			}
		}
	}
}