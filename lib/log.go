package lib

import (
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

const logDir = "data/logs/"

func init() {
	exi := IsDir(logDir)
	if !exi {
		err := os.MkdirAll(logDir, 0777)
		if err != nil {
			panic("log: 创建文件夹: " + logDir + " 失败, 请尝试手动创建")
		}
	}
}

func InitLogger() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "01-02 15:04:05",
		LogFormat:       "\033[37mATRI | %time% | %lvl% >> %msg% \n",
	})

	now := time.Now().Format("20060102-15")
	fileName := now + ".log"

	file, _ := os.OpenFile(logDir+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	mw := io.MultiWriter(os.Stdout, file)

	log.SetOutput(mw)
}
