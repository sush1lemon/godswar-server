package opcodes

import (
	"encoding/hex"
	"fmt"
	"godswar/pkg/decode"
	"godswar/pkg/logger"
	"godswar/pkg/networking"
	"godswar/pkg/packets"
	"godswar/pkg/utility"
)

type handler struct {
	decoded decode.Decode
	conn    networking.Connection
}

type OPCode interface {
	HandleOPCode() ([]byte, error)
}

func NewOPCodeHandler(decoded decode.Decode, conn networking.Connection) OPCode {
	return &handler{decoded, conn}
}

func (h handler) HandleOPCode() ([]byte, error) {

	reader := utility.NewPacketReader(h.decoded.DecodedBuffer)
	switch h.decoded.OPCode {
	case MSG_LOGIN:
		usernameShift := reader.Read(32)
		password := reader.Read(32)
		reader.Skip(4)
		clientMac := reader.Read(32)
		unk3 := reader.UInt32()
		username := utility.Parser(usernameShift)

		logger.BasicLog("OPCode", h.decoded.OPCode)
		logger.BasicLog("Username", username)
		logger.BasicLog("Password", password)
		logger.BasicLog("Client Mac", clientMac)
		logger.BasicLog("UNK3", unk3)

		h.conn.Send(packets.SERVER_LIST)
		break
	case MSG_SELECT_SERVER:
		h.conn.Send(packets.SEND_SERVER)
		break
	case MSG_LOGIN_RETURN_INFO:
		h.conn.Send(packets.NEW_GAME_SERVER)
		break
	case MSG_LOGIN_GAMESERVER:
		h.conn.Send(packets.AFTER_LOGIN)
		h.conn.Send(packets.CHAMP)
		break
	case MSG_CREATE_ROLE:
		h.conn.Send([]byte{0x0C, 0x00, 0xB4, 0x27, 0x13, 0x27, 0x8D, 0x0B,  0x01, 0x00, 0x00, 0x00})
		break
	case MSG_ENTER_GAME:
		h.conn.Send(packets.ENTER_PART1)
		h.conn.Send(packets.ENTER_PART2)
		h.conn.Send(packets.ENTER_PART3)
		h.conn.Send(packets.ENTER_PART4)
		break
	case 10015:
		/*
			Keep alive?
		 */
		h.conn.Send(h.decoded.Buffer)
		break
	case 10194:
		/*
			Walk ?
		 */
		break
	default:
		logger.BasicLog("Invalid OPCode:", h.decoded.OPCode)
		logger.BasicLog("Buff Count:", h.decoded.Len)
		fmt.Println(hex.Dump(h.decoded.Buffer))
		return []byte{0x00}, nil
	}

	return nil, nil
}
