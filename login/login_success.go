package login

import (
	"github.com/ingotmc/protocol/encode"
	"io"
)

type LoginSuccess struct {
	UUID, Username string
}

func (l LoginSuccess)  ID() int32 {
	return 0x02
}

func (l LoginSuccess) EncodeMC(w io.Writer) (err error) {
	err = encode.String(l.UUID, w)
	if err != nil {
		return
	}
	return encode.String(l.Username, w)
}
