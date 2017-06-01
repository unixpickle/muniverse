package chrome

import (
	"context"
	"strings"
	"time"

	"github.com/unixpickle/essentials"
)

// NavigateSync navigates to the given page and waits for
// it to finish loading.
func (c *Conn) NavigateSync(ctx context.Context, urlStr string) (err error) {
	defer essentials.AddCtxTo("DevTools synchronous page load", &err)

	// Prevent race condition where load happens before
	// we get to call Page.navigate.
	if err := c.call(ctx, "Page.stopLoading", nil, nil); err != nil {
		return err
	}

	pollChan := c.poll(ctx, "Page.loadEventFired", nil)

	if err := c.call(ctx, "Page.enable", nil, nil); err != nil {
		return err
	}
	defer c.call(ctx, "Page.disable", nil, nil)

	params := map[string]string{"url": urlStr}
	if err := c.call(ctx, "Page.navigate", params, nil); err != nil {
		return err
	}

	select {
	case err := <-pollChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
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
func (c *Conn) NavigateSafe(ctx context.Context, urlStr string) (err error) {
	essentials.AddCtxTo("navigate safe", &err)

RetryLoop:
	for {
		c.clearLog()
		if err = c.NavigateSync(ctx, urlStr); err != nil {
			return
		}

		// Give the log some time to come in.
		// TODO: see if this is necessary.
		select {
		case <-time.After(time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}

		for _, item := range c.ConsoleLog() {
			if strings.Contains(item, "net::ERR_NETWORK_CHANGED") {
				continue RetryLoop
			}
		}

		return
	}

}
