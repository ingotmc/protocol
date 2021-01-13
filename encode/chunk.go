package encode

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/bearmini/bitstream-go"
	"github.com/ingotmc/mc"
	"github.com/ingotmc/nbt"
	"io"
	"math"
)

type chunkEncoder struct {
	c mc.Chunk
}

func (c chunkEncoder) EncodeMC(w io.Writer) error {
	return ChunkData(c.c, w)
}

func Chunk(chunk mc.Chunk) Encoder {
	return chunkEncoder{chunk}
}

func ChunkBitmask(chunk mc.Chunk) EncodeFunc {
	f := func(w io.Writer) error {
		bitMask := int32(0x0000)
		for i, s := range chunk.Sections {
			if s != nil {
				bitMask = int32(0x01)<<i | bitMask
			}
		}
		return VarInt(bitMask, w)
	}
	return f
}

func Heightmap(heightmap mc.Heightmap) EncodeFunc {
	return func(w io.Writer) error {
		buf := bytes.NewBuffer([]byte{})
		bw := bitstream.NewWriter(buf)
		for z := 0; z < 16; z++ {
			for x := 0; x < 16; x++ {
				bw.WriteNBitsOfUint16BE(9, heightmap.HeightAt(x, z))
			}
		}
		if buf.Len() != 64/8*36 {
			return errors.New("heighmap buf len isn't what expected")
		}
		b := buf.Bytes()
		data := make([]int64, 36)
		for i := range data {
			data[i] = int64(binary.BigEndian.Uint64(b[i*8 : (i*8)+8]))
		}
		return nbt.Encode(nbt.Compound{
			"MOTION_BLOCKING": data,
		}, w)
	}
}

func ChunkData(c mc.Chunk, w io.Writer) error {
	chunkData := bytes.NewBuffer([]byte{})
	for _, s := range c.Sections {
		if s == nil {
			continue
		}
		err := chunkSection(*s, chunkData)
		if err != nil {
			return err
		}
	}
	err := VarInt(int32(chunkData.Len()), w)
	if err != nil {
		return err
	}
	_, err = w.Write(chunkData.Bytes())
	return err
}

func chunkSection(s mc.Section, w io.Writer) error {
	blockCount := 0
	air, _ := mc.GlobalPalette.FindByName("minecraft:air")
	for _, b := range s {
		if b.ID == air.DefaultState.ID {
			blockCount++
		}
	}
	palette := []int32{
		0,
		1,
	}
	Short(int16(blockCount), w)
	bpb := int(math.Ceil(math.Log2(float64(len(palette)))))
	if bpb < 4 {
		bpb = 4
	}
	UByte(uint8(bpb), w)
	VarInt(int32(len(palette)), w)
	for _, id := range palette {
		VarInt(id, w)
	}
	blocks := make([]uint64, 64*bpb)
	for blockIdx, block := range s {
		longIdx := blockIdx >> 6
		max := int(math.Floor(64.0 / float64(bpb)))
		for i := 0; i < max; i++ {
			mask := uint8(0xff >> (8 - bpb))
			for pi, pb := range palette {
				if pb != block.ID {
					continue
				}
				blocks[longIdx] |= uint64(pi) & uint64(mask) << uint64(bpb * i)
				break
			}
		}
	}
	VarInt(int32(len(blocks)), w)
	for _, l := range blocks {
		Long(int64(l), w)
	}
	return nil
}
