//go:build !windows
// +build !windows

// Fork from: https://github.com/Mrs4s/go-cqhttp/blob/master/global/terminal

package terminal

func RunningByDoubleClick() bool {
	return false
}
