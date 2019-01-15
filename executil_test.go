package executil

import (
	"os/exec"
	"runtime"
	"testing"
)

func TestStartWaitCombinedStdout(t *testing.T) {
	x := runtime.NumGoroutine()

	cmd := exec.Command("test/runsleep.sh")
	err := StartWaitCombinedStdout(cmd)
	if err != nil {
		t.Logf("%v", err)
		t.Fail()
	}

	y := runtime.NumGoroutine()
	if x != y {
		t.Logf("You're leaking goroutines...")
		t.Fail()
	}
}
