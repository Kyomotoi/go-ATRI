package manage

import (
	"encoding/json"
	"github.com/Kyomotoi/go-ATRI/service"
	"github.com/Kyomotoi/go-ATRI/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

const fileDIR = "data/database/manage/"

type D map[string]string

func checkDIR() {
	_, err := os.Stat(fileDIR)
	if err != nil {
		if !os.IsExist(err) {
			err := os.MkdirAll(fileDIR, 0777)
			if err != nil {
				log.Warning("目录 "+fileDIR+" 创建失败，请尝试手动创建")
			}
		}
	}
}

func loadBlockUser() D {
	var d D

	checkDIR()
	filePath := fileDIR + "blockUser.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		if !utils.IsExists(filePath) {
			err = ioutil.WriteFile(filePath, []byte("{}"), 0777)
			if err != nil {
				log.Warning("读取用户封禁名单失败，将返回空名单")
				return d
			}
		}
	}
	err = json.Unmarshal(data, &d)
	return d
}

func storeBlockUser(data D) error {
	_ = loadBlockUser()

	checkDIR()
	d, err := json.Marshal(data)
	if err != nil {
		log.Warning("解析用户封禁名单失败，用户封禁检测系统或许不再有效")
		return err
	}

	filePath := fileDIR + "blockUser.json"
	err = ioutil.WriteFile(filePath, d, 0777)
	if err != nil {
		log.Warning("用户封禁名单写入失败，用户封禁检测系统或许不再有效")
		return err
	}
	return nil
}

func loadBlockGroup() D {
	var d D

	checkDIR()
	filePath := fileDIR + "blockGroup.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		if !utils.IsExists(filePath) {
			t, _ := json.Marshal(d)
			err = ioutil.WriteFile(filePath, t, 0777)
			if err != nil {
				log.Warning("读取群封禁名单失败，将返回空名单")
				return d
			}
		}
	}
	err = json.Unmarshal(data, &d)
	return d
}

func storeBlockGroup(data D) error {
	_ = loadBlockGroup()

	checkDIR()
	d, err := json.Marshal(data)
	if err != nil {
		log.Warning("解析群封禁名单失败，群封禁检测系统或许不再有效")
		return err
	}

	filePath := fileDIR + "blockGroup.json"
	err = ioutil.WriteFile(filePath, d, 0777)
	if err != nil {
		log.Warning("群封禁名单写入失败，群封禁系统或许不再有效")
		return err
	}
	return nil
}

func BlockUser(userID string) error {
	data := loadBlockUser()
	data[userID] = time.Now().Format("2006-01-02 15:04:05")

	err := storeBlockUser(data)
	if err != nil {
		return err
	}
	return nil
}

func UnBlockUser(userID string) error {
	data := loadBlockUser()
	delete(data, userID)

	err := storeBlockUser(data)
	if err != nil {
		return err
	}
	return nil
}

func BlockGroup(groupID string) error {
	data := loadBlockGroup()
	data[groupID] = time.Now().Format("2006-01-02 15:04:05")

	err := storeBlockUser(data)
	if err != nil {
		return err
	}
	return nil
}

func UnBlockGroup(groupID string) error {
	data := loadBlockGroup()
	delete(data, groupID)

	err := storeBlockGroup(data)
	if err != nil {
		return err
	}
	return nil
}

func ControlGlobalService(serv string, isEnabled bool) {
	data := service.LoadServiceData(serv)
	data.Enabled = isEnabled
	service.StoreServiceData(serv, data)
}

func ControlGroupService(serv string, groupID string, isEnabled bool) {
	var t []string

	data := service.LoadServiceData(serv)
	if isEnabled {
		t = utils.DeleteAimInArray(groupID, data.DisableUser)
	} else {
		t = append(data.DisableGroup, groupID)
	}
	data.DisableGroup = t
	service.StoreServiceData(serv, data)
}


func ControlUserService(serv string, userID string, isEnabled bool) {
	var t []string

	data := service.LoadServiceData(serv)
	if isEnabled {
		t = utils.DeleteAimInArray(userID, data.DisableUser)
	} else {
		t = append(data.DisableUser, userID)
	}
	data.DisableUser = t
	service.StoreServiceData(serv, data)
}
