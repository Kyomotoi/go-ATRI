package main

import (
	"ATRI/utils"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	zero "github.com/wdvxdr1123/ZeroBot"

	_ "ATRI/plugins/hitokoto"
)


func init() {
	log.SetFormatter(
		&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "ATRI | %time% | %lvl% >> %msg% \n",
		},
	)
	log.SetLevel(log.DebugLevel)

	_, err := rotatelogs.New(path.Join("logs", "%Y-%m-%d.log"), rotatelogs.WithRotationTime(time.Hour*24))
	if err != nil {
		log.Errorf("创建 logs 文件失败: %v", err)
	}

	// log 保存相关，待补充，在此之前无法保存log
	// log.AddHook()

	if utils.Exists("config.json") {
		log.Info("正在导入设置...")
		log.Warn("即将使用 config.json 内的配置启动 ATRI.")
	} else {
		log.Warn("检测为初次启动，已自动于同目录下生成 config.json 请配置并重新启动！")
		utils.WriteConfig()
		log.Warn("将在5秒后退出...")
		time.Sleep(time.Second * 5)
		os.Exit(0)
	}

	log.Info("アトリは、高性能ですから！")
	log.Info("Project: github.com/Kyomotoi/go-ATRI")
	time.Sleep(time.Second * 3)
}

func main() {
	conf := utils.ReadConfig("config.json")

	zero.Run(
		zero.Option{
			Host: conf.Host,
			Port: conf.Port,
			AccessToken: "",
			NickName: conf.NickName,
			CommandPrefix: conf.CommandPrefix,
			SuperUsers: conf.SuperUsers,
		},
	)

	select{}
}