package configs

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func GenerateConfig() error {
	conf := &ConfigModel{
		WebsocketURL: "ws://127.0.0.1:13140",
		Debug:        false,
		SuperUsers:   []int64{1314000},
		Nickname:     []string{"ATRI", "Atri", "atri", "亚托莉", "アトリ"},
		GoCQHTTP: ConfigGoCQHTTPModel{
			Enabled:         false,
			Protocol:        "5",
			DownloadVersion: "latest",
		},
		SetuIsUseProxy: true,
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

func ConfigDealer() (*ConfigModel, error) {
	data := &ConfigModel{}
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
