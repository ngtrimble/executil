package executil

import (
	"io"
	"os"
	"os/exec"
)

//StartWaitCombinedStdout is used to run a command and have its outputs stream
//continously to stdout of the current process.
func StartWaitCombinedStdout(cmd *exec.Cmd) error {
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	go func() {
		io.Copy(os.Stdout, stdout)
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
