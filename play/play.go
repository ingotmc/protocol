package play

import (
	"fmt"
	"github.com/ingotmc/protocol/decode"
)

func DecodeByID(id int32) (decode.PacketFunc, error) {
	switch id {
	case 0x00:
		return readTeleportConfirm, nil
	case 0x04:
		return readClientStatus, nil
	case 0x05:
		return readClientSettings, nil
	case 0x0f:
		return readKeepAlive, nil
	case 0x11:
		return readPlayerPosition, nil
	default:
		return nil, fmt.Errorf("unknown packet id %d for state play", id)
	}
}
