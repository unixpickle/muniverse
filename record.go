package muniverse

import (
	"bytes"
	"errors"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/unixpickle/essentials"
)

type recordEnv struct {
	Env
	destDir   string
	resetTime string
	frameIdx  int
}

// Record wraps an Env and saves observations as image
// files.
//
// Image files are saved to the given directory.
// If the directory does not exist, it will be created.
//
// Filenames are formatted as:
//
//     env<reset_time>-<frame_idx>.png
//
// The resulting files are meant to be used with a tool
// like ffmpeg to construct a video.
// For example:
//
//     ffmpeg -i env1496014190_%04d.png video.webm
//
// In practice, you will likely want to specify the
// framerate, codec, and dimensions:
//
//     ffmpeg -r 10 \
//         -i env1496015380692491427-%04d.png \
//         -c:v libx264 \
//         -pix_fmt yuv420p \
//         -crf 23 \
//         -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" \
//         output.mp4
//
func Record(e Env, destDir string) (wrapped Env, err error) {
	defer essentials.AddCtxTo("record environment", &err)
	info, err := os.Stat(destDir)
	if os.IsNotExist(err) {
		if err = os.Mkdir(destDir, 0755); err != nil {
			return
		}
	} else if err != nil {
		return
	} else if !info.IsDir() {
		return nil, errors.New("not a directory: " + destDir)
	}
	return &recordEnv{
		Env:     e,
		destDir: destDir,
	}, nil
}

func (r *recordEnv) Reset() error {
	r.resetTime = strconv.FormatInt(time.Now().UnixNano(), 10)
	r.frameIdx = 0
	return r.Env.Reset()
}

func (r *recordEnv) Observe() (obs Obs, err error) {
	obs, err = r.Env.Observe()
	if err == nil {
		defer essentials.AddCtxTo("record frame", &err)
		var data []byte
		data, err = pngDataFromObs(obs)
		if err != nil {
			return
		}
		name := fmt.Sprintf("env%s-%04d.png", r.resetTime, r.frameIdx)
		path := filepath.Join(r.destDir, name)
		err = ioutil.WriteFile(path, data, 0644)
		r.frameIdx++
	}
	return
}

func pngDataFromObs(obs Obs) ([]byte, error) {
	if po, ok := obs.(pngObs); ok {
		return po, nil
	} else {
		img, err := obs.Image()
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
}
