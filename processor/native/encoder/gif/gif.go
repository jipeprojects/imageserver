package gif

import (
	"bytes"
	"image"
	"image/gif"

	"github.com/pierrre/imageserver"
	imageserver_processor_native "github.com/pierrre/imageserver/processor/native"
)

//Encoder encodes a native Image to a raw Image in gif.
type Encoder struct {
}

//Encode encodes the image.
func (e *Encoder) Encode(nim image.Image, params imageserver.Params) (*imageserver.Image, error) {
	buf := new(bytes.Buffer)
	err := gif.Encode(buf, nim, &gif.Options{})
	if err != nil {
		return nil, err
	}
	return &imageserver.Image{
		Format: "gif",
		Data:   buf.Bytes(),
	}, nil
}

func init() {
	imageserver_processor_native.RegisterEncoder("gif", &Encoder{})
}
