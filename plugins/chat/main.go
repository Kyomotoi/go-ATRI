package chat

import (
	"github.com/Kyomotoi/go-ATRI/service"
	"github.com/Kyomotoi/go-ATRI/utils"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension"
	"strconv"
)

func init() {
	cmds := make(map[string]string)
	cmds["@（内容）"] = "闲聊（文爱"
	cmds["@ 叫我"] = "更改闲聊（划掉 文爱）时的称呼\n此外触发的命令还有：我叫、我是"
	cmds["@ 复读"] = "复读姬！\n此外触发的命令还有：说"
	cmds["@ 更新词库"] = "仅供咱的master触发~"
	service.RegisterService("闲聊", "有点涩？", cmds)

	zero.OnMessage(zero.OnlyToMe, service.CheckBlock).SetPriority(5).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				if !service.IsServiceEnabled("闲聊", ctx) {
					return
				}

				msg := ctx.Event.Message.String()
				userID := strconv.FormatInt(ctx.Event.UserID, 10)
				repo, err := Kimo(msg, userID)
				if err != nil {
					return
				}
				ctx.Send(repo)
			}()
		})

	zero.OnCommandGroup([]string{"叫我", "我叫", "我是"}, zero.OnlyToMe, service.CheckBlock).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				if !service.IsServiceEnabled("闲聊", ctx) {
					return
				}

				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
				}

				if cmd.Args == "" {
					ctx.Send("欧尼酱想让咱如何称呼呢！0w0")
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
				ctx.Send("好欸！" + cmd.Args + "ちゃん~~~")
			}()
		})

	zero.OnCommandGroup([]string{"复读", "说"}, zero.OnlyToMe, service.CheckBlock).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				if !service.IsServiceEnabled("闲聊", ctx) {
					return
				}

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
}
