package muniverse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"

	"github.com/unixpickle/essentials"
	"github.com/unixpickle/muniverse/chrome"
)

const portRange = "9000-9999"

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
	containerID string
	devConn     *chrome.Conn

	// Used to garbage collect the container if we
	// exit ungracefully.
	killSocket net.Conn
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
		containerID: id,
		devConn:     conn,
		killSocket:  killSock,
	}, nil
}

// Close cleans up the resources associated with the
// environment.
func (e *Env) Close() error {
	errs := []error{
		e.devConn.Close(),
		exec.Command("docker", "kill", e.containerID).Run(),
	}

	// TODO: look into if this can ever produce an error,
	// since the container might already have closed the
	// socket by now.
	//
	// We don't close this *before* stopping the container
	// since `docker kill` might fail if the container
	// already died and was cleaned up.
	e.killSocket.Close()

	// Any calls after Close() should trigger simple panics.
	e.devConn = nil
	e.killSocket = nil

	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
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
		container,
		fmt.Sprintf("--window-size=%d,%d", spec.Width, spec.Height),
		"http://localhost/" + spec.Name,
	}

	output, err := exec.Command("docker", args...).Output()
	if err != nil {
		return "", essentials.AddCtx("docker run", err)
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
