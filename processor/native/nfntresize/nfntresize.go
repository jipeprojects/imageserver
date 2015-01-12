package nfntresize

import (
	"image"

	"github.com/nfnt/resize"
	"github.com/pierrre/imageserver"
)

type Processor struct {
}

func (proc *Processor) Process(nim image.Image, params imageserver.Params) (image.Image, error) {
	width := 100
	height := 0
	interp := resize.Lanczos3
	nim = resize.Resize(uint(width), uint(height), nim, interp)
	return nim, nil
}
