package plugins

import (
	"github.com/wdvxdr1123/ZeroBot"
	"github.com/Kyomotoi/go-ATRI/models"
	"strings"
)

func init() {
	chat := ChatPlugin{}
	zero.RegisterPlugin(chat)
}

type ChatPlugin struct {}

func (ChatPlugin) GetPluginInfo() zero.PluginInfo {
	return zero.PluginInfo{
		Author: "kyomotoi",
		PluginName: "Chat",
		Version: "0.0.0.1",
		Details: "处理日常聊天",
	}
}

func (ChatPlugin) Start() {
	zero.OnCommand("骂我").Handle(func(_ *zero.Matcher, event zero.Event, state zero.State) zero.Response {
		msg := models.HttpGET("https://nmsl.shadiao.app/api.php?level=min&lang=zh_cn")
		zero.Send(event, msg)
		return zero.FinishResponse
	})

	zero.OnMessage().Handle(func(_ *zero.Matcher, event zero.Event, state zero.State) zero.Response {
		msg := event.Message.CQString()

		var nickname []string
		nickname[0] = "ATRI"
		nickname[1] = "atri"
		nickname[2] = "亚托莉"
		nickname[3] = "アトリ"

		if strings.ContainsAny(msg, "萝卜子") {
			zero.Send(event, "萝卜子是对咱的蔑称！")
		} else if models.StringInSlice(msg, nickname) {
			zero.Send(event, "叫咱有啥事吗w?")
		}

		return zero.FinishResponse
	})
}