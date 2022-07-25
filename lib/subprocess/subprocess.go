package subprocess

import (
	"bytes"
	"os/exec"
	"syscall"
)

func Run(execu string, args ...string) SubProcessResponse {
	var resp SubProcessResponse
	var outbuf, errbuf bytes.Buffer

	cmd := exec.Command(execu, args...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	resp.Stdout = outbuf.String()
	resp.Stderr = errbuf.String()

	err := cmd.Run()
	if err != nil {
		if exitError, isOK := err.(*exec.ExitError); isOK {
			resp.ExitCode = exitError.Sys().(syscall.WaitStatus).ExitStatus()
		} else {
			resp.ExitCode = 1
		}
	} else {
		resp.ExitCode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	}

	if resp.Stderr == "" && resp.ExitCode != 0 {
		resp.Stderr = err.Error()
	}

	return resp
}
