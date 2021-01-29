package utils

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// Exists 检测文件是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// ReadExists 读取文件，错误时返回None
func ReadExists(path string) string {
	info, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("读取文件" + path + "失败.", err.Error())
		return ""
	}
	return string(info)
}