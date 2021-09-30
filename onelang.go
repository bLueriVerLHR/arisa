package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
)

func OneSentence(typeName string) OneLang {
	if regexp.MustCompile(`动画`).Match([]byte(typeName)) {
		typeName = "a"
	} else if regexp.MustCompile(`漫画`).Match([]byte(typeName)) {
		typeName = "b"
	} else if regexp.MustCompile(`游戏`).Match([]byte(typeName)) {
		typeName = "c"
	} else if regexp.MustCompile(`文学`).Match([]byte(typeName)) {
		typeName = "d"
	} else if regexp.MustCompile(`原创`).Match([]byte(typeName)) {
		typeName = "e"
	} else if regexp.MustCompile(`网络`).Match([]byte(typeName)) {
		typeName = "f"
	} else if regexp.MustCompile(`随便`).Match([]byte(typeName)) {
		typeName = "g"
	} else if regexp.MustCompile(`影视|电视|电影`).Match([]byte(typeName)) {
		typeName = "h"
	} else if regexp.MustCompile(`诗词`).Match([]byte(typeName)) {
		typeName = "i"
	} else if regexp.MustCompile(`网易云|网抑云`).Match([]byte(typeName)) {
		typeName = "j"
	} else if regexp.MustCompile(`哲学|哲思`).Match([]byte(typeName)) {
		typeName = "k"
	} else if regexp.MustCompile(`抖机灵`).Match([]byte(typeName)) {
		typeName = "l"
	} else {
		typeName = "d"
	}
	response, err := http.Get("https://v1.hitokoto.cn/?c=" + typeName)
	CheckError(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	var ol OneLang
	err = json.Unmarshal(body, &ol)
	CheckError(err)
	return ol
}
