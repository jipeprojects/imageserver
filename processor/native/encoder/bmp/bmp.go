package bmp

import (
	"bytes"
	"image"

	"github.com/pierrre/imageserver"
	imageserver_processor_native "github.com/pierrre/imageserver/processor/native"
	"golang.org/x/image/bmp"
)

//Encoder encodes a native Image to a raw Image in bmp.
type Encoder struct {
}

//Encode encodes the image.
func (e *Encoder) Encode(nim image.Image, params imageserver.Params) (*imageserver.Image, error) {
	buf := new(bytes.Buffer)
	err := bmp.Encode(buf, nim)
	if err != nil {
		return nil, err
	}
	return &imageserver.Image{
		Format: "bmp",
		Data:   buf.Bytes(),
	}, nil
}

func init() {
	imageserver_processor_native.RegisterEncoder("bmp", &Encoder{})
}
