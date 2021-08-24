package service

import (
	"encoding/json"
	"github.com/Kyomotoi/go-ATRI/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

const FileDIR = "data/service/"

type cmds map[string]string

func init() {
	utils.InitLogger()
}

func checkDIR() error {
	_, err := os.Stat(FileDIR)
	if err != nil {
		if !os.IsExist(err) {
			err := os.MkdirAll(FileDIR, 0777)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func generateServiceConfig(service string, docs string, commands cmds) error {
	err := checkDIR()
	if err != nil {
		log.Error("Failed to create service folder.")
		return err
	}

	tData := &Service{
		Service:      service,
		Docs:         docs,
		Commands:     commands,
		Enabled:      true,
		DisableUser:  []string{},
		DisableGroup: []string{},
	}
	data, err := json.Marshal(tData)
	if err != nil {
		return err
	}

	filePath := FileDIR + service + ".json"
	err = ioutil.WriteFile(filePath, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

func RegisterService(service string, docs string, commands cmds) {
	filePath := FileDIR + service + ".json"
	if !utils.IsExists(filePath) {
		err := generateServiceConfig(service, docs, commands)
		if err != nil {
			log.Warning("服务 " + service + " 注册失败，将强制禁用，并无法启用")
			return
		}
	}

	itemCmds := len(commands)

	log.Printf("成功加载服务：%s，从中读到 %d 个触发器", service, itemCmds)
}

func LoadServiceData(service string) Service {
	var j Service

	filePath := FileDIR + service + ".json"
	if !utils.IsExists(filePath) {
		panic("无法读取服务 " + service + " 中的信息，请重启以尝试修复该错误")
	}
	data, _ := os.ReadFile(filePath)
	_ = json.Unmarshal(data, &j)
	return j
}

func StoreServiceData(serv string, d Service) {
	filePath := FileDIR + serv + ".json"
	if !utils.IsExists(filePath) {
		panic("无法读取服务 " + serv + " 中的信息，请重启以尝试修复该错误")
	}

	data, _ := json.Marshal(d)
	err := ioutil.WriteFile(filePath, data, 0777)
	if err != nil {
		panic("写入服务 " + serv + " 时失败，请检查是否给予程序足够的权限")
	}
}
