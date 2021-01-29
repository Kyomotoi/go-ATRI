package hitokoto

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	zero.RegisterPlugin(hitokoto{})
}

type hitokoto struct{}

func (hitokoto) GetPluginInfo() zero.PluginInfo {
	return zero.PluginInfo{
		Author:     "Kyomotoi",
		PluginName: "Hitokoto",
		Version:    "0.0.1",
		Details:    "一言",
	}
}


func (hitokoto) Start() {
	zero.OnCommand("test").Handle(
		func(matcher *Matcher, event Event, state State) zero.Response {
			zero.Send(event, "Hello Golang!")
			return zero.FinishResponse
		},
	)
}

// handleEcho 插件处理逻辑 
func handleEcho(_ *zero.Matcher, event zero.Event, state zero.State) zero.Response {
	zero.Send(event, "Hello Golang!")
	return zero.FinishResponse
}