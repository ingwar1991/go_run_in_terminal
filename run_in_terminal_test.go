package run_in_terminal

import "testing"

func TestRun(t *testing.T) {
    execName := "echo"
    args := []string{"hello", "world"}
    dirPath := ""
    noWait := false

    out, err := Run(execName, args, dirPath, noWait)
    if err != nil {
        t.Error(err)
    }

    if out != "hello world\n" {
        t.Error("Expected 'hello world\n', got", out)
    }
}
