package png

import (
	"bytes"
	"image"
	"image/png"

	"github.com/pierrre/imageserver"
	imageserver_processor_native "github.com/pierrre/imageserver/processor/native"
)

//Encoder encodes a native Image to a raw Image in png.
type Encoder struct {
}

var encoder = &png.Encoder{
	CompressionLevel: png.BestCompression,
}

//Encode encodes the image.
func (e *Encoder) Encode(nim image.Image, params imageserver.Params) (*imageserver.Image, error) {
	buf := new(bytes.Buffer)
	err := encoder.Encode(buf, nim)
	if err != nil {
		return nil, err
	}
	return &imageserver.Image{
		Format: "png",
		Data:   buf.Bytes(),
	}, nil
}

func init() {
	imageserver_processor_native.RegisterEncoder("png", &Encoder{})
}
