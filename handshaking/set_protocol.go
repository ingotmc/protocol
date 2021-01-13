package handshaking

import (
	"github.com/ingotmc/protocol/decode"
	"io"
)

type SetProtocol struct {
	ProtocolVersion int32
	ServerAddr string
	ServerPort uint16
	NextState int32
}

func readSetProtocol(r io.Reader) (out interface{}, err error) {
	sp := SetProtocol{}
	sp.ProtocolVersion, err = decode.VarInt(r)
	if err != nil {
		return
	}
	sp.ServerAddr, err = decode.String(r)
	if err != nil {
		return
	}
	sp.ServerPort, err = decode.UShort(r)
	if err != nil {
		return
	}
	sp.NextState, err = decode.VarInt(r)
	out = sp
	return
}
