package gocqhttp

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Kyomotoi/go-ATRI/lib"
	log "github.com/sirupsen/logrus"
)

const gocqReleasesList = "https://api.github.com/repos/Mrs4s/go-cqhttp/releases"

func getGoCQHTTPReleaseInfo() (GithubReleaseList, error) {
	var model GithubReleaseList

	resp, err := http.Get(gocqReleasesList)
	if err != nil && resp.StatusCode != http.StatusOK {
		log.Warn("Driver: 在请求链接 " + gocqReleasesList + "时发生错误. Response msg: " + resp.Status)
		log.Warn("Driver: 内置 gocqhttp 将失效")
		return model, err
	}

	data, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	_ = json.Unmarshal(data, &model)
	return model, nil
}

func getFileNameOfGoCQHTTP() string {
	var a, b, p string

	a, b = lib.GetPlatformType()
	if a == "windows" {
		p = ".exe"
	} else {
		p = ".tar.gz"
	}

	return "go-cqhttp_" + a + "_" + b + p
}

func downloadGoCQHTTP(v string) error {
	rli, err := getGoCQHTTPReleaseInfo()
	if err != nil {
		return err
	}

	if rli == nil {
		log.Warn("Driver: 没有可用的 gocqhttp 发行版")
		return errors.New("driver: no gocqhttp version availabled")
	}

	var downloadURL string

	if v == "" || v == "latest" {
		latestRelease := rli[0]
		for _, item := range latestRelease.Assets {
			if item.Name == getFileNameOfGoCQHTTP() {
				downloadURL = item.BrowserDownloadURL
				break
			}
		}
	} else {
		for _, item := range rli {
			if item.TagName == v {
				latestRelease := item
				for _, item := range latestRelease.Assets {
					if item.Name == getFileNameOfGoCQHTTP() {
						downloadURL = item.BrowserDownloadURL
						break
					}
				}
				break
			}
		}
	}

	if downloadURL == "" {
		latestRelease := rli[0]
		for _, item := range latestRelease.Assets {
			if item.Name == getFileNameOfGoCQHTTP() {
				downloadURL = item.BrowserDownloadURL
				break
			}
		}
	}

	filePath := resourceDIR + "/" + v + "/" + getFileNameOfGoCQHTTP()
	err = lib.DownloadFile(filePath, downloadURL)
	if err != nil {
		log.Warn("Driver: 下载 gocqhttp 时发生错误. Error: " + err.Error())
		return err
	}

	return nil
}
