package protocol

import (
	"bytes"
	"github.com/ingotmc/protocol/decode"
	"github.com/ingotmc/protocol/encode"
	"io"
	"io/ioutil"
)

// ReadPacketFunc describes a function which reads a packet id + payload from an io.Reader
type ReadPacketFunc func (r io.Reader) (id int32, data []byte, err error)

// SendPacketFunc describes a function which encodes a clientbound packet to a io.Writer
type SendPacketFunc func (pkt Clientbound, w io.Writer) error

func SendPacket(pkt Clientbound, w io.Writer) (err error) {
	buf := bytes.NewBuffer([]byte{})
	err = encode.VarInt(pkt.ID(), buf)
	if err != nil {
		return
	}
	err = pkt.EncodeMC(buf)
	if err != nil {
		return
	}
	err = encode.VarInt(int32(buf.Len()), w)
	if err != nil {
		return
	}
	_, err = io.CopyN(w, buf, int64(buf.Len()))
	return
}

func ReadWirePacket(r io.Reader) (id int32, data []byte, err error) {
	l, err := decode.VarInt(r)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer([]byte{})
	buf.Grow(int(l))
	_, err = io.CopyN(buf, r, int64(l))
	if err != nil {
		return
	}
	id, err = decode.VarInt(buf)
	if err != nil {
		return
	}
	data, err = ioutil.ReadAll(buf)
	return
}