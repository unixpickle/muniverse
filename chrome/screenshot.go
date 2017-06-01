package chrome

import (
	"bytes"
	"context"
	"encoding/base64"
	"io/ioutil"

	"github.com/unixpickle/essentials"
)

// ScreenshotPNG captures a screenshot of the page and
// returns the raw PNG image data.
func (c *Conn) ScreenshotPNG(ctx context.Context) (imageData []byte, err error) {
	defer essentials.AddCtxTo("capture PNG screenshot", &err)
	return c.captureScreenshot(ctx, nil)
}

// ScreenshotJPEG captures a compressed screenshot of the
// page and returns the raw JPEG image data.
//
// The quality is an integer in the range [0, 100].
func (c *Conn) ScreenshotJPEG(ctx context.Context, quality int) (imageData []byte,
	err error) {
	defer essentials.AddCtxTo("capture JPEG screenshot", &err)
	return c.captureScreenshot(ctx, map[string]interface{}{
		"format":  "jpeg",
		"quality": quality,
	})
}

func (c *Conn) captureScreenshot(ctx context.Context, params interface{}) ([]byte,
	error) {
	var resObj struct {
		Data string `json:"data"`
	}
	if err := c.call(ctx, "Page.captureScreenshot", params, &resObj); err != nil {
		return nil, err
	}
	dataReader := bytes.NewReader([]byte(resObj.Data))
	decoder := base64.NewDecoder(base64.StdEncoding, dataReader)
	return ioutil.ReadAll(decoder)
}
