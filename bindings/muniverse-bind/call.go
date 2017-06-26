package main

import (
	"encoding/binary"
	"io"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/muniverse"
	"github.com/unixpickle/muniverse/chrome"

	"gopkg.in/mgo.v2/bson"
)

// Call is a message sent from an API front-end to the API
// back-end (this program).
type Call struct {
	ID string

	SpecForName     *CallSpecForName
	NewEnv          *CallNewEnv
	NewEnvContainer *CallNewEnvContainer
	NewEnvChrome    *CallNewEnvChrome
	CloseEnv        *CallCloseEnv
	Reset           *CallReset
	Step            *CallStep
	Observe         *CallObserve
}

// ReadCall decodes a Call from an input stream.
func ReadCall(r io.Reader) (call *Call, err error) {
	defer essentials.AddCtxTo("read call", &err)
	var len uint32
	if err := binary.Read(r, binary.LittleEndian, &len); err != nil {
		return nil, err
	}
	data := make([]byte, int(len))
	if _, err := io.ReadFull(r, data); err != nil {
		return nil, err
	}
	err = bson.Unmarshal(data, &call)
	return
}

type CallSpecForName struct {
	Name string
}

type CallNewEnv struct {
	Spec *muniverse.EnvSpec
}

type CallNewEnvContainer struct {
	Container string
	Spec      *muniverse.EnvSpec
}

type CallNewEnvChrome struct {
	Host     string
	GameHost string
	Spec     *muniverse.EnvSpec
}

type CallCloseEnv struct {
	UID string
}

type CallReset struct {
	UID string
}

type CallStep struct {
	UID     string
	Seconds float64
	Events  []*Event
}

type CallObserve struct {
	UID string
}

type Event struct {
	KeyEvent   *chrome.KeyEvent
	MouseEvent *chrome.MouseEvent
}
