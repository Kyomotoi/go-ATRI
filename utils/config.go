package utils

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	WebsocketURL  string   `yaml:"WebsocketURL"`
	Debug         bool     `yaml:"Debug"`
	SuperUsers    []string `yaml:"SuperUsers"`
	Nickname      []string `yaml:"Nickname"`
	CommandPrefix string   `yaml:"CommandPrefix"`
	AccessToken   string   `yaml:"AccessToken"`
	SauceNaoKey   string   `yaml:"SauceNaoKey"`
}

func generateConfig() error {
	conf := &Config{
		WebsocketURL:  "ws://127.0.0.1:13140",
		Debug:         false,
		SuperUsers:    []string{"1314000"},
		Nickname:      []string{"ATRI", "Atri", "atri", "亚托莉", "アトリ"},
		CommandPrefix: "",
		AccessToken:   "",
		SauceNaoKey:   "",
	}
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("config.yml", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig() error {
	if IsExists("config.yml") {
		log.Info("正在导入设置...")
		log.Info("即将使用 config.yml 内的配置启动ATRI")
		time.Sleep(time.Second*3)
	} else {
		log.Warning("检查为初次启动，已自动于同目录下生成 config.yml，请配置并重新启动！")
		err := generateConfig()
		if err != nil {
			log.Error("无法创建文件：config.yml，请确认是否给足系统权限")
			return err
		}
		log.Warning("将于5秒后退出")
		time.Sleep(time.Second * 5)
		os.Exit(1)
	}
	return nil
}

func ConfigDealer() (Config, error) {
	var data Config

	content, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Error("无法读取 config.yml")
		return data, err
	}
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		log.Error("解析 config.yml 失败，请检查格式、内容是否输入正确")
		return data, err
	}
	return data, nil
}
