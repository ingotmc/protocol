package play

import (
	"github.com/ingotmc/protocol/decode"
	"io"
)

type ClientSettings struct {
	Locale string
	ViewDistance byte
	ChatMode int32
	ChatColors bool
	DisplayedSkinParts uint8
	MainHand int32
}

func readClientSettings(r io.Reader) (out interface{}, err error) {
	cs := ClientSettings{}
	cs.Locale, err = decode.String(r)
	cs.ViewDistance, err = decode.UByte(r)
	cs.ChatMode, err = decode.VarInt(r)
	cs.ChatColors, err = decode.Bool(r)
	cs.DisplayedSkinParts, err = decode.UByte(r)
	cs.MainHand, err = decode.VarInt(r)
	out = cs
	return
}

