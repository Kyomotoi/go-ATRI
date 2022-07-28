package gocqhttp

import (
	"os"
	"path"
	"strconv"

	"github.com/Kyomotoi/go-ATRI/internal/gocqhttp/device"
	"github.com/Kyomotoi/go-ATRI/lib"
	log "github.com/sirupsen/logrus"
)

const resourceDIR = "data/protocols/gocqhttp/"

func init() {
	exi := lib.IsDir(resourceDIR)
	if !exi {
		err := os.MkdirAll(resourceDIR, 0777)
		if err != nil {
			log.Error("Driver: 创建文件夹失败: " + resourceDIR + " 请尝试手动创建")
		}
	}
}

func InitGoCQHTTP(version string, account int64, password string, host string, port string, proc int) error {
	var err error

	gocqV := getFileNameOfGoCQHTTP()
	if version == "" {
		version = "latest"
	}

	gocqDIR := resourceDIR + version
	exi := lib.IsDir(gocqDIR)
	if !exi {
		err = os.MkdirAll(gocqDIR, 0777)
		if err != nil {
			log.Error("Driver: 创建文件夹失败: " + gocqDIR + " 请尝试手动创建")
			log.Warn("Driver: 内置 gocqhttp 将失效")
			return err
		}
	}

	gocqPath := gocqDIR + "/" + gocqV
	if !lib.IsExists(gocqPath) {
		log.Info("Driver: 正在下载: " + gocqV)
		err = downloadGoCQHTTP(version)
		if err != nil {
			log.Warn("Driver: 内置 gocqhttp 将失效")
			return err
		}
		log.Info("Driver: gocqhttp 下载完成")
	}

	device.InitDevice(version, account, proc)

	log.Info("Driver: 正在初始化 gocqhttp 设置")
	err = InitConfig(version, strconv.Itoa(int(account)), password, host, port)
	if err != nil {
		log.Warn("Driver: 内置 gocqhttp 将失效")
		return err
	}

	wd, _ := os.Getwd()
	dc := path.Join(wd, gocqDIR)

	go lib.Run(dc+"/"+getFileNameOfGoCQHTTP(), dc, "-c", gocqDIR+"/config.yml")

	return nil
}
