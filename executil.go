package executil

import (
	"io"
	"os"
	"os/exec"
)

//StartWaitCombinedStdout is used to run a command and have its outputs stream
//continously to stdout of the current process.
func StartWaitCombinedStdout(cmd *exec.Cmd) error {
	return StartWait(cmd, os.Stdout, os.Stdout)
}

//StartWaitCombined is used to run a command and have its outputs stream
//continously to stdout and stderr of the current process.
func StartWaitCombined(cmd *exec.Cmd) error {
	return StartWait(cmd, os.Stdout, os.Stderr)
}

//StartWait is used to run a command have its outputs stream continously to
//stdoutDst and stderrDst.
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
