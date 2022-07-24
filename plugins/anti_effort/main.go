package antieffort

import (
	"github.com/Kyomotoi/go-ATRI/service"
	"github.com/Kyomotoi/go-ATRI/utils"
	log "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
)

const (
	getURLmsg = `请键入wakatime share embed URL:
获取方法:
  - 前往 wakatime.com/share/embed
  - Format 选择 JSON
  - Chart Type 选择 Coding Activity
  - Date Range 选择 Last 7 Days
  - 所需url就在下一栏 HTML 中的 url`
	serviceName = "谁是卷王"
)

var joinAttitude = []string{"y", "Y", "是", "希望", "同意"}

func init() {
	antiEffort := service.NewService(serviceName, " 谁是卷王！", false, "/ae", service.CheckBlock, service.IsServiceEnabled(serviceName), zero.OnlyGroup)

	_ = antiEffort.OnCommand("111我也要卷", "加入卷王统计榜", []string{}).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			go func() {
				var url string
				var nickname string
				var isJoinGlobal string

				ctx.Send(getURLmsg)
				next := ctx.FutureEvent("message", ctx.CheckSession())
				recv, cancel := next.Repeat()
				for i := range recv {
					msg := i.MessageString()
					url = msg
					cancel()
				}

				ctx.Send("如何在排行榜中称呼你捏")
				next = ctx.FutureEvent("message", ctx.CheckSession())
				recv, cancel = next.Repeat()
				for i := range recv {
					msg := i.MessageString()
					nickname = msg
					cancel()
				}

				ctx.Send("你希望加入公共排行榜吗？(y/n)")
				next = ctx.FutureEvent("message", ctx.CheckSession())
				recv, cancel = next.Repeat()
				for i := range recv {
					msg := i.MessageString()
					isJoinGlobal = msg
					cancel()
				}

				result, err := addUser(ctx.Event.GroupID, ctx.Event.UserID, nickname, url)
				if err != nil {
					ctx.Send("操作失败惹...")
					return
				}
				if utils.StringInArray(isJoinGlobal, joinAttitude) {
					addUser(0, ctx.Event.UserID, nickname, url)
				}
				ctx.Send(result)
			}()
		})

	_ = antiEffort.OnCommand("111参加公共卷王", "加入公共卷王统计榜", []string{}).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			rawData := getData(ctx.Event.GroupID)
			if len(rawData.Data) != 0 {
				data := rawData.Data
				for _, v := range data {
					if v.UserID == ctx.Event.UserID {
						nickname := v.UserNickname
						url := v.WakaURL
						addUser(0, ctx.Event.UserID, nickname, url)
						ctx.Send("完成~！")
						return
					}
				}
			}

			var url string
			var nickname string

			ctx.Send(getURLmsg)
			next_0 := ctx.FutureEvent("message", ctx.CheckSession())
			recv_0, cancel := next_0.Repeat()
			for i := range recv_0 {
				msg := i.MessageString()
				url = msg
				cancel()
			}

			ctx.Send("如何在排行榜中称呼你捏")
			next_1 := ctx.FutureEvent("message", ctx.CheckSession())
			recv_1, cancel := next_1.Repeat()
			for i := range recv_1 {
				msg := i.MessageString()
				nickname = msg
				cancel()
			}

			result, err := addUser(ctx.Event.GroupID, ctx.Event.UserID, nickname, url)
			if err != nil {
				ctx.Send("操作失败惹...")
				return
			}
			_, err = addUser(0, ctx.Event.UserID, nickname, url)
			if err != nil {
				ctx.Send("操作失败惹...")
				return
			}
			ctx.Send(result)
		})

	_ = antiEffort.OnCommand("111我不卷了", "退出卷王统计榜", []string{}).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			result := delUser(ctx.Event.GroupID, ctx.Event.UserID)
			delUser(0, ctx.Event.UserID)
			ctx.Send(result)
		})

	_ = antiEffort.OnCommand("今日卷王", "查看今日卷王榜", []string{}).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send("别急！正在统计！")
			rawData := getData(ctx.Event.GroupID)
			if len(rawData.Data) == 0 {
				ctx.Send("暂无数据！")
				return
			}

			result := genRank(rawData, ctx.Event.UserID, "today")
			ctx.Send(result)
		})

	_ = antiEffort.OnCommand("周卷王", "查看最近七天卷王榜", []string{}).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send("别急！正在统计！")
			rawData := getData(ctx.Event.GroupID)
			if len(rawData.Data) == 0 {
				ctx.Send("暂无数据！")
				return
			}

			result := genRank(rawData, ctx.Event.UserID, "week")
			ctx.Send(result)
		})

	_ = antiEffort.OnCommand("公共卷王", "查看今日公共卷王榜", []string{}).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send("别急！正在统计！")
			rawData := getData(0)
			if len(rawData.Data) == 0 {
				ctx.Send("暂无数据！")
				return
			}

			result := genRank(rawData, ctx.Event.UserID, "global_today")
			ctx.Send(result)
		})

	_ = antiEffort.OnCommand("公共周卷王", "查看最近七天公共卷王榜", []string{}).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send("别急！正在统计！")
			rawData := getData(0)
			if len(rawData.Data) == 0 {
				ctx.Send("暂无数据！")
				return
			}

			result := genRank(rawData, ctx.Event.UserID, "global_week")
			ctx.Send(result)
		})

	_ = antiEffort.OnCommand("更新卷王榜", "更新卷王榜", []string{}).
		SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			updateData()
			ctx.Send("完成~！")
		})

	scheduler := utils.Scheduler

	_ = scheduler.AddFunc("* 10 * * * ? ", func() {
		log.Debug("anti_effort: 更新卷王们的数据")
		updateData()
	})
}
