package shell

import (
	"bytes"
	"fmt"
	"testing"
)

const (
	testScript1 = `#!/bin/bash

echo "Hello From testScript1"

for f in *; do
	echo $f
done
`
)

func TestRunCombined(t *testing.T) {
	RunCombined("ls -l")
	RunCombined(testScript1)
	RunCombined("test/testscript2.sh Hello World")
}

func TestRun(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	Run("ls -l", stdout, stderr, []string{})
	fmt.Printf("%s", stdout.String())
	fmt.Printf("%s", stderr.String())

	stdout.Reset()
	stderr.Reset()
	Run(testScript1, stdout, stderr, []string{})
	fmt.Printf("%s", stdout.String())
	fmt.Printf("%s", stderr.String())

	stdout.Reset()
	stderr.Reset()
	Run("test/testscript2.sh Hello World", stdout, stderr, []string{"FOO=bar"})
	fmt.Printf("%s", stdout.String())
	fmt.Printf("%s", stderr.String())
}
