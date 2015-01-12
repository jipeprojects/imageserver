package png

import (
	"image"
	"testing"

	"github.com/pierrre/imageserver"
	imageserver_processor_native "github.com/pierrre/imageserver/processor/native"
)

func TestEncoder(t *testing.T) {
	nim := image.NewRGBA(image.Rect(0, 0, 100, 100))
	e := &Encoder{}
	_, err := e.Encode(nim, imageserver.Params{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncoderInterface(t *testing.T) {
	var _ imageserver_processor_native.Encoder = &Encoder{}
}
