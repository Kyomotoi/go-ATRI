package service

import (
	"encoding/json"
	"os"

	"github.com/Kyomotoi/go-ATRI/lib"
	log "github.com/sirupsen/logrus"
)

const serviceDIR = "data/services/"

func init() {
	exi := lib.IsDir(serviceDIR)
	if !exi {
		err := os.MkdirAll(serviceDIR, 0777)
		if err != nil {
			panic("service: 创建文件夹: " + serviceDIR + " 失败, 请尝试手动创建")
		}
	}
}

func generateServiceConfig(service string, docs string) {
	tData := &ServiceInfo{
		Service:     service,
		Docs:        docs,
		CommandList: make(map[string]CommandInfo),
		Enabled:     true,
	}
	data, _ := json.Marshal(tData)

	filePath := serviceDIR + service + ".json"
	err := os.WriteFile(filePath, data, 0777)
	if err != nil {
		log.Error("Write service info failed!")
	}
}

func LoadCommandList(service string) map[string]CommandInfo {
	filePath := serviceDIR + service + ".json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			generateServiceConfig(service, "nothing")
		}
		data, _ = os.ReadFile(filePath)
	}
	var si ServiceInfo
	_ = json.Unmarshal(data, &si)
	return si.CommandList
}

func StoneCommandList(service string, cmds map[string]CommandInfo) {
	sd := LoadServiceData(service)
	sd.CommandList = cmds
	StoreServiceData(service, sd)
}

func LoadServiceData(service string) ServiceInfo {
	var s ServiceInfo

	filePath := serviceDIR + service + ".json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			panic("无法读取服务 " + service + " 中的信息，请重启以尝试修复该错误")
		}
	}
	_ = json.Unmarshal(data, &s)
	return s
}

func StoreServiceData(serv string, d ServiceInfo) {
	filePath := serviceDIR + serv + ".json"
	if !lib.IsExists(filePath) {
		panic("无法读取服务 " + serv + " 中的信息，请重启以尝试修复该错误")
	}

	data, _ := json.Marshal(d)
	err := os.WriteFile(filePath, data, 0777)
	if err != nil {
		panic("写入服务 " + serv + " 时失败，请检查是否给予程序足够的权限")
	}
}
