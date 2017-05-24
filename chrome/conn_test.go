package chrome

import (
	"reflect"
	"testing"
)

func TestConnCalls(t *testing.T) {
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
