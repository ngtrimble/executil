# executil

## Introduction

This package provides some convenient wrappers around os/exec. Specifically it solves the problem of running
external processes until they complete while continously streaming stdout and stderr to an io.Writer of your
choice. This package also makes piping other commands in sequence convenient. 

## Docs

[Read docs on https://godoc.org/github.com/ngtrimble/executil](https://godoc.org/github.com/ngtrimble/executil)

## Examples

```
cmd1 := exec.Command("echo", "hello world")
cmd2 := exec.Command("cat")
stdout, stderr, err := StartWaitPipe(nil, cmd1, cmd2)
if err != nil {
    fmt.Fprintf(os.Stderr, "%v", err)
} else {
    fmt.Printf("%s", stdout)
}
```

This is equivalent to running the following in the a shell


```
echo "hello world" | cat
```
