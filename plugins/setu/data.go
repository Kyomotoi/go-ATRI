package setu

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const SetuURL = "https://api.lolicon.app/setu/v2"

func request(tag string) (gjson.Result, error) {
	var u string

	if tag != "" {
		u = SetuURL + "?tag=" + url.QueryEscape(tag)
	} else {
		u = SetuURL
	}
	resp, err := http.Get(u)
	if err != nil {
		return gjson.Result{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return gjson.Result{}, err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return gjson.Result{}, err
	}

	j := gjson.ParseBytes(data)
	isOK := j.Get("error").String()
	if isOK != "" {
		log.Error("Lolicon API exists error: "+isOK)
		return gjson.Result{}, err
	}
	return j, nil
}

func RandomSetu(ctx *zero.Ctx) error {
	j, err := request("")
	if err != nil {
		log.Error("Failed to request url: "+SetuURL)
		return err
	}

	title := j.Get("data.0.title").String()
	pid := j.Get("data.0.pid").String()
	img := j.Get("data.0.urls.original").String()

	repo := "Title: " + title + "\nPid: " + pid
	ctx.Send(repo)
	setu := ctx.Send(message.Image(img))
	time.Sleep(30*time.Second)
	ctx.DeleteMessage(setu)
	log.Info("Recall setu: "+img)
	return nil
}

func TagSetu(tag string, ctx *zero.Ctx) error {
	j, err := request(tag)
	if err != nil {
		log.Error("Failed to request url: "+SetuURL)
		return err
	}
	title := j.Get("data.0.title").String()
	pid := j.Get("data.0.pid").String()
	img := j.Get("data.0.urls.original").String()

	repo := "Title: " + title + "\nPid: " + pid
	ctx.Send(repo)
	setu := ctx.Send(message.Image(img))
	time.Sleep(30*time.Second)
	ctx.DeleteMessage(setu)
	log.Info("Recall setu: "+img)
	return nil
}
