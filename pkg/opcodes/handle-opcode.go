package opcodes

import (
	"encoding/hex"
	"fmt"
	"github.com/nats-io/nats.go"
	"godswar/pkg/decode"
	"godswar/pkg/experimental"
	"godswar/pkg/logger"
	"godswar/pkg/networking"
	"godswar/pkg/packets"
	"godswar/pkg/services"
	"time"
)

type handler struct {
	decoded decode.Decode
	conn    *networking.Connection
	backend *services.Service
	rdb *nats.EncodedConn
}

type OPCode interface {
	HandleOPCode() ([]byte, error)
}

func NewOPCodeHandler(decoded decode.Decode, conn *networking.Connection, backend *services.Service, rdb *nats.EncodedConn) OPCode {
	return &handler{decoded, conn, backend, rdb}
}


func (h handler) HandleOPCode() ([]byte, error) {

	//reader := utility.NewPacketReader(h.decoded.DecodedBuffer)
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
		h.conn.Send(packets.AFTER_LOGIN)
		go h.backend.Account.GetAccountCharacters(&h.decoded)
		break
	case MSG_CREATE_ROLE:
		h.backend.Account.CreateAccountCharacter(&h.decoded)
		break
	case MSG_ENTER_GAME:
		l, _ := h.rdb.Subscribe("GAME:MESSAGING", func(p decode.Decode) {
			fmt.Println(hex.Dump(p.Buffer))
			h.conn.Send(p.Buffer)
		})

		h.conn.AttachGameStateListener(l)
		//h.conn.Send(experimental.ENTER_PART4)

		h.conn.Send(*h.conn.CharacterBytes)
		//h.conn.Send(experimental.ENTER_PART1)
		h.conn.Send(experimental.ENTER_PART2_UNK)
		h.conn.Send(experimental.ENTER_PART2)
		h.conn.Send(experimental.ENTER_PART4)

		//h.conn.Send(experimental.PHD1)
		//h.conn.Send(experimental.PHD2)
		//h.conn.Send(experimental.PHD3)
		break
	case 10015:
		/*
			PING
		*/
		h.conn.Send(h.decoded.Buffer)
		break
	case 10035:
		h.rdb.Publish("GAME:MESSAGING", h.decoded)
		break
	case MSG_WALK:
		h.rdb.Publish("GAME:MESSAGING", h.decoded)
		/*
			Walk ?
		*/
		//_ = reader.Read(2)
		//_ = reader.Read(1)
		//reader.Skip(1)
		//
		//x := reader.Read(4)
		//y := reader.Read(4)
		//mp := reader.Read(4)
		//
		//coorx := utility.Float64frombytes(x)
		//coory := utility.Float64frombytes(y)
		//coorz := utility.Float64frombytes(mp)
		//
		//fmt.Println(hex.Dump(x))
		//fmt.Println(hex.Dump(y))
		//fmt.Println(hex.Dump(mp))
		//fmt.Println("x:",coorx, "y:", coory, "mp:", coorz)
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
	case 10117:
		// Forge
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
