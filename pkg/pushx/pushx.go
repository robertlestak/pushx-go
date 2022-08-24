package pushx

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
)

func findExec(name string) (string, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return "", err
	}
	if path == "" {
		path = "/usr/local/bin/" + name
	}
	return path, nil
}

func Pushx(input io.Reader, args []string) error {
	path, e := findExec("pushx")
	if e != nil {
		return e
	}
	cmd := exec.Command(path, args...)
	cmd.Env = os.Environ()
	cmd.Stdin = input
	var out bytes.Buffer
	var err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err
	if e != nil {
		var err bytes.Buffer
		err.WriteString(err.String())
		err.WriteString(out.String())
		return errors.New(err.String())
	}
	return nil
}
