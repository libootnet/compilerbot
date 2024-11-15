package src

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

func CreateVM(id, images, extension string) (string, int, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", 1, err
	}

	ctx := context.Background()

	// images := "docker.io/library/python:3.9"

	_, _, err = cli.ImageInspectWithRaw(ctx, images)
	if err != nil {
		if client.IsErrNotFound(err) {
			// install
			out, err := cli.ImagePull(ctx, images, image.PullOptions{})
			if err != nil {
				return "", 1, err
			}
			defer out.Close()
			io.Copy(os.Stdout, out)
		} else {
			return "", 1, err
		}
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", 1, err
	}

	path := "/scripts/" + id + "." + extension

	command := Methods(extension, path)

	config := &container.Config{
		// Image: "python:3.9",
		Image: LanguageType(images),
		Cmd:   command,
		Tty:   true,
		/*Env: []string{
			"GOCACHE=/go-tmp/go-cache",
			"GOPATH=/go-tmp/go-path",
			"GOTMPDIR=/go-tmp",
		},*/
	}

	hostConfig := &container.HostConfig{
		NetworkMode: "none",
		Privileged:  false,
		// ReadonlyRootfs: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: dir + "/scripts",
				Target: "/scripts",
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, "")
	if err != nil {
		return "", 1, err
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", 1, err
	}
	// fmt.Println(resp.ID)

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	timeoutCh := time.After(20 * time.Second)

	var errWait error
	var timeout bool
	select {
	case err := <-errCh:
		errWait = err
	case <-statusCh:
		errWait = nil
	case <-timeoutCh:
		timeout = true
	}

	if timeout {
		if err := cli.ContainerStop(ctx, resp.ID, container.StopOptions{}); err != nil {
			return "", 1, err
		}
		if err := cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true}); err != nil {
			return "", 1, err
		}
		return "Time Out", 1, nil
	}

	if errWait != nil {
		return "", 1, errWait
	}

	logs, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return "", 1, err
	}

	inspect, err := cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		return "", 1, err
	}

	logOutput, err := io.ReadAll(logs)
	if err != nil {
		return "", 1, err
	}
	// fmt.Println(string(logOutput))

	if err := cli.ContainerStop(ctx, resp.ID, container.StopOptions{}); err != nil {
		return "", 1, err
	}
	if err := cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{}); err != nil {
		return "", 1, err
	}

	// fmt.Println(resp.ID)

	return string(logOutput), inspect.State.ExitCode, nil
}
