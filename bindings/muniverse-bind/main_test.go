package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/unixpickle/muniverse"
	"github.com/unixpickle/muniverse/chrome"
)

func TestProtocol(t *testing.T) {
	backendIn, clientOut, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer backendIn.Close()
	defer clientOut.Close()
	clientIn, backendOut, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer clientIn.Close()
	defer backendOut.Close()

	errChan := make(chan error, 1)
	go func() {
		server := &Server{
			Input:  backendIn,
			Output: backendOut,
		}
		errChan <- server.Run()
	}()

	safeCall := func(c *Call) *Response {
		return ensureNoError(t, makeTestCall(t, clientIn, clientOut, c))
	}

	// Test SpecForName.
	for _, name := range []string{"KatanaFruits-v0", "", "KatanaFruits-v0"} {
		resp := safeCall(&Call{
			SpecForName: &CallSpecForName{Name: name},
		})
		spec := resp.Spec
		expected := muniverse.SpecForName(name)
		if !reflect.DeepEqual(spec, expected) {
			t.Errorf("bad spec for name %#v: got %#v but expected %#v",
				name, spec, expected)
		}
	}

	// Test KeyForCode.
	for _, code := range []string{"ArrowLeft", "", "Space"} {
		resp := safeCall(&Call{
			KeyForCode: &CallKeyForCode{Code: code},
		})
		actual := resp.KeyEvent
		evt, ok := chrome.KeyEvents[code]
		var expected *chrome.KeyEvent
		if ok {
			expected = &evt
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("bad key for code %#v: got %#v but expected %#v",
				code, actual, expected)
		}
	}

	// Test environment lifecycle.
	for _, name := range []string{"KatanaFruits-v0", "MinimalDots-v0"} {
		spec := muniverse.SpecForName(name)
		if spec == nil {
			t.Fatal("no environment:", name)
		}
		resp := safeCall(&Call{
			NewEnv: &CallNewEnv{Spec: spec},
		})
		uid := *resp.UID

		safeCall(&Call{
			Reset: &CallReset{UID: uid},
		})

		// Make sure observations are the right size.
		resp = safeCall(&Call{
			Observe: &CallObserve{UID: uid},
		})
		if resp.Observation.Width != spec.Width ||
			resp.Observation.Height != spec.Height {
			t.Errorf("bad observation size: %dx%d (expected %dx%d)",
				resp.Observation.Width, resp.Observation.Height,
				spec.Width, spec.Height)
		}
		if len(resp.Observation.RGB) != 3*spec.Width*spec.Height {
			t.Errorf("bad RGB length: %d (expected %d)",
				len(resp.Observation.RGB), 3*spec.Width*spec.Height)
		}

		// Make sure taking steps does not fail.
		keyEvt := chrome.KeyEvents["ArrowLeft"]
		keyEvt.Type = chrome.KeyDown
		mouseEvt := &chrome.MouseEvent{
			Type: chrome.MouseMoved,
			X:    100,
			Y:    100,
		}
		resp = safeCall(&Call{
			Step: &CallStep{
				UID:     uid,
				Seconds: 0.1,
				Events: []*Event{
					&Event{KeyEvent: &keyEvt},
					&Event{MouseEvent: mouseEvt},
				},
			},
		})
		if resp.StepResult.Done {
			t.Error("environment finished before expected")
		}
		if resp.StepResult.Reward != 0 {
			t.Errorf("environment gave reward before expected")
		}

		safeCall(&Call{
			CloseEnv: &CallCloseEnv{UID: uid},
		})

		resp = makeTestCall(t, clientIn, clientOut, &Call{
			Reset: &CallReset{UID: uid},
		})
		if resp.Error == nil {
			t.Error("expected error when using closed environment")
		}
	}

	clientOut.Close()
	select {
	case <-errChan:
	case <-time.After(time.Second * 10):
		t.Error("failed to shutdown after close")
	}
}

func makeTestCall(t *testing.T, in io.Reader, out io.Writer, call *Call) *Response {
	call.ID = fmt.Sprintf("%d", rand.Int63())
	if err := writeObject(out, call); err != nil {
		t.Fatal(err)
	}
	var resp *Response
	if err := readObject(in, &resp); err != nil {
		t.Fatal(err)
	}
	if resp.ID != call.ID {
		t.Fatalf("mismatching ID: got %#v but sent %#v", resp.ID, call.ID)
	}
	return resp
}

func ensureNoError(t *testing.T, r *Response) *Response {
	if r.Error != nil {
		t.Fatal("unexpected error from backend:", *r.Error)
	}
	return r
}
