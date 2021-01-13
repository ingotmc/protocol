package play

import (
	"github.com/ingotmc/protocol/encode"
	"io"
)

type UpdateViewPosition struct {
	ChunkX, ChunkZ int32
}

func (u UpdateViewPosition) ID() int32 {
	return 0x41
}

func (u UpdateViewPosition) EncodeMC(w io.Writer) error {
	encode.VarInt(u.ChunkX, w)
	encode.VarInt(u.ChunkZ, w)
	return nil
}
