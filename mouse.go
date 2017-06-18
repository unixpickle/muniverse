package muniverse

import (
	"image"
	"image/color"
	"time"

	"github.com/unixpickle/muniverse/chrome"
)

const (
	mouseSize = 13
	mouseStem = 5
)

type mouseEnv struct {
	Env

	curX int
	curY int
}

// MouseEnv creates a wrapped environment which renders a
// pointer at the current mouse location.
//
// By default, Chrome will not render a mouse.
// Thus, it is necessary to render a mouse manually.
//
// When the resulting Env is closed, e is closed as well.
func MouseEnv(e Env) Env {
	return &mouseEnv{Env: e}
}

func (m *mouseEnv) Reset() error {
	m.curX = -20
	m.curY = -20
	return m.Env.Reset()
}

func (m *mouseEnv) Step(t time.Duration, events ...interface{}) (reward float64,
	done bool, err error) {
	reward, done, err = m.Env.Step(t, events...)
	for _, evt := range events {
		if mouse, ok := evt.(*chrome.MouseEvent); ok {
			m.curX = mouse.X
			m.curY = mouse.Y
		}
	}
	return
}

func (m *mouseEnv) Observe() (obs Obs, err error) {
	obs, err = m.Env.Observe()
	if err != nil {
		return
	}
	img, err := obs.Image()
	if err != nil {
		return
	}

	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			widthForY := (y - m.curY) / 2
			switch true {
			case y >= m.curY+mouseSize && y < m.curY+mouseSize+mouseStem &&
				x == m.curX:
				fallthrough
			case y > m.curY && y < m.curY+mouseSize && x < m.curX+widthForY &&
				x > m.curX-widthForY:
				newImg.Set(x, y, color.Gray{Y: 0})
			case (x == m.curX+widthForY || x == m.curX-widthForY) &&
				y > m.curY && y < m.curY+mouseSize:
				fallthrough
			case y == m.curY && x == m.curX:
				fallthrough
			case y >= m.curY+mouseSize && y < m.curY+mouseSize+mouseStem &&
				(x == m.curX-1 || x == m.curX+1):
				fallthrough
			case y == m.curY+mouseSize+mouseStem &&
				(x >= m.curX-1 && x <= m.curX+1):
				fallthrough
			case y == m.curY+mouseSize && x < m.curX+widthForY &&
				x > m.curX-widthForY:
				newImg.Set(x, y, color.Gray{Y: 0xff})
			default:
				newImg.Set(x, y, img.At(x, y))
			}
		}
	}

	return &imageObs{Img: newImg}, nil
}
