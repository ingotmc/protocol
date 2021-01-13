package login

import (
	"fmt"
	"github.com/ingotmc/protocol/decode"
)

func DecodeByID(id int32) (decode.PacketFunc, error) {
	switch id {
	case 0x00:
		return readLoginStart, nil
	default:
		return nil, fmt.Errorf("state login can't handle packet id %d", id)
	}
}