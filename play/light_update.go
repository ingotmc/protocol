package play

import (
	"github.com/ingotmc/mc/light"
	"github.com/ingotmc/protocol/encode"
	"io"
)

type UpdateLight struct {
	ChunkX, ChunkZ int32
	SkyLightArrays [18]*light.Section
	BlockLightArrays [18]*light.Section
}

func (u UpdateLight) ID() int32 {
	return 0x25
}

func masksFromSections(sections [18]*light.Section) (mask, emptyMask int32) {
	for i, s := range sections {
		if s == nil {
			emptyMask |= int32(0x01) << i
			continue
		}
		mask |= int32(0x01) << i
	}
	return
}

func (u UpdateLight) EncodeMC(w io.Writer) (err error) {
	defer func() {
		if r := recover(); r != nil {
			e, ok := r.(error)
			if !ok {
				panic(r)
			}
			err = e
		}
	}()
	check := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	check(encode.VarInt(u.ChunkX, w))
	check(encode.VarInt(u.ChunkZ, w))
	skyLightMask, emptySkyLightMask := masksFromSections(u.SkyLightArrays)
	blockLightMask, emptyBlockLightMask := masksFromSections(u.BlockLightArrays)
	check(encode.VarInt(skyLightMask, w))
	check(encode.VarInt(blockLightMask, w))
	check(encode.VarInt(emptySkyLightMask, w))
	check(encode.VarInt(emptyBlockLightMask, w))
	//skyArrayLen := bits.OnesCount32(uint32(skyLightMask))
	//encode.VarInt(int32(2048 * skyArrayLen), w)
	for _, s := range u.SkyLightArrays {
		if s == nil {
			continue
		}
		check(encode.LightSection(*s).EncodeMC(w))
	}
	// TODO: implement blocklight
	return
}
