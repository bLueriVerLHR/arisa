package plugins

import (
	"arisa/tools"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Setu struct {
	Error string `json:"error"`
	Data  []struct {
		Pid        int      `json:"pid"`
		P          int      `json:"p"`
		R18        bool     `json:"r18"`
		UID        int      `json:"uid"`
		Title      string   `json:"title"`
		Author     string   `json:"author"`
		Tags       []string `json:"tags"`
		Width      int      `json:"width"`
		Height     int      `json:"height"`
		Ext        string   `json:"ext"`
		UploadDate int64    `json:"uploadDate"`
		Urls       struct {
			Original string `json:"original"`
		} `json:"urls"`
	} `json:"data"`
}

func SetuKoKo(msg string) Setu {
	var r18 bool
	var keyword string
	var uid string
	var tag string
	params := "?num=1&proxy=i.pixiv.cat"
	if regexp.MustCompile(`r18`).Match([]byte(msg)) {
		r18 = true
	} else {
		r18 = false
	}
	if regexp.MustCompile(`keyword`).Match([]byte(msg)) {
		res := regexp.MustCompile(`keyword=[^=\s]*`).FindAllString(msg, 10)
		keyword = res[0][8:]
	}
	if regexp.MustCompile(`uid`).Match([]byte(msg)) {
		res := regexp.MustCompile(`uid=[0-9]*`).FindAllString(msg, 10)
		uid = res[0][4:]
	}
	if regexp.MustCompile(`tag`).Match([]byte(msg)) {
		res := regexp.MustCompile(`tag=[^=\s]*`).FindAllString(msg, 10)
		tag = res[0][4:]
	}
	if r18 {
		params = params + "&r18=1"
	}
	if keyword != "" {
		params = params + "&keyword=" + keyword
	}
	if uid != "" {
		params = params + "&uid=" + uid
	}
	if tag != "" {
		params = params + "&tag=" + tag
	}
	response, err := http.Get("https://api.lolicon.app/setu/v2" + params)
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var st Setu
	err = json.Unmarshal(body, &st)
	tools.Check(err)
	return st
}
