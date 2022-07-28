package lib

import (
	"os"
	"os/exec"
)

func Run(execu string, runDir string, args ...string) error {
	cmd := exec.Command(execu, args...)
	cmd.Dir = runDir
	cmd.Args = append([]string{}, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	return cmd.Wait()
}
