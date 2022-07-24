package service

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/Kyomotoi/go-ATRI/lib"
	log "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
)

type D map[string]string

const manageDIR = "data/plugins/manage/"

func checkMaDIR() {
	_, err := os.Stat(manageDIR)
	if err != nil {
		if !os.IsExist(err) {
			err := os.MkdirAll(manageDIR, 0777)
			if err != nil {
				log.Warning("目录 " + manageDIR + " 创建失败，请尝试手动创建")
			}
		}
	}
}

func loadBlockList(typ string) D {
	var d D
	var aimFile string

	checkMaDIR()

	if typ == "user" {
		aimFile = "block_user.json"
	} else {
		aimFile = "block_group.json"
	}

	filePath := manageDIR + aimFile
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warning("封禁名单不存在，即将创建")
			t, _ := json.Marshal(d)
			err = os.WriteFile(filePath, t, 0777)
			if err != nil {
				log.Warning("封禁名单写入默认参数失败，请检查是否给足权限")
			}
			return d
		}
	}
	_ = json.Unmarshal(data, &d)
	return d
}

func CheckBlock(ctx *zero.Ctx) bool {
	evTyp := ctx.Event.MessageType
	if evTyp == "user" {
		data := loadBlockList("user")
		userID := strconv.FormatInt(ctx.Event.UserID, 10)
		_, isOK := data[userID]
		if isOK {
			return false
		}
	}

	if evTyp == "group" {
		data := loadBlockList("group")
		groupID := strconv.FormatInt(ctx.Event.GroupID, 10)
		_, isOK := data[groupID]
		if isOK {
			return false
		}
	}

	return true
}

func IsServiceEnabled(serv string) zero.Rule {
	return func(ctx *zero.Ctx) bool {
		data := LoadServiceData(serv)

		userID := strconv.FormatInt(ctx.Event.UserID, 10)
		servBlockUserList := data.DisableUser
		if lib.StringInArray(userID, servBlockUserList) {
			return false
		}

		groupID := strconv.FormatInt(ctx.Event.GroupID, 10)
		servBlockGroupList := data.DisableGroup
		if lib.StringInArray(groupID, servBlockGroupList) {
			return false
		}

		return data.Enabled
	}
}
