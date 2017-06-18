package muniverse

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/muniverse/chrome"
)

const recordingPerm = 0755

// Recording is an on-disk record of observations and
// actions in an environment rollout.
//
// It is not safe to run methods on a Recording from more
// than one Goroutine at a time.
// Further, it is not safe to open a Recording if another
// Goroutine or program is writing to it.
type Recording struct {
	dir  string
	info recordingInfo
}

// OpenRecording opens an existing Recording from a
// directory at the given path.
func OpenRecording(path string) (rec *Recording, err error) {
	defer essentials.AddCtxTo("open recording", &err)
	rec = &Recording{dir: path}
	infoData, err := ioutil.ReadFile(rec.infoPath())
	if err != nil {
		return
	}
	err = json.Unmarshal(infoData, &rec.info)
	return
}

// CreateRecording creates a new Recording at the given
// path.
func CreateRecording(path string) (rec *Recording, err error) {
	defer essentials.AddCtxTo("create recording", &err)
	if err = os.Mkdir(path, recordingPerm); err != nil {
		return
	}
	defer func() {
		if err != nil {
			os.RemoveAll(path)
		}
	}()
	rec = &Recording{dir: path}
	if err = os.Mkdir(rec.framesPath(), recordingPerm); err != nil {
		return
	}
	data, err := json.Marshal(rec.info)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(rec.infoPath(), data, recordingPerm)
	return
}

// WriteObs adds an observation to the Recording.
func (r *Recording) WriteObs(obs Obs) (err error) {
	defer essentials.AddCtxTo("write observation", &err)

	data, err := ObsPNG(obs)
	if err != nil {
		return
	}
	name := strconv.Itoa(r.info.NumObs) + ".png"
	imgPath := filepath.Join(r.framesPath(), name)
	err = ioutil.WriteFile(imgPath, data, recordingPerm)
	if err != nil {
		return
	}

	r.info.NumObs++
	if err = r.writeInfo(); err != nil {
		r.info.NumObs--
		os.Remove(imgPath)
	}
	return
}

// WriteStep adds a step to the Recording.
func (r *Recording) WriteStep(info *StepInfo) (err error) {
	step := recordingStep{
		Time:   info.Time,
		Reward: info.Reward,
		Done:   info.Done,
	}
	for _, event := range info.Events {
		switch event := event.(type) {
		case *chrome.KeyEvent:
			step.Events = append(step.Events, recordingEvent{KeyEvent: event})
		case *chrome.MouseEvent:
			step.Events = append(step.Events, recordingEvent{MouseEvent: event})
		default:
			return fmt.Errorf("unsupported event type: %T", event)
		}
	}
	oldSteps := r.info.Steps
	r.info.Steps = append(r.info.Steps, step)
	err = r.writeInfo()
	if err != nil {
		r.info.Steps = oldSteps
	}
	return
}

// NumObs returns the number of recorded observations.
func (r *Recording) NumObs() int {
	return r.info.NumObs
}

// NumSteps returns the number of recorded steps.
func (r *Recording) NumSteps() int {
	return len(r.info.Steps)
}

// ReadObs reads the observation at the given index.
// It fails if the index is out of range.
func (r *Recording) ReadObs(idx int) (obs Obs, err error) {
	defer essentials.AddCtxTo("read observation", &err)
	if idx < 0 || idx >= r.NumObs() {
		return nil, errors.New("index out of range")
	}

	name := strconv.Itoa(idx) + ".png"
	imgPath := filepath.Join(r.framesPath(), name)
	data, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return
	}
	return pngObs(data), nil
}

// ReadStep reads the step at the given index.
// It fails if the index is out of range.
func (r *Recording) ReadStep(idx int) (info *StepInfo, err error) {
	defer essentials.AddCtxTo("read step", &err)
	if idx < 0 || idx >= r.NumSteps() {
		err = errors.New("index out of range")
		return
	}
	step := r.info.Steps[idx]
	info = &StepInfo{
		Time:   step.Time,
		Reward: step.Reward,
		Done:   step.Done,
	}
	for _, evt := range step.Events {
		if evt.KeyEvent != nil {
			info.Events = append(info.Events, evt.KeyEvent)
		} else if evt.MouseEvent != nil {
			info.Events = append(info.Events, evt.MouseEvent)
		}
	}
	return
}

func (r *Recording) writeInfo() error {
	data, err := json.Marshal(r.info)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.infoPath(), data, recordingPerm)
}

func (r *Recording) infoPath() string {
	return filepath.Join(r.dir, "info.json")
}

func (r *Recording) framesPath() string {
	return filepath.Join(r.dir, "frames")
}

type recordEnv struct {
	Env
	destDir string
	curRec  *Recording
	gen     *rand.Rand
}

// RecordEnv creates an Env wrapper which stores a new
// Recording for every episode.
//
// Recordings are stored inside the given directory.
// Each recording is assigned a pseudo-random directory
// name to prevent collisions.
// However, the names all begin with "recording_".
//
// Closing the resulting Env will automatically close e.
func RecordEnv(e Env, dir string) Env {
	return &recordEnv{
		Env:     e,
		destDir: dir,
		gen:     rand.New(rand.NewSource(rand.Int63() ^ time.Now().UnixNano())),
	}
}

func (r *recordEnv) Reset() (err error) {
	defer essentials.AddCtxTo("reset recorded environment", &err)
	if err = r.createDir(); err != nil {
		return
	}
	name := fmt.Sprintf("recording_%d_%d", time.Now().UnixNano(), r.gen.Int63())
	rec, err := CreateRecording(filepath.Join(r.destDir, name))
	if err != nil {
		return
	}
	r.curRec = rec
	return r.Env.Reset()
}

func (r *recordEnv) Observe() (obs Obs, err error) {
	defer essentials.AddCtxTo("observe recorded environment", &err)
	obs, err = r.Env.Observe()
	if err != nil {
		return
	}
	err = r.curRec.WriteObs(obs)
	return
}

func (r *recordEnv) Step(t time.Duration, events ...interface{}) (reward float64,
	done bool, err error) {
	defer essentials.AddCtxTo("step recorded environment", &err)
	reward, done, err = r.Env.Step(t, events...)
	if err != nil {
		return
	}
	err = r.curRec.WriteStep(&StepInfo{
		Time:   t,
		Events: events,
		Reward: reward,
		Done:   done,
	})
	return
}

func (r *recordEnv) createDir() error {
	info, err := os.Stat(r.destDir)
	if os.IsNotExist(err) {
		return os.Mkdir(r.destDir, recordingPerm)
	} else if !info.IsDir() {
		return errors.New("record to " + r.destDir + ": not a directory")
	}
	return nil
}

// StepInfo stores information about a step in an
// environment and the result of that step.
type StepInfo struct {
	Time   time.Duration
	Events []interface{}

	Reward float64
	Done   bool
}

type recordingInfo struct {
	NumObs int
	Steps  []recordingStep
}

type recordingStep struct {
	Time   time.Duration
	Reward float64
	Done   bool
	Events []recordingEvent
}

type recordingEvent struct {
	KeyEvent   *chrome.KeyEvent
	MouseEvent *chrome.MouseEvent
}
