package help

import (
	"fmt"
	"github.com/Kyomotoi/go-ATRI/service"
	"github.com/Kyomotoi/go-ATRI/utils"
	"io/ioutil"
	"strings"
)

func Menu() string {
	return menu
}

func About() string {
	conf, _ := utils.ConfigDealer()
	nickName := strings.Join(conf.Nickname, "、")
	version := service.Version()
	repo := fmt.Sprintf(projectInfo, nickName, version)
	return repo
}

func ServiceList() string {
	var tD []string
	files, _ := ioutil.ReadDir(service.FileDIR)
	for _, f := range files {
		fN := f.Name()
		tD = append(tD, strings.Replace(fN, ".json", "", -1))
	}
	servList := strings.Join(tD, "、")
	repo := fmt.Sprintf(serviceList, servList)
	return repo
}

func ServiceInfo(serv string) string {
	data := service.LoadServiceData(serv)
	servName := data.Service
	servDocs := data.Docs

	tD := utils.GetMapKeys(data.Commands)
	servCmds := strings.Join(tD, "、")

	repo := fmt.Sprintf(serviceInfo, servName, servDocs, servCmds)
	return repo
}

func CommandInfo(serv string, cmd string) string {
	data := service.LoadServiceData(serv)
	cmds := data.Commands
	cmdInfo := cmds[cmd]
	if cmdInfo == "" {
		return "请检查命令名是否输入正确（"
	}

	repo := fmt.Sprintf(commandInfo, cmd, cmdInfo)
	return repo
}

const menu = `
哦呀？~需要帮助？
@ 关于 -查看bot基本信息
@ 服务列表 -以查看所有可用服务
@ 帮助 [服务] -以查看对应服务帮助
@ 菜单 -以打开此页面
`

const projectInfo = `
唔...是来认识咱的么
可以称呼咱：%s
咱的型号是：%s
想进一步了解：
https://github.com/Kyomotoi/go-ATRI
`

const serviceList = `
咱搭载了以下服务~
%s
@ 帮助 [服务] -以查看对应帮助
如服务无响应，或许是权限不足的原因，或者已禁用
`

const serviceInfo = `
服务名：%s
说明：%s
可用命令：%s
Tip: @ 帮助 [服务] [命令] -以查看对应详细信息
`

const commandInfo = `
命令：%s
说明：%s
`
