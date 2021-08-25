package waifu2x

import (
	"github.com/Kyomotoi/go-ATRI/service"
	"github.com/Kyomotoi/go-ATRI/utils"
	log "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"time"
)

var deepaikey string

func init() {
	conf, _ := utils.ConfigDealer()
	deepaikey = conf.DeepAiKey
	cmds := make(map[string]string)
	cmds["放大图片"] = "放大图片,仅适用于分辨率较小的图片"
	service.RegisterService("waifu2x", "放大图片", cmds)
	engine := zero.New()
	engine.OnCommand("放大图片", service.CheckBlock).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		if !service.IsServiceEnabled("waifu2x", ctx) {
			return
		}
		ctx.Send("请20秒内发送一张图片")
		next := zero.NewFutureEvent("message", 999, false, zero.CheckUser(ctx.Event.UserID))
		recv := next.Next()
		select {
		case <-time.After(time.Second * 20):
			ctx.Send("接收图片超时")
		case e := <-recv:
			nextCtx := &zero.Ctx{Event: e, State: zero.State{}}
			if nextCtx.Event.Message[0].Type == "image" {
				ctx.Send("放大中...")
				rt, err := waifu2x(nextCtx.Event.Message[0].Data["url"], deepaikey)
				if err != nil {
					log.Error("waifu2x: 请求api失败，error: ", err)
					ctx.Send("放大失败！")
				} else {
					ctx.Send("放大成功！")
					ctx.Send(message.Image(rt))
				}
			} else {
				ctx.Send("未接收到图片")
			}
		}
	})
}
