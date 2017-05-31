package chrome

import (
	"errors"
	"strings"
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

// NavigateSafe is like NavigateSync, but it checks for
// network related errors in the console and retries if
// such errors are encountered.
//
// This is meant to deal with a bug in Chromium when
// running under Docker where a net::ERR_NETWORK_CHANGED
// error is triggered.
// It is yet to be clear whether the error originates due
// to Docker networking or due to Chromium.
func (c *Conn) NavigateSafe(urlStr string, timeout <-chan time.Time) (err error) {
	essentials.AddCtxTo("navigate safe", &err)

RetryLoop:
	for {
		c.clearLog()
		if err = c.NavigateSync(urlStr, timeout); err != nil {
			return
		}

		// Give the log some time to come in.
		// TODO: see if this is necessary.
		time.Sleep(time.Second)

		for _, item := range c.ConsoleLog() {
			if strings.Contains(item, "net::ERR_NETWORK_CHANGED") {
				continue RetryLoop
			}
		}

		return
	}

}
