package chrome

import "github.com/unixpickle/essentials"

// EvalPromise evaluates a piece of JavaScript code which
// returns a Promise.
//
// If retObj is non-nil, it is overwritten with the
// unmarshalled result of the Promise.
// Unmarshalling works like it does in encoding/json.
func (c *Conn) EvalPromise(code string, retObj interface{}) error {
	var rawResult struct {
		Result struct {
			Type  string      `json:"type"`
			Value interface{} `json:"value"`
		} `json:"result"`
	}

	rawResult.Result.Value = retObj

	resErr := c.call("Runtime.evaluate", map[string]interface{}{
		"expression":    code,
		"returnByValue": true,
		"awaitPromise":  true,
	}, &rawResult)
	if resErr != nil {
		return essentials.AddCtx("eval promise", resErr)
	}
	return nil
}
