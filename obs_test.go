package muniverse

import (
	"bytes"
	"image"
	"image/png"
	"math/rand"
	"reflect"
	"testing"
)

func TestObsRGB(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 200))
	expected := make([]uint8, 0, 100*200*3)
	for i := 0; i < 100*200; i++ {
		idx := i * 4
		for j := 0; j < 3; j++ {
			val := uint8(rand.Intn(0x100))
			img.Pix[idx+j] = val
			expected = append(expected, val)
		}
		img.Pix[idx+3] = 0xff
	}

	var pngData bytes.Buffer
	if err := png.Encode(&pngData, img); err != nil {
		t.Fatal(err)
	}

	actual, width, height, err := RGB(pngObs(pngData.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	if width != 100 || height != 200 {
		t.Errorf("expected size 100x200 but got %dx%d", width, height)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("got different pixel buffer")
	}
}

func BenchmarkObsRGB(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 200))
	expected := make([]uint8, 0, 100*200*3)
	for i := 0; i < 100*200; i++ {
		idx := i * 4
		for j := 0; j < 3; j++ {
			val := uint8(rand.Intn(0x100))
			img.Pix[idx+j] = val
			expected = append(expected, val)
		}
		img.Pix[idx+3] = 0xff
	}
	var pngData bytes.Buffer
	if err := png.Encode(&pngData, img); err != nil {
		b.Fatal(err)
	}
	obs := pngObs(pngData.Bytes())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RGB(obs)
	}
}
