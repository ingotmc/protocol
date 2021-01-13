package encode

import (
	"encoding/binary"
	"github.com/ingotmc/mc"
	"io"
	"math"
)

// Encoder represents the ability to encode itself in a minecraft protocol compliant way.
type Encoder interface {
	EncodeMC(w io.Writer) error
}

type EncodeFunc func(w io.Writer) error

func (ef EncodeFunc) EncodeMC(w io.Writer) error {
	return ef(w)
}

func VarInt(value int32, w io.Writer) error {
	i := 0
	v := uint32(value)
	for {
		b := byte(v & 0x7F)
		v >>= 7
		if v != 0 {
			b |= 0x80
		}
		_, err := w.Write([]byte{b})
		if err != nil {
			return err
		}
		i++
		if v == 0 {
			break
		}
	}
	return nil
}

func String(s string, w io.Writer) error {
	l := int32(len(s))
	err := VarInt(l, w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(s))
	return err
}

func Int(i int32, w io.Writer) error {
	return binary.Write(w, binary.BigEndian, &i)
}

func Long(i int64, w io.Writer) error {
	return binary.Write(w, binary.BigEndian, &i)
}

func Bool(b bool, w io.Writer) error {
	v := byte(0x00)
	if b {
		v = 0x01
	}
	_, err := w.Write([]byte{v})
	return err
}

func Double(f float64, w io.Writer) error {
	return binary.Write(w, binary.BigEndian, &f)
}

func Short(f int16, w io.Writer) error {
	return binary.Write(w, binary.BigEndian, &f)
}

func Float(f float32, w io.Writer) error {
	return binary.Write(w, binary.BigEndian, &f)
}

func UByte(b uint8, w io.Writer) error {
	_, err := w.Write([]byte{b})
	return err
}

func Coords(pos mc.Coords, w io.Writer) error {
	x := int64(math.Floor(pos.X))
	y := int64(math.Floor(pos.Y))
	z := int64(math.Floor(pos.Z))
	v := ((x & 0x3ffffff) << 38) | ((z & 0x3ffffff) << 12) | (y & 0xfff)
	return Long(v, w)
}