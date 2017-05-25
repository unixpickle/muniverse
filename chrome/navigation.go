package chrome

import (
	"errors"
	"time"

	"github.com/unixpickle/essentials"
)

// RefreshSync refreshes the current page and waits for it
// to finish loading.
// It fails with an error if the timeout channel is either
// closed or sent a message before the load completes.
func (c *Conn) RefreshSync(timeout <-chan time.Time) (err error) {
	essentials.AddCtxTo("DevTools synchronous page refresh", &err)

	// Prevent race condition where load happens before
	// we get to call Page.reload.
	if err := c.call("Page.stopLoading", nil, nil); err != nil {
		return err
	}

	pollChan := c.poll("Page.loadEventFired", nil)

	if err := c.call("Page.enable", nil, nil); err != nil {
		return err
	}
	defer c.call("Page.disable", nil, nil)

	if err := c.call("Page.reload", nil, nil); err != nil {
		return err
	}

	select {
	case err := <-pollChan:
		return err
	case <-timeout:
		return errors.New("timeout exceeded")
	}
}
