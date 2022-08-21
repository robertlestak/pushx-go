package pushx

import (
	"bytes"
	"io"
	"os/exec"
)

func Pushx(input io.Reader, args []string) error {
	cmd := exec.Command("pushx", args...)
	cmd.Stdin = input
	var out bytes.Buffer
	var err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err
	return cmd.Run()
}
