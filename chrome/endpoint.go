package chrome

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/unixpickle/essentials"
)

// Endpoint stores information about an active page,
// extension, iframe, or any other API endpoint.
type Endpoint struct {
	Description string `json:"description"`
	FrontendURL string `json:"devtoolsFrontendUrl"`
	FaviconURL  string `json:"faviconUrl"`
	ID          string `json:"id"`
	Title       string `json:"title"`

	// Options: page, service_worker, background_page,
	// iframe, and possibly others.
	Type string `json:"type"`

	URL          string `json:"url"`
	WebSocketURL string `json:"webSocketDebuggerUrl"`
}

// Endpoints lists the endpoints at the given API host.
//
// For example, the host might be "localhost:9222" and the
// result might include a list of open tabs.
func Endpoints(ctx context.Context, host string) (endpoints []*Endpoint,
	err error) {
	defer essentials.AddCtxTo("list DevTools endpoints", &err)
	u := url.URL{
		Host:   host,
		Scheme: "http",
		Path:   "/json",
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &endpoints)
	return
}
