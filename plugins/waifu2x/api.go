package waifu2x

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func waifu2x(ImageUrl string, Key string) (string, error) {
	rjson := make(map[string]string)
	client := &http.Client{}
	data := url.Values{"image": {ImageUrl}}
	req, _ := http.NewRequest("POST", "https://api.deepai.org/api/waifu2x", strings.NewReader(data.Encode()))
	req.Header.Add("api-key", Key)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, doErr := client.Do(req)
	if doErr != nil {
		return "", doErr
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	unmErr := json.Unmarshal(body, &rjson)
	if unmErr != nil {
		return "", unmErr
	}
	if rjson["err"] != "" {
		return "", errors.New(rjson["err"])
	}
	return rjson["output_url"], nil
}
