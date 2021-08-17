package main

import (
	_ "github.com/Kyomotoi/go-ATRI/plugins/chat"
	_ "github.com/Kyomotoi/go-ATRI/plugins/help"
	_ "github.com/Kyomotoi/go-ATRI/plugins/manage"
	_ "github.com/Kyomotoi/go-ATRI/plugins/setu"
	"github.com/Kyomotoi/go-ATRI/service"
	"github.com/Kyomotoi/go-ATRI/utils"
	"github.com/Kyomotoi/go-ATRI/utils/terminal"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

func init() {

	utils.InitLogger()

	log.Info("项目地址：https://github.com/Kyomotoi/go-ATRI")
	log.Info("当前版本："+service.Version())
	log.Info("后宫裙：567297659")

	err := utils.InitConfig()
	if err != nil {
		log.Warning("初始化 config.yml 失败")
		log.Warning("将于5秒后退出")
		time.Sleep(time.Second * 5)
		os.Exit(1)
	}

	conf, err := utils.ConfigDealer()
	if err != nil {
		log.Warning("处理 config.yml 失败")
		log.Warning("将于5秒后退出")
		time.Sleep(time.Second * 5)
		os.Exit(1)
	}

	if conf.Debug {
		log.Info("DEBUG 已启用")
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if terminal.RunningByDoubleClick() {
		log.Warning("不建议直接双击运行本程序，这将导致一些非可预料后果，请通过控制台启动本程序")
		log.Warning("将等待10秒后启动")
		time.Sleep(time.Second*10)
	}

	log.Info("アトリは、高性能ですから！")
}

func main() {
	conf, _ := utils.ConfigDealer()
	zero.Run(zero.Config{
		NickName:      conf.Nickname,
		CommandPrefix: conf.CommandPrefix,
		SuperUsers:    conf.SuperUsers,
		Driver:        []zero.Driver{
			&driver.WSClient{
				Url:         conf.WebsocketURL,
				AccessToken: conf.AccessToken,
			},
		},
	})
	select {}
}
