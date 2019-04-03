package executil

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

//StartWaitCombinedStdout executes StartWait using os.Stdout as output destinations
func StartWaitCombinedStdout(cmd *exec.Cmd) error {
	return StartWait(cmd, os.Stdout, os.Stdout)
}

//StartWaitCombined executes StartWait using os.Stdout and os.Stderr as output
//destinations
func StartWaitCombined(cmd *exec.Cmd) error {
	return StartWait(cmd, os.Stdout, os.Stderr)
}

//StartWaitPipe executes each cmd in order and pipes stdout from one cmd to the next.
//stdin can be nil if you do not want to pipe data to stdin on the first command.
func StartWaitPipe(stdin io.Reader, cmds ...*exec.Cmd) (io.Reader, io.Reader, error) {
	var stdout *bytes.Buffer
	var stderr *bytes.Buffer
	for _, cmd := range cmds {
		stdout = &bytes.Buffer{}
		stderr = &bytes.Buffer{}
		err := StartWaitStdin(cmd, stdin, stdout, stderr)
		if err != nil {
			return stdout, stderr, err
		}
		stdin = stdout
	}
	return stdout, stderr, nil
}

//StartWait exec's cmd, waits and streams outputs continously to stdoutDst and stderrDst.
func StartWait(cmd *exec.Cmd, stdoutDst, stderrDst io.Writer) error {
	return StartWaitStdin(cmd, nil, stdoutDst, stderrDst)
}

//StartWaitStdin exec's cmd with input from stdin, waits and streams outputs continously to
//stdoutDst and stderrDst.
func StartWaitStdin(cmd *exec.Cmd, stdin io.Reader, stdoutDst, stderrDst io.Writer) error {
	if stdin != nil {
		cmd.Stdin = stdin
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	go func() {
		io.Copy(stdoutDst, stdout)
	}()

	go func() {
		io.Copy(stderrDst, stderr)
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
