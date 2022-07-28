package main

import (
	// _ "github.com/Kyomotoi/go-ATRI/plugins/help"
	// _ "github.com/Kyomotoi/go-ATRI/plugins/manage"

	"fmt"
	"strconv"
	"time"

	_ "github.com/Kyomotoi/go-ATRI/plugins/anti_effort"
	_ "github.com/Kyomotoi/go-ATRI/plugins/setu"

	"github.com/Kyomotoi/go-ATRI/configs"
	"github.com/Kyomotoi/go-ATRI/internal"
	"github.com/Kyomotoi/go-ATRI/lib"
	"github.com/Kyomotoi/go-ATRI/lib/terminal"

	log "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

var config *configs.Config

func init() {
	lib.InitLogger()

	log.Info("项目地址: https://github.com/Kyomotoi/go-ATRI")
	log.Info("当前版本：" + internal.Version())
	log.Info("后宫裙: 567297659")

	config = configs.Parse()
	log.Info("config.yml 加载成功")

	if config.Bot.Debug {
		log.Info("DEBUG 已启用")
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if terminal.RunningByDoubleClick() {
		log.Warning("不建议直接双击运行本程序, 这将导致一些非可预料后果, 请通过控制台启动本程序")
		log.Warning("将等待10秒后启动")
		time.Sleep(time.Second * 10)
	}

	timelocal, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = timelocal

	lib.InitSchedule()
	log.Info("定时任务已启动")

	if config.Driver.Gocqhttp.Enabled {
		log.Info("Driver: 初始化协议中")
		err := internal.InitDriver(
			config.Driver.Gocqhttp.DownloadVersion,
			config.Driver.Gocqhttp.Account,
			config.Driver.Gocqhttp.Password,
			config.Bot.Host,
			strconv.Itoa(config.Bot.Port),
			config.Driver.Gocqhttp.Protocol,
		)
		if err != nil {
			log.Warn("Driver: 初始化内置协议端失败. 请使用外置协议端连接")
		}
		log.Info("Driver: 协议初始化完成")
	}

	log.Info("アトリは、高性能ですから！")
}

func main() {
	wsClientURL := fmt.Sprintf("ws://%s:%d", config.Bot.Host, config.Bot.Port)
	zero.RunAndBlock(zero.Config{
		NickName:      config.Bot.Nickname,
		CommandPrefix: config.Bot.CommandPrefix,
		SuperUsers:    config.Bot.Superusers,
		Driver: []zero.Driver{
			driver.NewWebSocketClient(wsClientURL, config.Bot.AccessToken),
		},
	}, nil)
	select {}
}
