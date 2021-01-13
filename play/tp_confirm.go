package play

import (
	"github.com/ingotmc/protocol/decode"
	"io"
)

type TeleportConfirm struct {
	TeleportID int32
}

func readTeleportConfirm(r io.Reader) (out interface{}, err error) {
	tp := TeleportConfirm{}
	tp.TeleportID, err = decode.VarInt(r)
	out = tp
	return
}

