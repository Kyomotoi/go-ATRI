package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type JsonConfig struct {
	Bot *JsonConfigBot `json:"bot"`
	Api *JsonConfigApi `json:"api"`
}

type JsonConfigBot struct {
		Host           string            `json:"host"`
		Port           string            `json:"port"`
	    CommandStart   string            `json:"command_start"`
	    Nickname       []string `json:"nickname"`
		Superusers     []string `json:"superusers"`
	}

type JsonConfigApi struct {
	LoliconAPI     string `json:"lolicon_api"`
	SauceNaoKEY    string `json:"sauce_nao_key"`
	FaceplusAPI    string `json:"faceplus_api"`
	FaceplusSECRET string `json:"faceplus_secret"`
}

func DefaultConfig() *JsonConfig {
	return &JsonConfig{
		Bot: &JsonConfigBot{
				Host:         "127.0.0.1",
				Port:         "8080",
				CommandStart: "",
				Nickname:     []string{},
				Superusers:   []string{},

		},
		Api: &JsonConfigApi{
				LoliconAPI:     "",
				SauceNaoKEY:    "",
				FaceplusAPI:    "",
				FaceplusSECRET: "",
		},
	}
}

func LoadConfig(p string) *JsonConfig {
	filePtr, err := os.Open("config.json")
	if err != nil {
		log.Warnf("加载 config.json 时出错，无法读取到该文件")
		return nil
	}
	defer filePtr.Close()

	c := JsonConfig{}
	err = json.Unmarshal([]byte(p), &c)
	if err != nil {
		log.Warnf("加载 config.json 时出错：%v", p)
		log.Infof("正在备份原文件...")
		os.Rename(p, p+".backup"+strconv.FormatInt(time.Now().Unix(), 10))
		return nil
	}
	return &c
}

func (c *JsonConfig) ConfigSave(p string) error {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(p, []byte(data), 0644)
	return nil
}