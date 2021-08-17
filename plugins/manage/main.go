package manage

import (
	"github.com/Kyomotoi/go-ATRI/service"
	"github.com/Kyomotoi/go-ATRI/utils"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension"
	"strconv"
)

func init() {
	cmds := make(map[string]string)
	cmds["封禁用户 qid"] = "封禁用户，仅供维护者触发~"
	cmds["解封用户 qid"] = "解封用户，仅供维护者触发~"
	cmds["封禁群 gid"] = "封禁群，仅供维护者触发~"
	cmds["解封群 gid"] = "解封群，仅供维护者触发~"
	cmds["全局启用 service"] = "全局启用服务，仅供维护者触发~"
	cmds["全局禁用 service"] = "全局禁用服务，仅供维护者触发~"
	cmds["启用 service"] = "针对所在群启用服务，维护者及群管理可触发"
	cmds["禁用 service"] = "针对所在群禁用服务，维护者及群管理可触发"
	cmds["对用户(.*?)(禁用|启用)(.*?)$"] = "针对单个用户启用，禁用服务\nExample: 对用户114514禁用涩图"
	service.RegisterService("管理", "控制bot的各项服务", cmds)

	zero.OnCommand("封禁用户", zero.SuperUserPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
					return
				}

				if cmd.Args == "" {
					ctx.Send("哪位？GKD！")
					next := ctx.FutureEvent("message", ctx.CheckSession())
					recv, cancel := next.Repeat()
					for i := range recv {
						msg := i.Message.String()
						if utils.StringInArray(msg, []string{"算了", "罢了", "取消"}) {
							ctx.Send("...看来有人逃过一劫呢")
						} else {
							if msg != "" {
								cmd.Args = msg
								cancel()
								continue
							}
						}
					}
				}

				err = BlockUser(cmd.Args)
				if err != nil {
					ctx.Send("封禁失败了呢...")
				}
				ctx.Send("用户 "+cmd.Args+" 危！")
			}()
		})

	zero.OnCommand("解封用户", zero.SuperUserPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
					return
				}

				if cmd.Args == "" {
					ctx.Send("哪位？GKD！")
					next := ctx.FutureEvent("message", ctx.CheckSession())
					recv, cancel := next.Repeat()
					for i := range recv {
						msg := i.Message.String()
						if utils.StringInArray(msg, []string{"算了", "罢了", "取消"}) {
							ctx.Send("...有人又得继续在小黑屋呆一阵子了")
						} else {
							if msg != "" {
								cmd.Args = msg
								cancel()
								continue
							}
						}
					}
				}

				err = UnBlockUser(cmd.Args)
				if err != nil {
					ctx.Send("解封失败了呢...")
				}
				ctx.Send("用户 "+cmd.Args + " 重获新生！")
			}()
		})

	zero.OnCommand("封禁群", zero.SuperUserPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
					return
				}

				if cmd.Args == "" {
					ctx.Send("哪个群？GKD！")
					next := ctx.FutureEvent("message", ctx.CheckSession())
					recv, cancel := next.Repeat()
					for i := range recv {
						msg := i.Message.String()
						if utils.StringInArray(msg, []string{"算了", "罢了", "取消"}) {
							ctx.Send("...看来有一群逃过一劫呢")
						} else {
							if msg != "" {
								cmd.Args = msg
								cancel()
								continue
							}
						}
					}
				}

				err = BlockGroup(cmd.Args)
				if err != nil {
					ctx.Send("封禁失败了呢...")
				}
				ctx.Send("群 "+cmd.Args + " 危！")
			}()
		})

	zero.OnCommand("解封群", zero.SuperUserPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) 	{
			go func() {
				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
					return
				}

				if cmd.Args == "" {
					ctx.Send("哪个群？GKD！")
					next := ctx.FutureEvent("message", ctx.CheckSession())
					recv, cancel := next.Repeat()
					for i := range recv {
						msg := i.Message.String()
						if utils.StringInArray(msg, []string{"算了", "罢了", "取消"}) {
							ctx.Send("...有一群又得继续在小黑屋呆一阵子了")
						} else {
							if msg != "" {
								cmd.Args = msg
								cancel()
								continue
							}
						}
					}
				}

				err = UnBlockGroup(cmd.Args)
				if err != nil {
					ctx.Send("解封失败了呢...")
				}
				ctx.Send("群 "+cmd.Args+" 重获新生！")
			}()
		})

	zero.OnCommand("全局启用", zero.SuperUserPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
					return
				}

				if cmd.Args == "" {
					ctx.Send("目标服务呢~？")
					next := ctx.FutureEvent("message", ctx.CheckSession())
					recv, cancel := next.Repeat()
					for i := range recv {
						msg := i.Message.String()
						if utils.StringInArray(msg, []string{"算了", "罢了", "取消"}) {
							ctx.Send("...好吧")
						} else {
							if msg != "" {
								cmd.Args = msg
								cancel()
								continue
							}
						}
					}
				}

				defer func() {
					if err := recover(); err != nil {
						ctx.Send(err)
					}
					return
				}()

				ControlGlobalService(cmd.Args, true)
				ctx.Send("完成~！服务 "+cmd.Args+" 已全局启用")
			}()
		})

	zero.OnCommand("全局禁用", zero.SuperUserPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
					return
				}

				if cmd.Args == "" {
					ctx.Send("目标服务呢~？")
					next := ctx.FutureEvent("message", ctx.CheckSession())
					recv, cancel := next.Repeat()
					for i := range recv {
						msg := i.Message.String()
						if utils.StringInArray(msg, []string{"算了", "罢了", "取消"}) {
							ctx.Send("...好吧")
						} else {
							if msg != "" {
								cmd.Args = msg
								cancel()
								continue
							}
						}
					}
				}

				defer func() {
					if err := recover(); err != nil {
						ctx.Send(err)
					}
					return
				}()

				ControlGlobalService(cmd.Args, false)
				ctx.Send("完成~！服务 "+cmd.Args+" 已全局禁用")
			}()
		})

	zero.OnCommand("启用", zero.AdminPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				if ctx.Event.MessageType != "group" {
					ctx.Send("该功能只能在群内使用（")
					return
				}

				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
					return
				}

				if cmd.Args == "" {
					ctx.Send("目标服务呢~？")
					next := ctx.FutureEvent("message", ctx.CheckSession())
					recv, cancel := next.Repeat()
					for i := range recv {
						msg := i.Message.String()
						if utils.StringInArray(msg, []string{"算了", "罢了", "取消"}) {
							ctx.Send("...好吧")
						} else {
							if msg != "" {
								cmd.Args = msg
								cancel()
								continue
							}
						}
					}
				}

				defer func() {
					if err := recover(); err != nil {
						ctx.Send(err)
					}
					return
				}()

				groupID := strconv.FormatInt(ctx.Event.GroupID, 10)
				ControlGroupService(cmd.Args, groupID, true)
				ctx.Send("完成！～已允许本群使用服务："+cmd.Args)
			}()
		})

	zero.OnCommand("禁用", zero.AdminPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				if ctx.Event.MessageType != "group" {
					ctx.Send("该功能只能在群内使用（")
					return
				}

				var cmd extension.CommandModel
				err := ctx.Parse(&cmd)
				if err != nil {
					ctx.Send("...呜呜..出故障了")
					return
				}

				if cmd.Args == "" {
					ctx.Send("目标服务呢~？")
					next := ctx.FutureEvent("message", ctx.CheckSession())
					recv, cancel := next.Repeat()
					for i := range recv {
						msg := i.Message.String()
						if utils.StringInArray(msg, []string{"算了", "罢了", "取消"}) {
							ctx.Send("...好吧")
						} else {
							if msg != "" {
								cmd.Args = msg
								cancel()
								continue
							}
						}
					}
				}

				defer func() {
					if err := recover(); err != nil {
						ctx.Send(err)
					}
					return
				}()

				groupID := strconv.FormatInt(ctx.Event.GroupID, 10)
				ControlGroupService(cmd.Args, groupID, false)
				ctx.Send("完成！～已禁止本群使用服务："+cmd.Args)
			}()
		})

	zero.OnRegex("对用户(.*?)(禁用|启用)(.*?)$", zero.SuperUserPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				m := ctx.State["regex_matched"].([]string)

				aimUserID := m[1]
				isEnabled := m[2]
				aimService := m[3]

				sl := []string{"启用", "禁用"}
				if !utils.StringInArray(isEnabled, sl) {
					ctx.Send("请检查传入参数~！只能选择禁用 或 启用！")
					return
				}

				var isE bool
				if isEnabled == "启用" {
					isE = true
				} else {
					isE = false
				}

				defer func() {
					if err := recover(); err != nil {
						ctx.Send(err)
					}
					return
				}()

				ControlUserService(aimService, aimUserID, isE)
				ctx.Send("用户 "+aimUserID+" 服务 "+aimService+" 已处于"+isEnabled)
			}()
		})
}
