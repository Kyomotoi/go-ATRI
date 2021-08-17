package help

import zero "github.com/wdvxdr1123/ZeroBot"

func init() {
	zero.OnCommandGroup([]string{"help", "menu", "菜单"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				ctx.Send(Menu())
			}()
		})

	zero.OnCommandGroup([]string{"about", "关于", "你是谁"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				ctx.Send(About())
			}()
		})

	zero.OnCommandGroup([]string{"服务列表", "功能列表"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				ctx.Send(ServiceList())
			}()
		})

	zero.OnRegex("帮助\\s(.*?)\\s(.*?)$", zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				defer func() {
					if err := recover(); err != nil {
						ctx.Send(err)
					}
					return
				}()

				m := ctx.State["regex_matched"].([]string)

				servName := m[1]
				servCmd := m[2]

				if servCmd != "" {
					ctx.Send(CommandInfo(servName, servCmd))
				} else {
					ctx.Send(ServiceInfo(servName))
				}
			}()
		})
}
