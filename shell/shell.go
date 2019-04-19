package shell

import (
	"io"
	"os/exec"

	"github.com/ngtrimble/executil"
)

//RunCombined will execute 'sh -c shellCmd'. Stdout and and stderr will be written
//to stdout and stderr of the current process.
func RunCombined(shellCmd string) error {
	args := append([]string{"-c", shellCmd})
	cmd := exec.Command("sh", args...)
	return executil.StartWaitCombined(cmd)
}

//Run will execute 'sh -c shellCmd'. Stdout and stderr will be written to
//stdout and stderr respectively.
func Run(shellCmd string, stdout, stderr io.Writer) error {
	args := append([]string{"-c", shellCmd})
	cmd := exec.Command("sh", args...)
	return executil.StartWait(cmd, stdout, stderr)
}
