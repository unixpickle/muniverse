package main

import (
	"errors"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/muniverse"
	"github.com/unixpickle/muniverse/chrome"
)

func main() {
	essentials.Die((&Server{
		Input:  os.Stdin,
		Output: os.Stdout,
	}).Run())
}

type Server struct {
	Input  io.Reader
	Output io.Writer

	envsLock   sync.RWMutex
	currentUID int64
	envsByUID  map[string]muniverse.Env
}

// Run runs the server until an unrecoverable error is
// encountered.
// Typically, Run will not return until the input stream
// has been closed or produces an error.
func (b *Server) Run() error {
	b.envsByUID = map[string]muniverse.Env{}

	done := make(chan struct{})
	errChan := make(chan error, 1)
	gotError := func(e error) {
		select {
		case errChan <- e:
			close(done)
		default:
		}
	}

	var wg sync.WaitGroup

	// Synchronize all responses to prevent overlapping
	// writes to one io.Writer.
	respChan := make(chan *Response, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case r := <-respChan:
				if err := WriteResponse(b.Output, r); err != nil {
					gotError(err)
					return
				}
			case <-done:
				return
			}
		}
	}()

ReadLoop:
	for {
		call, err := ReadCall(b.Input)
		select {
		case <-done:
			break ReadLoop
		default:
		}
		if err != nil {
			gotError(err)
			break
		}
		go func() {
			response := b.handleCall(call)
			response.ID = call.ID
			select {
			case respChan <- response:
			case <-done:
			}
		}()
	}

	wg.Wait()

	return <-errChan
}

func (b *Server) handleCall(call *Call) *Response {
	switch true {
	case call.SpecForName != nil:
		return b.specForName(call)
	case call.NewEnv != nil, call.NewEnvContainer != nil, call.NewEnvChrome != nil:
		return b.newEnv(call)
	case call.CloseEnv != nil:
		return b.closeEnv(call)
	case call.Reset != nil:
		return b.reset(call)
	case call.Step != nil:
		return b.step(call)
	case call.Observe != nil:
		return b.observe(call)
	case call.KeyForCode != nil:
		return b.keyForCode(call)
	default:
		return ErrorResponse(errors.New("malformed call"))
	}
}

func (b *Server) specForName(call *Call) *Response {
	return &Response{
		Spec: muniverse.SpecForName(call.SpecForName.Name),
	}
}

func (b *Server) newEnv(call *Call) *Response {
	var env muniverse.Env
	var err error
	switch true {
	case call.NewEnv != nil:
		if call.NewEnv.Spec == nil {
			err = errors.New("null specification")
		} else {
			env, err = muniverse.NewEnv(call.NewEnv.Spec)
		}
	case call.NewEnvContainer != nil:
		if call.NewEnvContainer.Spec == nil {
			err = errors.New("null specification")
		} else {
			env, err = muniverse.NewEnvContainer(call.NewEnvContainer.Container,
				call.NewEnvContainer.Spec)
		}
	case call.NewEnvChrome != nil:
		if call.NewEnvChrome.Spec == nil {
			err = errors.New("null specification")
		} else {
			env, err = muniverse.NewEnvChrome(call.NewEnvChrome.Host,
				call.NewEnvChrome.GameHost, call.NewEnvChrome.Spec)
		}
	default:
		panic("unreachable")
	}
	response := &Response{}
	if err != nil {
		message := err.Error()
		response.Error = &message
	} else {
		b.envsLock.Lock()
		uid := strconv.FormatInt(b.currentUID, 10)
		b.currentUID++
		b.envsByUID[uid] = env
		b.envsLock.Unlock()
		response.UID = &uid
	}
	return response
}

func (b *Server) closeEnv(call *Call) *Response {
	b.envsLock.Lock()
	env, ok := b.envsByUID[call.CloseEnv.UID]
	delete(b.envsByUID, call.CloseEnv.UID)
	b.envsLock.Unlock()

	if !ok {
		return ErrorResponse(errors.New("environment does not exist"))
	} else if err := env.Close(); err != nil {
		return ErrorResponse(err)
	}
	return &Response{}
}

func (b *Server) reset(call *Call) *Response {
	env, errResp := b.lookupEnv(call.Reset.UID)
	if errResp != nil {
		return errResp
	}
	if err := env.Reset(); err != nil {
		return ErrorResponse(err)
	}
	return &Response{}
}

func (b *Server) step(call *Call) *Response {
	env, errResp := b.lookupEnv(call.Step.UID)
	if errResp != nil {
		return errResp
	}
	t := time.Duration(float64(time.Second) * call.Step.Seconds)
	var events []interface{}
	for _, evt := range call.Step.Events {
		if evt.KeyEvent != nil {
			events = append(events, evt.KeyEvent)
		} else if evt.MouseEvent != nil {
			events = append(events, evt.MouseEvent)
		}
	}
	reward, done, err := env.Step(t, events...)
	if err != nil {
		return ErrorResponse(err)
	}
	return &Response{
		StepResult: &StepResult{
			Reward: reward,
			Done:   done,
		},
	}
}

func (b *Server) observe(call *Call) *Response {
	env, errResp := b.lookupEnv(call.Observe.UID)
	if errResp != nil {
		return errResp
	}
	obs, err := env.Observe()
	if err != nil {
		return ErrorResponse(err)
	}
	data, width, height, err := muniverse.RGB(obs)
	if err != nil {
		return ErrorResponse(err)
	}
	return &Response{
		Observation: &Observation{
			Width:  width,
			Height: height,
			RGB:    data,
		},
	}
}

func (b *Server) keyForCode(call *Call) *Response {
	evt, ok := chrome.KeyEvents[call.KeyForCode.Code]
	if ok {
		return &Response{
			KeyEvent: &evt,
		}
	} else {
		return &Response{}
	}
}

func (b *Server) lookupEnv(uid string) (env muniverse.Env, errResp *Response) {
	b.envsLock.RLock()
	defer b.envsLock.RUnlock()
	if env, ok := b.envsByUID[uid]; ok {
		return env, nil
	} else {
		return nil, ErrorResponse(errors.New("environment does not exist"))
	}
}
