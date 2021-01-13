package play

import (
	"github.com/ingotmc/protocol/decode"
	"io"
)

type PlayerPosition struct {
	X, FeetY, Z float64
	OnGround    bool
}

func readPlayerPosition(r io.Reader) (out interface{}, err error) {
	ppos := PlayerPosition{}
	ppos.X, err = decode.Double(r)
	if err != nil {
		return
	}
	ppos.FeetY, err = decode.Double(r)
	if err != nil {
		return
	}
	ppos.Z, err = decode.Double(r)
	if err != nil {
		return
	}
	ppos.OnGround, err = decode.Bool(r)
	if err != nil {
		return
	}
	out = ppos
	return
}
