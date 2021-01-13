package decode

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type PacketFunc func(r io.Reader) (interface{}, error)

func VarInt(r io.Reader) (int32, error) {
	res := int32(0)
	for n := 0; n < 6; n++ {
		read := []byte{0}
		_, err := r.Read(read)
		if err != nil {
			return 0, err
		}
		v := int32(read[0] & 0b01111111)
		res |= v << (7 * n)
		if (read[0] & 0b10000000) == 0 {
			return res, nil
		}
	}
	return 0, errors.New("varint too big")
}

func String(r io.Reader) (string, error) {
	length, err := VarInt(r)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer([]byte{})
	n, err := buf.ReadFrom(io.LimitReader(r, int64(length)))
	if err != nil {
		return "", err
	}
	if int32(n) != length {
		return "", fmt.Errorf("mismatched string length: wanted %d, have %d", length, n)
	}
	return buf.String(), nil
}

func UShort(r io.Reader) (res uint16, err error) {
	err = binary.Read(r, binary.BigEndian, &res)
	return
}

func UByte(r io.Reader) (res uint8, err error) {
	err = binary.Read(r, binary.BigEndian, &res)
	return
}

func Bool(r io.Reader) (res bool, err error) {
	b, err := UByte(r)
	if b == 0x01 {
		res = true
	}
	return
}

func Long(r io.Reader) (res int64, err error) {
	err = binary.Read(r, binary.BigEndian, &res)
	return
}

func Double(r io.Reader) (res float64, err error) {
	err = binary.Read(r, binary.BigEndian, &res)
	return
}
