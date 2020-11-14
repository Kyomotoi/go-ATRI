package models

import (
	"io/ioutil"
	"net/http"
)

func HttpGET(link string) string {

	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(msg)
}