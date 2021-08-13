package setu

import (
	log "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/extension/single"
	"strconv"
	"time"
)

var limit = rate.NewManager(time.Minute*1, 1)

func init() {
	engine := zero.New()

	single.New(
		single.WithKeyFn(func(ctx *zero.Ctx) interface{} {
			return ctx.Event.UserID
		}),
		single.WithPostFn(func(ctx *zero.Ctx) {
			log.Info("Setu limited user: "+strconv.FormatInt(ctx.Event.UserID, 10))
		}),
	).Apply(engine)

	_ = engine.OnCommandGroup([]string{"来张涩图", "来份涩图", "来点涩图", "涩图来"}).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
				err := RandomSetu(ctx)
				if err != nil {
					return
				}
		})

	_ = engine.OnRegex("来[张点丶份](.*?)的[涩色🐍]图").
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
				tag, _ := ctx.State["regex_matched"].([]string)
				err := TagSetu(tag[0], ctx)
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
