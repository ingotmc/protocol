package play

import (
	"github.com/ingotmc/mc"
	"github.com/ingotmc/protocol/encode"
	"io"
)

type PlayerPositionAndLook struct {
	Position   mc.Coords
	Rotation   mc.Rotation
	Relative   uint8
	TeleportID int32
}

func (p PlayerPositionAndLook) ID() int32 {
	return 0x36
}

func (p PlayerPositionAndLook) EncodeMC(w io.Writer) (err error) {
	encode.Double(p.Position.X, w)
	encode.Double(p.Position.Y, w)
	encode.Double(p.Position.Z, w)
	encode.Float(p.Rotation.Yaw, w)
	encode.Float(p.Rotation.Pitch, w)
	encode.UByte(p.Relative, w)
	encode.VarInt(p.TeleportID, w)
	return
}
