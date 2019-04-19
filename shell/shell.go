package shell

import (
	"fmt"
	"os/exec"

	"github.com/ngtrimble/executil"
)

//Run will run cmd using sh on
func Run(shellCmd string, args ...string) error {
	args = append([]string{"-c", shellCmd}, args...)
	fmt.Printf("%v\n", args)
	cmd := exec.Command("sh", args...)
	return executil.StartWaitCombined(cmd)
}
