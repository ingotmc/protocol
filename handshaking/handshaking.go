package handshaking

import (
	"fmt"
	"github.com/ingotmc/protocol/decode"
)

func DecodeByID(id int32) (decode.PacketFunc, error) {
	switch id {
	case 0x00:
		return readSetProtocol, nil
	default:
		return nil, fmt.Errorf("unknown packet id %d for state handshaking", id)
	}
}
