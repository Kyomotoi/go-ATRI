package models

import "os"

// 判断 path 文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断 path 是否为文件夹
func IsDir(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return false
	}
	return p.IsDir()
}

// 判断 path 是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
