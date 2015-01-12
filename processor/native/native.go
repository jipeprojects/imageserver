package native

import (
	"bytes"
	"fmt"
	"image"

	"github.com/pierrre/imageserver"
)

/*
Processor is an Image Processor that uses the native Go Image

Steps:

- decode (from raw data to Go image)

- process (Go image)

- encode (from Go image to raw data)
*/
type Processor struct {
	Processor ProcessorNative
}

// Process processes an Image using native Go Image
func (processor *Processor) Process(im *imageserver.Image, params imageserver.Params) (*imageserver.Image, error) {
	nim, err := decode(im)
	if err != nil {
		return nil, err
	}
	nim, err = processor.Processor.Process(nim, params)
	if err != nil {
		return nil, err
	}
	im, err = encode(nim, im.Format, params)
	if err != nil {
		return nil, err
	}
	return im, nil
}

func decode(im *imageserver.Image) (image.Image, error) {
	nim, format, err := image.Decode(bytes.NewReader(im.Data))
	if err != nil {
		return nil, err
	}
	if format != im.Format {
		return nil, fmt.Errorf("image format \"%s\" does not match decoded format \"%s\"", im.Format, format)
	}
	return nim, nil
}

func encode(nim image.Image, format string, params imageserver.Params) (*imageserver.Image, error) {
	var err error
	if params.Has("format") {
		format, err = params.GetString("format")
		if err != nil {
			return nil, err
		}
	}
	encoder, ok := encoders[format]
	if !ok {
		return nil, &imageserver.ParamError{Param: "format", Message: fmt.Sprintf("no registered encoder for format \"%s\"", format)}
	}
	im, err := encoder.Encode(nim, params)
	if err != nil {
		return nil, err
	}
	return im, nil
}

// ProcessorNative processes a native Go Image
type ProcessorNative interface {
	Process(image.Image, imageserver.Params) (image.Image, error)
}

// ProcessorNativeFunc is a Processor func
type ProcessorNativeFunc func(image.Image, imageserver.Params) (image.Image, error)

// Process calls the func
func (f ProcessorNativeFunc) Process(nim image.Image, params imageserver.Params) (image.Image, error) {
	return f(nim, params)
}

// Encoder encodes a native Image to a raw Image
type Encoder interface {
	Encode(image.Image, imageserver.Params) (*imageserver.Image, error)
}

// EncoderFunc is a Encoder func
type EncoderFunc func(image.Image, imageserver.Params) (*imageserver.Image, error)

// Encode calls the func
func (f EncoderFunc) Encode(nim image.Image, params imageserver.Params) (*imageserver.Image, error) {
	return f(nim, params)
}

var encoders = make(map[string]Encoder)

// RegisterEncoder registers a Decoder for a format.
func RegisterEncoder(format string, encoder Encoder) {
	encoders[format] = encoder
}
