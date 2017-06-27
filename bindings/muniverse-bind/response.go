package main

import (
	"encoding/binary"
	"io"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/muniverse"
	"github.com/unixpickle/muniverse/chrome"

	"gopkg.in/mgo.v2/bson"
)

// Response is a message sent from the API back-end to an
// API front-end in response to a Call.
type Response struct {
	ID string `bson:"ID"`

	Error *string `bson:"Error,omitempty"`

	Spec        *muniverse.EnvSpec `bson:"Spec,omitempty"`
	UID         *string            `bson:"UID,omitempty"`
	StepResult  *StepResult        `bson:"StepResult,omitempty"`
	Observation *Observation       `bson:"Observation,omitempty"`
	KeyEvent    *chrome.KeyEvent   `bson:"KeyEvent,omitempty"`
}

// ErrorResponse creates a *Response with an error.
func ErrorResponse(e error) *Response {
	errMsg := e.Error()
	return &Response{Error: &errMsg}
}

// WriteResponse encodes a response to an output stream.
func WriteResponse(w io.Writer, r *Response) error {
	return essentials.AddCtx("write response", writeObject(w, r))
}

func writeObject(w io.Writer, obj interface{}) error {
	data, err := bson.Marshal(obj)
	if err != nil {
		return err
	}
	size := uint32(len(data))
	if err := binary.Write(w, binary.LittleEndian, size); err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

type StepResult struct {
	Reward float64 `bson:"Reward"`
	Done   bool    `bson:"Done"`
}

type Observation struct {
	Width  int    `bson:"Width"`
	Height int    `bson:"Height"`
	RGB    []byte `bson:"RGB"`
}
