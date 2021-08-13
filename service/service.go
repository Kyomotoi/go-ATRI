package service

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func checkDIR() error {
	var fileDIR string
	fileDIR = "data/service"
	_, err := os.Stat(fileDIR)
	if err != nil {
		if !os.IsExist(err) {
			err := os.MkdirAll(fileDIR, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func generateServiceConfig(service string, docs string) error {
	err := checkDIR()
	if err != nil {
		log.Error("Failed to create service folder.")
		return err
	}

	tData := &Service{
		Service:      service,
		Docs:         docs,
		Enabled:      false,
		OnlyAdmin:    false,
		DisableUser:  []string{},
		DisableGroup: []string{},
	}
	data, err := json.Marshal(tData)
	if err != nil {
		return err
	}

	filePath := "data/service/" + service + ".json"
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
