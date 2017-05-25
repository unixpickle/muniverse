package chrome

import (
	"reflect"
	"testing"
	"time"
)

func TestConnCall(t *testing.T) {
	conn := testingConn(t, false)
	defer conn.Close()

	var actual, expected struct {
		Result struct {
			Type  string `json:"type"`
			Value struct {
				String string `json:"string"`
				Number int    `json:"number"`
			} `json:"value"`
		} `json:"result"`
	}

	resErr := conn.call("Runtime.evaluate", map[string]interface{}{
		"expression":    "(function(){return {number: (3+2)*3, string: 'hi'};})()",
		"returnByValue": true,
	}, &actual)
	if resErr != nil {
		t.Fatal(resErr)
	}

	expected.Result.Type = "object"
	expected.Result.Value.Number = 15
	expected.Result.Value.String = "hi"

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestConnPoll(t *testing.T) {
	conn := testingConn(t, true)
	defer conn.Close()

	// Poll for event **before** triggering it.
	var result struct {
		Timestamp float64 `json:"timestamp"`
	}
	errChan := conn.poll("Page.loadEventFired", &result)

	// Trigger reload event.
	if err := conn.call("Page.enable", nil, nil); err != nil {
		t.Fatal(err)
	}
	defer conn.call("Page.disable", nil, nil)
	if err := conn.call("Page.reload", nil, nil); err != nil {
		t.Fatal(err)
	}

	select {
	case err := <-errChan:
		if err != nil {
			t.Error(err)
		} else if result.Timestamp == 0 {
			t.Error("got zero timestamp")
		}
	case <-time.After(time.Second * 20):
		t.Error("load timeout")
	}
}
