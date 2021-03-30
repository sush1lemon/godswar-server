package opcodes

import (
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
		h.conn.Send(packets.AFTER_LOGIN_TWO)
		break
	default:
		logger.BasicLog("Invalid OPCode:", h.decoded.OPCode)
		logger.BasicLog("Buff Count:", h.decoded.Len)
		return []byte{0x00}, nil
	}

	return nil, nil
}
