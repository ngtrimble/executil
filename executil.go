package executil

import (
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

//StartWait exec's cmd, waits and streams outputs continously stdoutDst and stderrDst.
func StartWait(cmd *exec.Cmd, stdoutDst io.Writer, stderrDst io.Writer) error {
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
