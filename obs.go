package muniverse

import (
	"bytes"
	"image"
	"image/png"
)

// Obs is an observation from an environment.
type Obs interface {
	Image() (image.Image, error)
}

type pngObs []byte

func (p pngObs) Image() (image.Image, error) {
	return png.Decode(bytes.NewReader(p))
}
