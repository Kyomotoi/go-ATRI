package chat

import (
	"github.com/Kyomotoi/go-ATRI/utils"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension"
	"strconv"
)

func init() {
	zero.OnMessage(zero.OnlyToMe).SetPriority(5).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				msg := ctx.Event.Message.String()
				userID := strconv.FormatInt(ctx.Event.UserID, 10)
				repo, err := Kimo(msg,userID)
				if err != nil {
					return
				}
				ctx.Send(repo)
			}()
		})

	zero.OnCommandGroup([]string{"叫我", "我叫", "我是"}, zero.OnlyToMe).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
				}

				if cmd.Args == "" {
					ctx.Send("想让咱如何称呼你呢（")
					next := ctx.FutureEvent("message", ctx.CheckSession())
					recv, cancel := next.Repeat()
					for i := range recv {
						msg := i.Message.String()
						if utils.StringInArray(msg, []string{"CQ"}) {
							ctx.Send("咱不喜欢这个，换一个吧（")
						} else {
							if msg != "" {
								cmd.Args = msg
								cancel()
								continue
							}
						}
					}
				}

				userID := strconv.FormatInt(ctx.Event.UserID, 10)
				err = StoreUserNickname(userID, cmd.Args)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
					return
				}
				ctx.Send("好欸！"+cmd.Args+"ちゃん~~~")
			}()
		})

	zero.OnCommand("更新词库", zero.OnlyToMe, zero.SuperUserPermission).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				err := UpdateData()
				if err != nil {
					ctx.Send("失败了呢...")
					return
				}
				ctx.Send("好欸！更新完成了！")
			}()
		})

	zero.OnCommandGroup([]string{"复读", "说"}, zero.OnlyToMe).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
				}

				if cmd.Args == "" {
					ctx.Send("想让咱复读啥呢？")
					next := ctx.FutureEvent("message", ctx.CheckSession())
					recv, cancel := next.Repeat()
					for i := range recv {
						msg := i.Message.String()
						if msg != "" {
							cmd.Args = msg
							cancel()
							continue
						}
					}
				}

				if utils.StringInArray(cmd.Args, []string{"CQ"}) {
					ctx.Send("咱不想复读这个！")
				} else {
					ctx.Send(cmd.Args)
				}
			}()
		})
}
