package opcodes

import (
	"encoding/hex"
	"fmt"
	"godswar/pkg/decode"
	"godswar/pkg/experimental"
	"godswar/pkg/logger"
	"godswar/pkg/networking"
	"godswar/pkg/packets"
	"godswar/pkg/services"
	"godswar/pkg/utility"
	"time"
)

type handler struct {
	decoded decode.Decode
	conn    networking.Connection
	backend *services.Service
}

type OPCode interface {
	HandleOPCode() ([]byte, error)
}

func NewOPCodeHandler(decoded decode.Decode, conn networking.Connection, backend *services.Service) OPCode {
	return &handler{decoded, conn, backend}
}


func (h handler) HandleOPCode() ([]byte, error) {

	reader := utility.NewPacketReader(h.decoded.DecodedBuffer)
	switch h.decoded.OPCode {
	case MSG_LOGIN:
		h.backend.Account.Login(&h.decoded)
		break
	case MSG_SELECT_SERVER:
		h.conn.Send(packets.SEND_SERVER)
		break
	case MSG_LOGIN_RETURN_INFO:
		time.Sleep(5000)
		h.conn.Send(packets.NEW_GAME_SERVER)
		break
	case MSG_LOGIN_GAMESERVER:
		time.Sleep(5000)
		h.conn.Send(packets.AFTER_LOGIN)
		time.Sleep(5000)
		h.conn.Send(packets.CHAMPTEST)
		break
	case MSG_CREATE_ROLE:
		h.conn.Send([]byte{0x0C, 0x00, 0xB4, 0x27, 0x13, 0x27, 0x8D, 0x0B, 0x01, 0x00, 0x00, 0x00})
		break
	case MSG_ENTER_GAME:
		/*
			request for files?
		*/
		h.conn.Send(packets.ENTER_PART1)
		h.conn.Send(packets.ENTER_PART4)
		break
	case 10015:
		/*
			Keep alive?
		*/
		h.conn.Send(h.decoded.Buffer)
		break
	case 10035:
		break
	case 10194:
		/*
			Walk ?
		*/
		_ = reader.Read(2)
		_ = reader.Read(1)
		reader.Skip(1)

		x := reader.Read(4)
		y := reader.Read(4)
		mp := reader.Read(4)

		coorx := utility.Float64frombytes(x)
		coory := utility.Float64frombytes(y)
		coorz := utility.Float64frombytes(mp)
		//
		//fmt.Println(hex.Dump(x))
		//fmt.Println(hex.Dump(y))
		//fmt.Println(hex.Dump(mp))
		fmt.Println("x:",coorx, "y:", coory, "mp:", coorz)
		//fmt.Println(mp)
		break
	case 10311:
		h.conn.Send(experimental.SERVER_TIME)
		break
	case 10007:
		/*
			request for files?
		*/
		//h.conn.Send(experimental.ATHENS_NPC)
		break
	case 10200:
		//h.conn.Send(experimental.SYNC_AP)
		break
	default:
		logger.BasicLog("Invalid OPCode:", h.decoded.OPCode)
		logger.BasicLog("Buff Count:", h.decoded.Len)
		fmt.Println(hex.Dump(h.decoded.Buffer))
		return []byte{0x00}, nil
	}

	return nil, nil
}
