package lib

import "runtime"

// GetPlatformType 获取系统架构
func GetPlatformType() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}
