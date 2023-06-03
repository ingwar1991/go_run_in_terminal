package run_in_terminal

import (
	"bytes"
	"os/exec"
)

func Run(execName string, args []string, dirPath string, noWait bool) (string, error) {
	execPath, err := exec.LookPath(execName)
	if err != nil {
		return "", err
	}

	args = append([]string{execPath}, args...)
	cmd := exec.Cmd{
		Path: execPath,
		Args: args,
	}

	if dirPath != "" {
		cmd.Dir = dirPath
	}

	if noWait {
		err = cmd.Start()

		return "", err
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()

	return out.String(), err
}
