package muniverse

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/unixpickle/essentials"
	"golang.org/x/net/context"
)

func dockerRun(ctx context.Context, container, volume string,
	spec *EnvSpec) (id string, err error) {
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
	}
	if volume != "" {
		if strings.Contains(volume, ":") {
			return "", errors.New("path contains colons: " + volume)
		}
		args = append(args, "-v", volume+":/downloaded_games")
	}
	args = append(args, container,
		fmt.Sprintf("--window-size=%d,%d", spec.Width, spec.Height))

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

func dockerIPAddress(ctx context.Context, containerID string) (addr string, err error) {
	if runtime.GOOS != "windows" {
		return "localhost", nil
	}
	defer essentials.AddCtxTo("get IP of Docker VM", &err)
	cmd := exec.CommandContext(ctx, "docker-machine", "ip", "default")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	cleanIP := strings.TrimSpace(string(output))

	// Will detect pretty much any error message.
	if strings.ContainsAny(cleanIP, " \n\t") {
		return "", errors.New("unexpected output: " + string(output))
	}

	return cleanIP, nil
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
