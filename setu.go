package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func SetuKoKo(msg RecvMsg) Setu {
	if !regexp.MustCompile(strconv.Itoa(msg.GroupID)).Match([]byte(setuAllow)) {
		return Setu{}
	}
	var r18 bool
	var keyword string
	var uid string
	var tag string
	params := "?num=1&proxy=i.pixiv.cat"
	if regexp.MustCompile(`r18`).Match([]byte(msg.RawMessage)) {
		r18 = true
	} else {
		r18 = false
	}
	if regexp.MustCompile(`keyword`).Match([]byte(msg.RawMessage)) {
		res := regexp.MustCompile(`keyword=[^=\s]*`).FindAllString(msg.RawMessage, 10)
		keyword = res[0][8:]
	}
	if regexp.MustCompile(`uid`).Match([]byte(msg.RawMessage)) {
		res := regexp.MustCompile(`uid=[0-9]*`).FindAllString(msg.RawMessage, 10)
		uid = res[0][4:]
	}
	if regexp.MustCompile(`tag`).Match([]byte(msg.RawMessage)) {
		res := regexp.MustCompile(`tag=[^=\s]*`).FindAllString(msg.RawMessage, 10)
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
	response, err := http.Get(SETU_URL + params)
	CheckError(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	var st Setu
	st.Read(string(body))
	return st
}

func SetuRank(msg RecvMsg) (Pixiv, string) {
	var rank string
	var date string
	var pic_type string
	var params string
	if regexp.MustCompile(`rank`).Match([]byte(msg.RawMessage)) {
		res := regexp.MustCompile(`rank=[0-9]{1,2}`).FindAllString(msg.RawMessage, 10)
		rank = res[0][5:]
	}
	if regexp.MustCompile(`date`).Match([]byte(msg.RawMessage)) {
		res := regexp.MustCompile(`date=[0-9]{4}-[0-9]{1,2}-[0-9]{1,2}`).FindAllString(msg.RawMessage, 10)
		date = res[0][5:]
	}
	if regexp.MustCompile(`type`).Match([]byte(msg.RawMessage)) {
		res := regexp.MustCompile(`type=(pixiv_normal|pixiv_male_r18|pixiv_male|yandere)`).FindAllString(msg.RawMessage, 10)
		pic_type = res[0][5:]
	}
	if pic_type != "" && regexp.MustCompile(strconv.Itoa(msg.GroupID)).Match([]byte(setuAllow)) && regexp.MustCompile(pic_type).Match([]byte("pixiv_normal, pixiv_male_r18, pixiv_male, yandere")) {
		params = params + "&type=" + pic_type
	} else {
		params = params + "&type=" + "pixiv_normal"
	}
	if date != "" {
		params = params + "&date=" + date
	} else {
		loc, err := time.LoadLocation("Asia/Shanghai")
		CheckError(err)
		params = params + "&date=" + time.Now().In(loc).Format("2006-01-02 15:04:05")[:11]
	}
	response, err := http.Get(gallery + "?" + params)
	CheckError(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	var p Pixiv
	if regexp.MustCompile(`Nobody here but us chickens!|Fatal error`).Match(body) {
		return p, rank
	}
	p.Read(string(body))
	return p, rank
}
