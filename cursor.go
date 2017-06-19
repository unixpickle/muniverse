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

type cursorEnv struct {
	Env

	initX int
	initY int

	curX int
	curY int
}

// CursorEnv creates a wrapped environment which renders a
// cursor at the current mouse location.
//
// At every episode, the mouse is initialized to the given
// x and y coordinates.
//
// By default, Chrome will not render a mouse.
// Thus, it is necessary to render a mouse manually.
//
// When the resulting Env is closed, e is closed as well.
func CursorEnv(e Env, initX, initY int) Env {
	return &cursorEnv{Env: e, initX: initX, initY: initY}
}

func (c *cursorEnv) Reset() error {
	c.curX = c.initX
	c.curY = c.initY
	return c.Env.Reset()
}

func (c *cursorEnv) Step(t time.Duration, events ...interface{}) (reward float64,
	done bool, err error) {
	reward, done, err = c.Env.Step(t, events...)
	for _, evt := range events {
		if mouse, ok := evt.(*chrome.MouseEvent); ok {
			c.curX = mouse.X
			c.curY = mouse.Y
		}
	}
	return
}

func (c *cursorEnv) Observe() (obs Obs, err error) {
	obs, err = c.Env.Observe()
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
			widthForY := (y - c.curY) / 2
			switch true {
			case y >= c.curY+mouseSize && y < c.curY+mouseSize+mouseStem &&
				x == c.curX:
				fallthrough
			case y > c.curY && y < c.curY+mouseSize && x < c.curX+widthForY &&
				x > c.curX-widthForY:
				newImg.Set(x, y, color.Gray{Y: 0})
			case (x == c.curX+widthForY || x == c.curX-widthForY) &&
				y > c.curY && y < c.curY+mouseSize:
				fallthrough
			case y == c.curY && x == c.curX:
				fallthrough
			case y >= c.curY+mouseSize && y < c.curY+mouseSize+mouseStem &&
				(x == c.curX-1 || x == c.curX+1):
				fallthrough
			case y == c.curY+mouseSize+mouseStem &&
				(x >= c.curX-1 && x <= c.curX+1):
				fallthrough
			case y == c.curY+mouseSize && x < c.curX+widthForY &&
				x > c.curX-widthForY:
				newImg.Set(x, y, color.Gray{Y: 0xff})
			default:
				newImg.Set(x, y, img.At(x, y))
			}
		}
	}

	return &imageObs{Img: newImg}, nil
}
