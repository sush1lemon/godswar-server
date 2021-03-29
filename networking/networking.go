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

const MAX_BUFFER = 8192
var hashPointer = 0

type network struct {
	conn net.Conn
}

type Network interface {
	HandleConnection()
	Disconnect()
	Send([]byte)
}

func NewNetwork(conn net.Conn) Network  {
	return &network{conn}
}

func (n network) Disconnect() {
	logger.BasicLog("Client Disconnect")
	n.conn.Close()
}

func (n network) Send(m []byte) {
	n.conn.Write(decode.Crypt(m))
}

func (n network) HandleConnection()  {
	// Make a buffer to hold incoming data.
	logger.BasicLog("Client Connected")
	defer n.Disconnect()
	buffer := make([]byte, 1024)

	hashPointer := 0
	for {
		/*
			BUFF COUNT = PACKET LEN = RECV DATA
		 */
		reader := bufio.NewReaderSize(n.conn, MAX_BUFFER)
		packet, err := decode.NewDecoder(reader, buffer, &hashPointer)
		if err != nil {
			n.Disconnect()
			break
		}

		if packet.Len == 0 || packet.Len > MAX_BUFFER {
			n.Disconnect()
			break
		}

		fmt.Println(buffer[0:10])

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

			n.Send(res)
		}
	}
}