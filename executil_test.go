package executil

import (
	"bytes"
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

func TestStartWaitCombined(t *testing.T) {
	x := runtime.NumGoroutine()

	cmd := exec.Command("test/run.sh")
	err := StartWaitCombined(cmd)
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

func TestStartWait(t *testing.T) {
	x := runtime.NumGoroutine()

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	cmd := exec.Command("test/run.sh")
	err := StartWait(cmd, stdout, stderr)
	if err != nil {
		t.Logf("%v", err)
		t.Fail()
	}

	if stdout.String() != "stdout\n" {
		t.Logf("This is captured from stdout: %s", stdout.String())
		t.Fail()
	}

	if stderr.String() != "stderr\n" {
		t.Logf("This is captured from stderr: %s", stderr.String())
		t.Fail()
	}

	y := runtime.NumGoroutine()
	if x != y {
		t.Logf("You're leaking goroutines...")
		t.Fail()
	}
}
