package muniverse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/muniverse/chrome"
)

const (
	portRange        = "9000-9999"
	defaultContainer = "unixpickle/muniverse:0.42.0"
)

const (
	callTimeout           = time.Minute * 2
	chromeConnectAttempts = 10
)

// This error message occurs very infrequently when doing
// `docker run` on my machine running Ubuntu 16.04.1.
const occasionalDockerErr = "Error response from daemon: device or resource busy."

// An Env controls and observes an environment.
//
// It is not safe to run an methods on an Env from more
// than one Goroutine at a time.
type Env interface {
	// Spec returns details about the environment.
	Spec() *EnvSpec

	// Reset resets the environment to a start state.
	Reset() error

	// Step sends the given events and advances the
	// episode by the given amount of time.
	//
	// If done is true, then the episode has ended.
	// After an episode ends, Reset must be called once
	// before Step may be called again.
	// However, observations may be made even after the
	// episode has ended.
	//
	// Typical event types are *chrome.MouseEvent and
	// *chrome.KeyEvent.
	Step(t time.Duration, events ...interface{}) (reward float64,
		done bool, err error)

	// Observe produces an observation for the current
	// state of the environment.
	Observe() (Obs, error)

	// Close cleans up resources used by the environment.
	//
	// After Close is called, the Env should not be used
	// anymore by any Goroutine.
	Close() error

	// Log returns internal log messages.
	// For example, it might return information about 404
	// errors.
	//
	// The returned list is a copy and may be modified by
	// the caller.
	Log() []string
}

type rawEnv struct {
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
func NewEnv(spec *EnvSpec) (Env, error) {
	return NewEnvContainer(defaultContainer, spec)
}

// NewEnvContainer creates a new environment inside a
// Docker container of the given name.
func NewEnvContainer(container string, spec *EnvSpec) (env Env, err error) {
	defer essentials.AddCtxTo("create environment", &err)

	ctx, cancel := callCtx()
	defer cancel()

	var id string

	// Retry as a workaround for an occasional error given
	// by `docker run`.
	for i := 0; i < 3; i++ {
		id, err = dockerRun(ctx, container, spec)
		if err == nil || !strings.Contains(err.Error(), occasionalDockerErr) {
			break
		}
	}

	if err != nil {
		return
	}

	ports, err := dockerBoundPorts(ctx, id)
	if err != nil {
		return
	}

	conn, err := connectDevTools(ctx, "localhost:"+ports["9222/tcp"])
	if err != nil {
		return
	}

	killSock, err := (&net.Dialer{}).DialContext(ctx, "tcp",
		"localhost:"+ports["1337/tcp"])
	if err != nil {
		conn.Close()
		return
	}

	return &rawEnv{
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
func NewEnvChrome(host, gameHost string, spec *EnvSpec) (env Env, err error) {
	defer essentials.AddCtxTo("create environment", &err)

	ctx, cancel := callCtx()
	defer cancel()

	conn, err := connectDevTools(ctx, host)
	if err != nil {
		return
	}

	return &rawEnv{
		spec:     *spec,
		gameHost: gameHost,
		devConn:  conn,
	}, nil
}

func (r *rawEnv) Spec() *EnvSpec {
	res := r.spec
	return &res
}

func (r *rawEnv) Reset() (err error) {
	defer essentials.AddCtxTo("reset environment", &err)

	ctx, cancel := callCtx()
	defer cancel()

	err = r.devConn.NavigateSafe(ctx, r.envURL())
	if err != nil {
		return
	}

	initCode := "window.muniverse.init(" + r.spec.Options + ");"
	err = r.devConn.EvalPromise(ctx, initCode, nil)
	if err != nil {
		return
	}

	err = r.devConn.EvalPromise(ctx, "window.muniverse.score();", &r.lastScore)
	err = essentials.AddCtx("get score", err)

	return
}

func (r *rawEnv) Step(t time.Duration, events ...interface{}) (reward float64,
	done bool, err error) {
	defer essentials.AddCtxTo("step environment", &err)

	ctx, cancel := callCtx()
	defer cancel()

	for _, event := range events {
		switch event := event.(type) {
		case *chrome.MouseEvent:
			err = r.devConn.DispatchMouseEvent(ctx, event)
		case *chrome.KeyEvent:
			err = r.devConn.DispatchKeyEvent(ctx, event)
		default:
			err = fmt.Errorf("unsupported event type: %T", event)
		}
		if err != nil {
			return
		}
	}

	millis := int(t / time.Millisecond)
	timeStr := strconv.Itoa(millis)
	err = r.devConn.EvalPromise(ctx, "window.muniverse.step("+timeStr+");", &done)
	if err != nil {
		return
	}

	lastScore := r.lastScore
	err = r.devConn.EvalPromise(ctx, "window.muniverse.score();", &r.lastScore)
	if err != nil {
		err = essentials.AddCtx("get score", err)
		return
	}
	reward = r.lastScore - lastScore

	return
}

func (r *rawEnv) Observe() (obs Obs, err error) {
	ctx, cancel := callCtx()
	defer cancel()

	pngData, err := r.devConn.ScreenshotPNG(ctx)
	if err != nil {
		return
	}
	return pngObs(pngData), nil
}

func (r *rawEnv) Close() (err error) {
	defer essentials.AddCtxTo("close environment", &err)

	ctx, cancel := callCtx()
	defer cancel()

	errs := []error{
		r.devConn.Close(),
	}
	if r.containerID != "" {
		_, e := dockerCommand(ctx, "kill", r.containerID)
		errs = append(errs, e)
	}

	if r.killSocket != nil {
		// TODO: look into if this can ever produce an error,
		// since the container might already have closed the
		// socket by now.
		//
		// We don't close this *before* stopping the container
		// since `docker kill` might fail if the container
		// already died and was cleaned up.
		r.killSocket.Close()
	}

	// Any calls after Close() should trigger simple errors.
	r.devConn = nil
	r.killSocket = nil

	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *rawEnv) Log() []string {
	return r.devConn.ConsoleLog()
}

func (r *rawEnv) envURL() string {
	baseName := r.spec.Name
	if r.spec.VariantOf != "" {
		baseName = r.spec.VariantOf
	}
	return "http://" + r.gameHost + "/" + baseName
}

func callCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), callTimeout)
}

func dockerRun(ctx context.Context, container string, spec *EnvSpec) (id string,
	err error) {
	args := []string{
		"run",
		"-p",
		portRange + ":9222",
		"-p",
		portRange + ":1337",
		"--shm-size=200m",
		"-d",   // Run in detached mode.
		"--rm", // Automatically delete the container.
		"-i",   // Give netcat a stdin to read from.
		container,
		fmt.Sprintf("--window-size=%d,%d", spec.Width, spec.Height),
	}

	output, err := dockerCommand(ctx, args...)
	if err != nil {
		return "", essentials.AddCtx("docker run",
			fmt.Errorf("%s (make sure docker is up-to-date)", err))
	}

	return strings.TrimSpace(string(output)), nil
}

func dockerBoundPorts(ctx context.Context,
	containerID string) (mapping map[string]string, err error) {
	defer essentials.AddCtxTo("docker inspect", &err)
	rawJSON, err := dockerCommand(ctx, "inspect", containerID)
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

var dockerLock sync.Mutex

func dockerCommand(ctx context.Context, args ...string) (output []byte, err error) {
	dockerLock.Lock()
	defer dockerLock.Unlock()
	output, err = exec.CommandContext(ctx, "docker", args...).Output()
	if err != nil {
		if eo, ok := err.(*exec.ExitError); ok && len(eo.Stderr) > 0 {
			stderrMsg := strings.TrimSpace(string(eo.Stderr))
			err = fmt.Errorf("%s: %s", eo.String(), stderrMsg)
		}
	}
	return
}

func connectDevTools(ctx context.Context, host string) (conn *chrome.Conn,
	err error) {
	for i := 0; i < chromeConnectAttempts; i++ {
		conn, err = attemptDevTools(ctx, host)
		if err == nil {
			return
		}
		select {
		case <-time.After(time.Second):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	return
}

func attemptDevTools(ctx context.Context, host string) (conn *chrome.Conn,
	err error) {
	endpoints, err := chrome.Endpoints(ctx, host)
	if err != nil {
		return
	}

	for _, ep := range endpoints {
		if ep.Type == "page" && ep.WebSocketURL != "" {
			return chrome.NewConn(ctx, ep.WebSocketURL)
		}
	}

	return nil, errors.New("no Chrome page endpoint")
}
