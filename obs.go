package muniverse

import (
	"bytes"
	"image"
	"image/png"

	"github.com/unixpickle/essentials"
)

// Obs is an observation from an environment.
type Obs interface {
	Image() (image.Image, error)
}

// RGB generates an RGB pixel buffer for an observation.
//
// The buffer is packed the same way as the pixel data in
// *image.RGBA, but without an alpha channel.
func RGB(o Obs) (buffer []uint8, width, height int, err error) {
	img, err := o.Image()
	if err != nil {
		return
	}
	// TODO: optimize for *image.NRGBA and *image.RGBA.
	switch img := img.(type) {
	default:
		width, height = img.Bounds().Dx(), img.Bounds().Dy()
		min := img.Bounds().Min
		buffer = make([]uint8, width*height*3)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				r, g, b, _ := img.At(x+min.X, y+min.Y).RGBA()
				idx := (x + y*width) * 3
				buffer[idx] = uint8(r >> 8)
				buffer[idx+1] = uint8(g >> 8)
				buffer[idx+2] = uint8(b >> 8)
			}
		}
	}
	return
}

type pngObs []byte

func (p pngObs) Image() (img image.Image, err error) {
	defer essentials.AddCtxTo("decode PNG observation", &err)
	return png.Decode(bytes.NewReader(p))
}
