package main

import (
	"encoding/binary"
	"io"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/muniverse"

	"gopkg.in/mgo.v2/bson"
)

// Response is a message sent from the API back-end to an
// API front-end in response to a Call.
type Response struct {
	ID string

	Error *string

	Spec        *muniverse.EnvSpec
	UID         *string
	StepResult  *StepResult
	Observation *Observation
}

// ErrorResponse creates a *Response with an error.
func ErrorResponse(e error) *Response {
	errMsg := e.Error()
	return &Response{Error: &errMsg}
}

// WriteResponse encodes a response to an output stream.
func WriteResponse(w io.Writer, r *Response) (err error) {
	defer essentials.AddCtxTo("write response", &err)
	data, err := bson.Marshal(r)
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
	Reward float64
	Done   bool
}

type Observation struct {
	Width  int
	Height int
	RGB    []byte
}
