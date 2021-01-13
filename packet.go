package protocol

import (
	"fmt"
	"github.com/ingotmc/protocol/decode"
	"github.com/ingotmc/protocol/encode"
	"github.com/ingotmc/protocol/handshaking"
	"github.com/ingotmc/protocol/login"
	"github.com/ingotmc/protocol/play"
)

// Clientbound represents a packet that can be sent to a client.
type Clientbound interface {
	ID() int32
	encode.Encoder
}

func DecodeByStateID(state State, id int32) (decode.PacketFunc, error) {
	switch state {
	case Handshaking:
		return handshaking.DecodeByID(id)
	case Login:
		return login.DecodeByID(id)
	case Play:
		return play.DecodeByID(id)
	default:
		return nil, fmt.Errorf("unknown state %d", state)
	}
}