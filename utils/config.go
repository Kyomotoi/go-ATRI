package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	WebsocketURL   string   `yaml:"WebsocketURL"`
	Debug          bool     `yaml:"Debug"`
	SuperUsers     []int64  `yaml:"SuperUsers"`
	Nickname       []string `yaml:"Nickname"`
	CommandPrefix  string   `yaml:"CommandPrefix"`
	AccessToken    string   `yaml:"AccessToken"`
	SetuIsUseProxy bool     `yaml:"SetuIsUseProxy"`
	SauceNaoKey    string   `yaml:"SauceNaoKey"`
}

func GenerateConfig() error {
	conf := &Config{
		WebsocketURL:   "ws://127.0.0.1:13140",
		Debug:          false,
		SuperUsers:     []int64{1314000},
		Nickname:       []string{"ATRI", "Atri", "atri", "亚托莉", "アトリ"},
		CommandPrefix:  "",
		AccessToken:    "",
		SetuIsUseProxy: true,
		SauceNaoKey:    "",
	}
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	err = os.WriteFile("config.yml", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ConfigDealer() (*Config, error) {
	data := &Config{}
	content, err := os.ReadFile("config.yml")
	if err != nil {
		log.Error("未找到 config.yml")
		return data, err
	}
	err = yaml.Unmarshal(content, data)
	if err != nil {
		panic("解析 config.yml 失败, 请检查格式、内容是否输入正确")
	}
	return data, nil
}
