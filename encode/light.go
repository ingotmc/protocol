package encode

import (
	"github.com/ingotmc/mc/light"
	"io"
)

func LightSection(l light.Section) EncodeFunc {
	return func(w io.Writer) error {
		err := VarInt(2048, w)
		if err != nil {
			return err
		}
		_, err = w.Write(l[:])
		return err
	}
}
