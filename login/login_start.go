package login

import (
	"github.com/ingotmc/protocol/decode"
	"io"
)

type LoginStart struct {
	Username string
}

func readLoginStart(r io.Reader) (out interface{}, err error) {
	ls := LoginStart{}
	ls.Username, err = decode.String(r)
	if err != nil {
		return
	}
	out = ls
	return
}

