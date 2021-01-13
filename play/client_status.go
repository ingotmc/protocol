package play

import (
	"github.com/ingotmc/protocol/decode"
	"io"
)

type ClientStatus struct {
	ActionID int32
}

func readClientStatus(r io.Reader) (out interface{}, err error) {
	cs := ClientStatus{}
	cs.ActionID, err = decode.VarInt(r)
	out = cs
	return
}

