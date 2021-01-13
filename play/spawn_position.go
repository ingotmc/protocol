package play

import (
	"github.com/ingotmc/mc"
	"github.com/ingotmc/protocol/encode"
	"io"
)

type SpawnPosition struct {
	Position mc.Coords
}

func (s SpawnPosition) ID() int32 {
	return 0x4e
}

func (s SpawnPosition) EncodeMC(w io.Writer) (err error) {
	return encode.Coords(s.Position, w)
}
