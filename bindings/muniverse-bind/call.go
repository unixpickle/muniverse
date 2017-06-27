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
	ID string `bson:"ID"`

	SpecForName     *CallSpecForName     `bson:"SpecForName"`
	NewEnv          *CallNewEnv          `bson:"NewEnv"`
	NewEnvContainer *CallNewEnvContainer `bson:"NewEnvContainer"`
	NewEnvChrome    *CallNewEnvChrome    `bson:"NewEnvChrome"`
	CloseEnv        *CallCloseEnv        `bson:"CloseEnv"`
	Reset           *CallReset           `bson:"Reset"`
	Step            *CallStep            `bson:"Step"`
	Observe         *CallObserve         `bson:"Observe"`
	KeyForCode      *CallKeyForCode      `bson:"KeyForCode"`
}

// ReadCall decodes a Call from an input stream.
func ReadCall(r io.Reader) (call *Call, err error) {
	defer essentials.AddCtxTo("read call", &err)
	err = readObject(r, &call)
	return
}

func readObject(r io.Reader, objOut interface{}) error {
	var len uint32
	if err := binary.Read(r, binary.LittleEndian, &len); err != nil {
		return err
	}
	data := make([]byte, int(len))
	if _, err := io.ReadFull(r, data); err != nil {
		return err
	}
	return bson.Unmarshal(data, objOut)
}

type CallSpecForName struct {
	Name string `bson:"Name"`
}

type CallNewEnv struct {
	Spec *muniverse.EnvSpec `bson:"Spec"`
}

type CallNewEnvContainer struct {
	Container string             `bson:"Container"`
	Spec      *muniverse.EnvSpec `bson:"Spec"`
}

type CallNewEnvChrome struct {
	Host     string             `bson:"Host"`
	GameHost string             `bson:"GameHost"`
	Spec     *muniverse.EnvSpec `bson:"Spec"`
}

type CallCloseEnv struct {
	UID string `bson:"UID"`
}

type CallReset struct {
	UID string `bson:"UID"`
}

type CallStep struct {
	UID     string   `bson:"UID"`
	Seconds float64  `bson:"Seconds"`
	Events  []*Event `bson:"Events"`
}

type CallObserve struct {
	UID string `bson:"UID"`
}

type CallKeyForCode struct {
	Code string `bson:"Code"`
}

type Event struct {
	KeyEvent   *chrome.KeyEvent   `bson:"KeyEvent"`
	MouseEvent *chrome.MouseEvent `bson:"MouseEvent"`
}
