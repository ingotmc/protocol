package play

import (
	"github.com/ingotmc/protocol/decode"
	"github.com/ingotmc/protocol/encode"
	"io"
)

type KeepAlive struct {
	KeepAliveID int64
}

func readKeepAlive(r io.Reader) (out interface{}, err error) {
	k := KeepAlive{}
	k.KeepAliveID, err = decode.Long(r)
	out = k
	return
}

func (k KeepAlive) ID() int32 {
	return 0x21
}

func (k KeepAlive) EncodeMC(w io.Writer) (err error) {
	return encode.Long(k.KeepAliveID, w)
}
