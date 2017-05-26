package chrome

import (
	"errors"
	"time"

	"github.com/unixpickle/essentials"
)

// NavigateSync navigates to the given page and waits for
// it to finish loading.
// It fails with an error if the timeout channel is either
// closed or sent a message before the load completes.
func (c *Conn) NavigateSync(urlStr string, timeout <-chan time.Time) (err error) {
	defer essentials.AddCtxTo("DevTools synchronous page load", &err)

	// Prevent race condition where load happens before
	// we get to call Page.navigate.
	if err := c.call("Page.stopLoading", nil, nil); err != nil {
		return err
	}

	pollChan := c.poll("Page.loadEventFired", nil)

	if err := c.call("Page.enable", nil, nil); err != nil {
		return err
	}
	defer c.call("Page.disable", nil, nil)

	params := map[string]string{"url": urlStr}
	if err := c.call("Page.navigate", params, nil); err != nil {
		return err
	}

	select {
	case err := <-pollChan:
		return err
	case <-timeout:
		return errors.New("timeout exceeded")
	}
}
