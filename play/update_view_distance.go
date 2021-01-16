package play

import (
	"github.com/ingotmc/protocol/encode"
	"io"
)

type UpdateViewDistance struct {
	Distance int32
}

func (u UpdateViewDistance) ID() int32 {
	return 0x42
}

func (u UpdateViewDistance) EncodeMC(w io.Writer) error {
	return encode.VarInt(u.Distance, w)
}

