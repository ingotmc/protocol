package play

import (
	"github.com/ingotmc/protocol/encode"
	"io"
)

type ChunkData struct {
	ChunkX, ChunkZ      int32
	FullChunk           bool
	PrimaryBitMask      encode.Encoder
	Heightmaps          encode.Encoder
	Biomes              []int32
	ChunkContent        encode.Encoder
	NumberBlockEntities int32
	BlockEntities       []encode.Encoder
}

func (c ChunkData) ID() int32 {
	return 0x22
}

func (c ChunkData) EncodeMC(w io.Writer) (err error) {
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
	check(encode.Int(c.ChunkX, w))
	check(encode.Int(c.ChunkZ, w))
	check(encode.Bool(c.FullChunk, w))
	check(c.PrimaryBitMask.EncodeMC(w))
	check(c.Heightmaps.EncodeMC(w))
	if c.FullChunk {
		for _, b := range c.Biomes {
			check(encode.Int(b, w))
		}
	}
	check(c.ChunkContent.EncodeMC(w))
	check(encode.VarInt(c.NumberBlockEntities, w))
	if c.NumberBlockEntities == 0 {
		return err
	}
	for _, b := range c.BlockEntities {
		check(b.EncodeMC(w))
	}
	return err
}