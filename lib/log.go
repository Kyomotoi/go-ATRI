package lib

import (
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func InitLogger() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "01-02 15:04:05",
		LogFormat:       "ATRI | %time% | %lvl% >> %msg% \n",
	})
}
