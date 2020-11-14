package main

import (
	"gopkg.in/yaml.v3"
)

var data = `
# ATRI 基本配置
bot:
  # 反向ws地址，建议127.0.0.1
  host: "127.0.0.1"
  # 反向ws端口
  port: "8080"
  # 超级用户，填自己的QQ号，用“，”隔开
  superusers: [""]
  # 机器人昵称，建议不动
  nickname: ["ATRI", "Atri", "atri", "亚托莉", "アトリ"]
  # 命令触发头，建议不动，除非你知道你在做什么
  command_start: ""

# 接口相关配置
api:
  # 涩图接口，URL：https://api.lolicon.app/#/setu
  LoliconAPI:
  # 换脸接口，URL：https://www.faceplusplus.com.cn/
  FaceplusAPI:
  FaceplusSECRET:
  # 搜图接口，URL：https://saucenao.com/
  SauceNaoKEY:
`

type config struct {
	bot struct{
		host          string   `yaml: "127.0.0.1"`
		port          string   `yaml: "8080"`
		superusers    []string `yaml: [""]`
		nickname      []string `yaml: ["ATRI", "Atri", "atri", "亚托莉", "アトリ"]`
		command_start string   `yaml: ""`
	}

	api struct{
		LoliconAPI string     `yaml: ""`
		FaceplusAPI string    `yaml: ""`
		FaceplusSECRET string `yaml: ""`
		SauceNaoKEY string    `yaml: ""`
	}
}