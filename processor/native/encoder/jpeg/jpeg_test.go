package jpeg

import (
	"image"
	"testing"

	"github.com/pierrre/imageserver"
	imageserver_processor_native "github.com/pierrre/imageserver/processor/native"
)

func TestEncoder(t *testing.T) {
	nim := image.NewRGBA(image.Rect(0, 0, 100, 100))
	e := &Encoder{}
	_, err := e.Encode(nim, imageserver.Params{"quality": 85})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncoderDefaultQuality(t *testing.T) {
	nim := image.NewRGBA(image.Rect(0, 0, 100, 100))
	e := &Encoder{
		DefaultQuality: 85,
	}
	_, err := e.Encode(nim, imageserver.Params{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncoderDefaultQuality2(t *testing.T) {
	nim := image.NewRGBA(image.Rect(0, 0, 100, 100))
	e := &Encoder{}
	_, err := e.Encode(nim, imageserver.Params{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncoderErrorQuality(t *testing.T) {
	nim := image.NewRGBA(image.Rect(0, 0, 100, 100))
	e := &Encoder{}
	for _, quality := range []int{-1, 101} {
		_, err := e.Encode(nim, imageserver.Params{"quality": quality})
		if err == nil {
			t.Fatal("no error")
		}
		errParam, ok := err.(*imageserver.ParamError)
		if !ok {
			t.Fatal("wrong error type")
		}
		if errParam.Param != "quality" {
			t.Fatal("wrong param")
		}
	}
}

func TestEncoderInterface(t *testing.T) {
	var _ imageserver_processor_native.Encoder = &Encoder{}
}
