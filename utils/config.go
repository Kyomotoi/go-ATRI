package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// Config 设置参数
type Config struct {
	Version       string
	Host          string
	Port          string
	Debug         bool
	SuperUsers    []string
	NickName      []string
	CommandPrefix string
}

// WriteConfig 写入配置
func WriteConfig() {
	info := Config{
		Version:       "YHN-001-C01",
		Host:          "127.0.0.0",
		Port:          "25565",
		Debug:         false,
		SuperUsers:    []string{"1172294279"},
		NickName:      []string{"ATRI", "atri", "亚托莉"},
		CommandPrefix: "/",
	}

	file, err := os.Create("config.json")
	if err != nil {
		log.Error("文件创建失败", err.Error())
		os.Exit(-1)
	}
	defer file.Close()

	msg, _ := json.MarshalIndent(info, "", "    ")
	err = ioutil.WriteFile("config.json", msg, 0644)
	if err != nil {
		log.Error("写入失败", err.Error())
		os.Exit(-1)
	} else {
		log.Debug("写入配置成功")
	}
}

// ReadConfig 读取配置
func ReadConfig(path string) *Config {
	var goJSONcnmNMSL *Config
	err := json.Unmarshal([]byte(ReadExists(path)), &goJSONcnmNMSL)
	if err != nil {
		log.Error("读取配置失败，请检查拼写.", err.Error())
		log.Warn("将在5秒后退出...")
		time.Sleep(time.Second * 5)
	}
	return goJSONcnmNMSL
}
