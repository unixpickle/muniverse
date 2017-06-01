package chrome

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/unixpickle/essentials"
)

// EvalPromise evaluates a piece of JavaScript code which
// returns a Promise.
//
// If retObj is non-nil, it is overwritten with the
// unmarshalled result of the Promise.
// Unmarshalling works like it does in encoding/json.
func (c *Conn) EvalPromise(ctx context.Context, code string,
	retObj interface{}) (err error) {
	defer essentials.AddCtxTo("eval promise", &err)

	var rawResult struct {
		ExceptionDetails interface{} `json:"exceptionDetails"`

		Result struct {
			Type  string      `json:"type"`
			Value interface{} `json:"value"`
		} `json:"result"`
	}

	rawResult.Result.Value = retObj

	resErr := c.call(ctx, "Runtime.evaluate", map[string]interface{}{
		"expression":    code,
		"returnByValue": true,
		"awaitPromise":  true,
	}, &rawResult)

	if resErr != nil {
		return resErr
	} else if rawResult.ExceptionDetails != nil {
		details, _ := json.Marshal(rawResult.ExceptionDetails)
		return errors.New("uncaught exception: " + string(details))
	}

	return nil
}
