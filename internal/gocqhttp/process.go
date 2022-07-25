package gocqhttp

import (
	"os"

	"github.com/Kyomotoi/go-ATRI/lib"
	log "github.com/sirupsen/logrus"
)

const resourceDIR = "data/protocols/gocqhttp/"

func init() {
	exi := lib.IsDir(resourceDIR)
	if !exi {
		err := os.MkdirAll(resourceDIR, 0777)
		if err != nil {
			log.Error("创建文件夹失败: " + resourceDIR + " 请尝试手动创建")
		}
	}
}

func InitGoCQHTTP(v string) error {
	gocqV := getFileNameOfGoCQHTTP()
	if v == "" {
		v = "latest"
	}

	gocqDIR := resourceDIR + "/" + v
	exi := lib.IsDir(gocqDIR)
	if !exi {
		err := os.MkdirAll(gocqDIR, 0777)
		if err != nil {
			log.Error("创建文件夹失败: " + gocqDIR + " 请尝试手动创建")
			log.Warn("init gocq: 内置 gocqhttp 将失效")
			return err
		}
	}

	gocqPath := gocqDIR + "/" + gocqV
	if !lib.IsExists(gocqPath) {
		log.Info("init gocq: 正在下载: " + gocqV)
		err := downloadGoCQHTTP(v)
		if err != nil {
			log.Warn("init gocq: 内置 gocqhttp 将失效")
			return err
		}
		log.Info("init gocq: gocqhttp 下载完成")
	}
	log.Info("init gocq: 初始化完成")
	return nil
}
