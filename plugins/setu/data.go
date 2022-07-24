package setu

import (
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	setuURL = "https://api.lolicon.app/setu/v2"

	rt_nice_patt  = "[hH好][sS涩色][oO哦]|[嗯恩摁社蛇🐍射]了|(硬|石更)了|[牛🐂][牛🐂]要炸了|[炼恋]起来|开?导"
	rt_nope_patt  = "不够[涩色]|就这|不行|不彳亍|一般|这也[是叫算]|[?？]|就这|爬|爪巴"
	rt_again_patt = "再来一张|不够|还要"
)

var (
	rt_nice_repo  = []string{"w", "好诶！", "ohh", "(///w///)", "🥵", "我也"}
	rt_nope_repo  = []string{"那你来发", "爱看不看", "你看不看吧", "看这种类型的涩图，是一件多么美妙的事情", "急了"}
	rt_again_repo = []string{"没了...", "自己找去"}
)

func request(tag string) (gjson.Result, error) {
	var u string

	if tag != "" {
		u = setuURL + "?tag=" + url.QueryEscape(tag)
	} else {
		u = setuURL
	}

	resp, err := http.Get(u)
	if err != nil && resp.StatusCode != http.StatusOK {
		log.Warning("setu: 在请求链接 " + u + "时发生错误. Response msg: " + resp.Status)
		return gjson.Result{}, err
	}

	data, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	j := gjson.ParseBytes(data)
	isOK := j.Get("error").String()
	if isOK != "" {
		log.Warning("setu: 接口返回出现错误，内容: \n" + isOK)
		return gjson.Result{}, err
	}

	return j, nil
}

func GetSetu(tag string) (string, message.MessageSegment, error) {
	data, err := request(tag)
	if err != nil {
		return "", message.MessageSegment{}, err
	}

	title := data.Get("data.0.title").String()
	pid := data.Get("data.0.pid").String()
	raw_img := data.Get("data.0.urls.original").String()

	img := strings.Replace(raw_img, "i.pixiv.cat", "i.pixiv.re", 1)

	msg := "Title: " + title + "\nPid: " + pid
	setu := message.Image(img)
	return msg, setu, nil
}

func RushedThinking(think string) string {
	nice_regex := regexp.MustCompile(rt_nice_patt)
	if matched := nice_regex.FindStringSubmatch(think); matched != nil {
		return rt_nice_repo[rand.Intn(len(rt_nice_repo))]
	}

	nope_regex := regexp.MustCompile(rt_nope_patt)
	if matched := nope_regex.FindStringSubmatch(think); matched != nil {
		return rt_nope_repo[rand.Intn(len(rt_nope_repo))]
	}

	again_regex := regexp.MustCompile(rt_again_patt)
	if matched := again_regex.FindStringSubmatch(think); matched != nil {
		return rt_again_repo[rand.Intn(len(rt_again_repo))]
	}

	return ""
}
