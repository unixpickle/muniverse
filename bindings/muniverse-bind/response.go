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

	Spec *muniverse.EnvSpec
	UID  *string
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
