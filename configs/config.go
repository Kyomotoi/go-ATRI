package configs

import (
	_ "embed"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

//go:embed default_config.yml
var defConfig string

func genConfig() error {
	sb := strings.Builder{}
	sb.WriteString(defConfig)

	err := os.WriteFile("config.yml", []byte(sb.String()), 0644)
	if err != nil {
		return err
	}
	return nil
}

func Parse() *Config {
	content, err := os.ReadFile("config.yml")
	if err != nil {
		err = genConfig()
		log.Warn("未检测到 config.yml, 已自动于同目录生成, 请配置并重新启动")
		if err != nil {
			log.Fatal("无法创建文件: config.yml, 请确认是否给足系统权限")
		}
		log.Warn("将于5秒后退出...")
		time.Sleep(time.Second * 5)
		os.Exit(1)
	}

	data := &Config{}
	err = yaml.Unmarshal(content, data)
	if err != nil {
		panic("解析 config.yml 失败, 请检查格式、内容是否输入正确")
	}
	return data
}
