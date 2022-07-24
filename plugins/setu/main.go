package setu

import (
	"time"

	"github.com/Kyomotoi/go-ATRI/service"
	zero "github.com/wdvxdr1123/ZeroBot"
)

// var limit = rate.NewLimiter(time.Minute*2, 1)
const serviceName = "涩图"

func init() {
	setu := service.NewService(serviceName, "hso!", false, "", service.CheckBlock, service.IsServiceEnabled(serviceName))

	_ = setu.OnCommand("来张涩图", "随机涩图, 冷却2分钟", []string{"涩图来", "来点涩图", "来份涩图"}).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				msg, setu, err := GetSetu("")
				if err != nil {
					ctx.Send("冲不起惹...")
				}

				ctx.Send(msg)
				rec := ctx.Send(setu)
				go func() {
					time.Sleep(30 * time.Second)
					ctx.DeleteMessage(rec)
				}()

				ctx.Send("看完不来点感想么0w0")
				next := ctx.FutureEvent("message", ctx.CheckSession())
				recv, cancel := next.Repeat()
				for i := range recv {
					msg := i.MessageString()
					repo := RushedThinking(msg)
					if repo != "" {
						ctx.Send(repo)
					}
					cancel()
				}
			}()
		})

	_ = setu.OnRegex("来[张点丶份](.*?)的[涩色🐍]图", "根据提供的tag查找涩图, 冷却2分钟").
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				m := ctx.State["regex_matched"].([]string)
				if len(m) == 0 {
					return
				}

				tag := m[1]
				msg, setu, err := GetSetu(tag)
				if err != nil {
					ctx.Send("冲不起惹...")
					return
				}

				ctx.Send(msg)
				rec := ctx.Send(setu)
				go func() {
					time.Sleep(30 * time.Second)
					ctx.DeleteMessage(rec)
				}()

				ctx.Send("看完不来点感想么0w0")
				next := ctx.FutureEvent("message", ctx.CheckSession())
				recv, cancel := next.Repeat()
				for i := range recv {
					msg := i.MessageString()
					repo := RushedThinking(msg)
					if repo != "" {
						ctx.Send(repo)
					}
					cancel()
				}
			}()
		})
}
