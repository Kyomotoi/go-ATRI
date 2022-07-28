package gocqhttp

import (
	_ "embed"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

//go:embed gocq_default_config.yml
var gocqdefConfig string

var confPath string

func genGocqConfig() error {
	sbg := strings.Builder{}
	sbg.WriteString(gocqdefConfig)

	err := os.WriteFile(confPath, []byte(sbg.String()), 0644)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig(version string, account string, password string, host string, port string) error {
	confPath = resourceDIR + version + "/config.yml"
	err := genGocqConfig()
	if err != nil {
		log.Fatal("Driver: 生成 gocqhttp 配置文件失败")
		return err
	}

	var data string

	rawData, _ := os.ReadFile(confPath)
	st := string(rawData)
	data = strings.Replace(st, "noraccount", account, -1)
	data = strings.Replace(data, "norpassword", password, -1)
	data = strings.Replace(data, "norhost", host, -1)
	data = strings.Replace(data, "norport", port, -1)

	err = os.WriteFile(confPath, []byte(data), 0644)
	if err != nil {
		log.Fatal("Driver: 处理 gocqhttp 配置文件失败")
		return err
	}

	log.Info("Driver: gocqhttp 配置文件初始化完成")
	return nil
}
