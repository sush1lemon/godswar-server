package opcodes

import (
	"godswar/decode"
	"godswar/logger"
	"godswar/packets"
)

type handler struct {
	decoded decode.Decode
}

type OPCode interface {
	HandleOPCode() ([]byte, error)
}

func NewOPCodeHandler(decoded decode.Decode) OPCode {
	return &handler{decoded}
}

func (h handler) HandleOPCode() ([]byte, error) {

	//reader := utility.NewPacketReader(h.decoded.DecodedBuffer)
	switch h.decoded.OPCode {
	case MSG_LOGIN:
		//usernameShift := reader.Read(32)
		//password := reader.Read(32)
		//reader.Skip(4)
		//clientMac := reader.Read(32)
		//unk3 := reader.UInt32()
		//username := utility.Parser(usernameShift)
		//
		//logger.BasicLog("OPCode", h.decoded.OPCode)
		//logger.BasicLog("Username", username)
		//logger.BasicLog("Password", password)
		//logger.BasicLog("Client Mac", clientMac)
		//logger.BasicLog("UNK3", unk3)

		return packets.SERVER_LIST, nil
	case MSG_SELECT_SERVER:
		return packets.SEND_SERVER, nil
	case MSG_LOGIN_RETURN_INFO:
		return nil, nil
	default:
		logger.BasicLog("Invalid OPCode:", h.decoded.OPCode)
		logger.BasicLog("Buff Count:", h.decoded.Len)
		logger.BasicLog("Data Size:", h.decoded.DataSize)
		return []byte{0x00}, nil
	}

}
