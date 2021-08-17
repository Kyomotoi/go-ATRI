package setu

import (
	"github.com/Kyomotoi/go-ATRI/service"
	log "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/extension/single"
	"strconv"
	"time"
)

var limit = rate.NewManager(time.Minute*2, 1)

func init() {
	cmds := make(map[string]string)
	cmds["来张涩图"] = "随机涩图！\n除此之外触发命令还有：来（份，点）涩图、涩图来。限制两分钟一张"
	cmds["来[张点丶份](.*?)的?[涩色🐍]图"] = "根据提供的tag查找涩图。限制两分钟一张"
	service.RegisterService("涩图", "hso!", cmds)

	engine := zero.New()

	single.New(
		single.WithKeyFn(func(ctx *zero.Ctx) interface{} {
			return ctx.Event.UserID
		}),
		single.WithPostFn(func(ctx *zero.Ctx) {
			log.Info("处于涩图限制的用户："+strconv.FormatInt(ctx.Event.UserID, 10))
		}),
	).Apply(engine)

	_ = engine.OnCommandGroup([]string{"来张涩图", "来份涩图", "来点涩图", "涩图来"}, service.CheckBlock).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if !service.IsServiceEnabled("涩图", ctx) {
				return
			}
			err := RandomSetu(ctx)
			if err != nil {
				return
			}
		})

	_ = engine.OnRegex("来[张点丶份](.*?)的[涩色🐍]图", service.CheckBlock).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if !service.IsServiceEnabled("涩图", ctx) {
				return
			}
			tag, _ := ctx.State["regex_matched"].([]string)
			err := TagSetu(tag[1], ctx)
			if err != nil {
				return
			}
		})

	engine.UsePreHandler(func(ctx *zero.Ctx) bool {
		if !limit.Load(ctx.Event.UserID).Acquire() {
			return false
		}
		return true
	})
}
