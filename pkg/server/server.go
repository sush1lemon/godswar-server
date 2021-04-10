package server

import (
	"bufio"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/upper/db/v4"
	"godswar/pkg/decode"
	"godswar/pkg/logger"
	"godswar/pkg/networking"
	"godswar/pkg/opcodes"
	"godswar/pkg/services"
	"net"
	"os"
)

const MaxBuffer = 8196

type Server struct {
	Db     db.Session
	Config Config
}

type Config struct {
	Host string
	Port string
	Type string
}


func (s Server) NewServer()  {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logger.BasicLog("Cannot connect to nats server")
		os.Exit(1)
	}
	c, err := nats.NewEncodedConn(nc, nats.GOB_ENCODER)
	if err != nil {
		logger.BasicLog("Cannot connect to nats server")
		os.Exit(1)
	}
	defer c.Close()
	defer nc.Close()

	l, err := net.Listen(s.Config.Type, fmt.Sprintf("%s:%s", s.Config.Host, s.Config.Port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()
	logger.BasicLog("Listening on " + s.Config.Host + ":" + s.Config.Port)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go s.handleConnection(conn, c)
	}
}

func (s Server) handleConnection(n net.Conn, rdb *nats.EncodedConn)  {

	logger.BasicLog("New socket connected")

	recvHashPointer := 0
	sentHashPointer := 0
	conn := networking.NewConnection(n, &recvHashPointer, &sentHashPointer, nil)
	service := services.NewService(s.Db, conn)
	buffer := make([]byte, MaxBuffer)

	for {
		/*
			BUFF COUNT = PACKET LEN = RECV DATA
		*/
		reader := bufio.NewReaderSize(n, MaxBuffer)
		packet, err := decode.NewDecoder(reader, buffer, &recvHashPointer)
		if err != nil {
			conn.Disconnect()
			break
		}

		if packet.Len == 0 || packet.Len > MaxBuffer {
			conn.Disconnect()
			break
		}
		//logger.BasicLog("OPCode:", packet.OPCode)
		//logger.BasicLog("Buff Count:", packet.Len)
		//logger.BasicLog("Buffer")
		//fmt.Println(hex.Dump(packet.Buffer))

		if packet.Len > 4 {
			opch := opcodes.NewOPCodeHandler(packet, conn, service, rdb)
			_, err := opch.HandleOPCode()
			if err != nil {
				conn.Disconnect()
			}
		}
	}

}