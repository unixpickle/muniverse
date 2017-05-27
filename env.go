package muniverse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/muniverse/chrome"
)

const (
	portRange        = "9000-9999"
	defaultContainer = "unixpickle/muniverse:0.1.3"
)

const refreshTimeout = time.Minute

// SpecForName finds the first entry in EnvSpecs with the
// given name.
// If no entry is found, nil is returned.
func SpecForName(name string) *EnvSpec {
	for _, s := range EnvSpecs {
		if s.Name == name {
			return s
		}
	}
	return nil
}

// An Env controls a running environment.
type Env struct {
	spec     EnvSpec
	gameHost string

	containerID string
	devConn     *chrome.Conn
	lastScore   float64

	// Used to garbage collect the container if we
	// exit ungracefully.
	killSocket net.Conn
}

// NewEnv creates a new environment inside the default
// Docker container.
// This may take a few minutes to run the first time,
// since it has to download a large container.
func NewEnv(spec *EnvSpec) (*Env, error) {
	return NewEnvContainer(defaultContainer, spec)
}

// NewEnvContainer creates a new environment inside a
// Docker container of the given name.
func NewEnvContainer(container string, spec *EnvSpec) (env *Env, err error) {
	defer essentials.AddCtxTo("create environment", &err)

	id, err := dockerRun(container, spec)
	if err != nil {
		return
	}

	ports, err := dockerBoundPorts(id)
	if err != nil {
		return
	}

	killSock, err := net.Dial("tcp", "localhost:"+ports["1337/tcp"])
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			killSock.Close()
		}
	}()

	conn, err := connectDevTools("localhost:" + ports["9222/tcp"])
	if err != nil {
		return
	}

	return &Env{
		spec:        *spec,
		gameHost:    "localhost",
		containerID: id,
		devConn:     conn,
		killSocket:  killSock,
	}, nil
}

// NewEnvChrome connects to an existing Chrome DevTools
// server and runs an environment in there.
//
// The gameHost argument specifies where to load games.
// For example, gameHost might be "localhost:8080" if the
// game "Foobar" should be loaded from
// "http://localhost/Foobar".
//
// The Chrome instance must have at least one page open,
// since an open page is selected and used to run the
// environment.
func NewEnvChrome(host, gameHost string, spec *EnvSpec) (env *Env, err error) {
	defer essentials.AddCtxTo("create environment", &err)

	conn, err := connectDevTools(host)
	if err != nil {
		return
	}

	return &Env{
		spec:     *spec,
		gameHost: gameHost,
		devConn:  conn,
	}, nil
}

// Spec returns a copy of the specification used to create
// this environment.
func (e *Env) Spec() *EnvSpec {
	res := e.spec
	return &res
}

// Reset resets the environment to a starting state.
//
// This should be called before the first step.
// It should also be called every time an episode finishes
// before steps are taken for the next episode.
func (e *Env) Reset() (err error) {
	defer essentials.AddCtxTo("reset environment", &err)

	err = e.devConn.NavigateSafe(e.envURL(), time.After(refreshTimeout))
	if err != nil {
		return
	}

	err = e.devConn.EvalPromise("window.muniverse.init();", nil)
	if err != nil {
		return
	}
	err = e.devConn.EvalPromise("window.muniverse.score();", &e.lastScore)

	return
}

// Step takes a step in the environment.
// In particular, it sends the given events and advances
// the game by the given number of milliseconds.
//
// If done is true, then the episode has ended and no more
// steps should be taken before Reset is called.
// However, observations may be made even after the
// episode has ended.
//
// The only supported event type is *chrome.MouseEvent.
func (e *Env) Step(millis int, events ...interface{}) (reward float64,
	done bool, err error) {
	defer essentials.AddCtxTo("step environment", &err)

	for _, event := range events {
		switch event := event.(type) {
		case *chrome.MouseEvent:
			err = e.devConn.DispatchMouseEvent(event)
			if err != nil {
				return
			}
		default:
			err = fmt.Errorf("unsupported event type: %T", event)
			return
		}
	}

	timeStr := strconv.Itoa(millis)
	err = e.devConn.EvalPromise("window.muniverse.step("+timeStr+");", &done)
	if err != nil {
		return
	}

	lastScore := e.lastScore
	err = e.devConn.EvalPromise("window.muniverse.score();", &e.lastScore)
	if err != nil {
		return
	}
	reward = e.lastScore - lastScore

	return
}

// Observe produces an observation for the current state.
func (e *Env) Observe() (obs Obs, err error) {
	pngData, err := e.devConn.ScreenshotPNG()
	if err != nil {
		return
	}
	return pngObs(pngData), nil
}

// Close cleans up the resources associated with the
// environment.
// You should only call Close() once you are done using
// the environment.
func (e *Env) Close() (err error) {
	defer essentials.AddCtxTo("close environment", &err)

	errs := []error{
		e.devConn.Close(),
	}
	if e.containerID != "" {
		errs = append(errs, exec.Command("docker", "kill", e.containerID).Run())
	}

	if e.killSocket != nil {
		// TODO: look into if this can ever produce an error,
		// since the container might already have closed the
		// socket by now.
		//
		// We don't close this *before* stopping the container
		// since `docker kill` might fail if the container
		// already died and was cleaned up.
		e.killSocket.Close()
	}

	// Any calls after Close() should trigger simple errors.
	e.devConn = nil
	e.killSocket = nil

	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Env) envURL() string {
	return "http://" + e.gameHost + "/" + e.spec.Name
}

func dockerRun(container string, spec *EnvSpec) (id string, err error) {
	args := []string{
		"run",
		"-p",
		portRange + ":9222",
		"-p",
		portRange + ":1337",
		"-d",   // Run in detached mode.
		"--rm", // Automatically delete the container.
		"-i",   // Give netcat a stdin to read from.
		container,
		fmt.Sprintf("--window-size=%d,%d", spec.Width, spec.Height),
	}

	output, err := exec.Command("docker", args...).Output()
	if err != nil {
		return "", essentials.AddCtx("docker run",
			fmt.Errorf("%s (make sure docker is up-to-date)", err))
	}

	return strings.TrimSpace(string(output)), nil
}

func dockerBoundPorts(containerID string) (mapping map[string]string, err error) {
	defer essentials.AddCtxTo("docker inspect", &err)
	rawJSON, err := exec.Command("docker", "inspect", containerID).Output()
	if err != nil {
		return nil, err
	}
	var info []struct {
		NetworkSettings struct {
			Ports map[string][]struct {
				HostPort string
			}
		}
	}
	if err := json.Unmarshal(rawJSON, &info); err != nil {
		return nil, err
	}
	if len(info) != 1 {
		return nil, errors.New("unexpected number of results")
	}
	rawMapping := info[0].NetworkSettings.Ports
	mapping = map[string]string{}
	for containerPort, hostPorts := range rawMapping {
		if len(hostPorts) != 1 {
			return nil, errors.New("unexpected number of host ports")
		}
		mapping[containerPort] = hostPorts[0].HostPort
	}
	return
}

func connectDevTools(host string) (conn *chrome.Conn, err error) {
	var endpoints []*chrome.Endpoint
	for i := 0; i < 5; i++ {
		endpoints, err = chrome.Endpoints(host)
		if err == nil {
			break
		}
		// Give Chrome a moment to start up if needed.
		time.Sleep(time.Second)
	}
	if err != nil {
		return
	}

	for _, ep := range endpoints {
		if ep.Type == "page" {
			return chrome.NewConn(ep.WebSocketURL)
		}
	}

	return nil, errors.New("no Chrome page endpoint")
}
