package shell

import (
	"testing"
)

const (
	testScript1 = `
#!/bin/bash

echo "Hello From testScript1"

for f in *; do
	echo $f
done

echo $1
echo $2
`
)

func TestRun(t *testing.T) {
	Run("ls", "-l")
	Run(testScript1, "arg_1", "arg_2")
	Run("test/testscript2.sh", "arg_1", "arg_2")
}
