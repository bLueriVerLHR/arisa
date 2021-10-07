package plugins

import (
	"arisa/tools"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type PixivTops struct {
	Code int `json:"code"`
	Data []struct {
		ID        int         `json:"id"`
		Source    string      `json:"source"`
		URL       string      `json:"url"`
		Thumb     string      `json:"thumb"`
		Rank      int         `json:"rank"`
		Date      string      `json:"date"`
		CreatedAt string      `json:"created_at"`
		UpdatedAt interface{} `json:"updated_at"`
		Title     string      `json:"title"`
	} `json:"data"`
}

func SetuRank(msg string) (PixivTops, string) {
	var rank string
	var date string
	var pic_type string
	var params string
	if regexp.MustCompile(`rank`).Match([]byte(msg)) {
		res := regexp.MustCompile(`rank=[0-9]{1,2}`).FindAllString(msg, 10)
		if len(res) == 0 {
			rank = "1"
		} else {
			rank = res[0][5:]
		}
	}
	if regexp.MustCompile(`date`).Match([]byte(msg)) {
		res := regexp.MustCompile(`date=[0-9]{4}-[0-9]{1,2}-[0-9]{1,2}`).FindAllString(msg, 10)
		if len(res) == 0 {
			date = ""
		} else {
			date = res[0][5:]
		}
	}
	if regexp.MustCompile(`type`).Match([]byte(msg)) {
		res := regexp.MustCompile(`type=(pixiv_normal|pixiv_male_r18|pixiv_male|yandere)`).FindAllString(msg, 10)
		if len(res) == 0 {
			pic_type = ""
		} else {
			pic_type = res[0][5:]
		}
	}
	if pic_type != "" && regexp.MustCompile(pic_type).Match([]byte("pixiv_normal, pixiv_male_r18, pixiv_male, yandere")) {
		params = params + "&type=" + pic_type
	} else {
		params = params + "&type=" + "pixiv_normal"
	}
	if date != "" {
		params = params + "&date=" + date
	} else {
		loc, err := time.LoadLocation("Asia/Shanghai")
		tools.Check(err)
		params = params + "&date=" + time.Now().In(loc).Format("2006-01-02 15:04:05")[:10]
	}
	response, err := http.Get("https://gallery.nyadora.moe/" + "?" + params)
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var p PixivTops
	if regexp.MustCompile(`Nobody here but us chickens!|Fatal error`).Match(body) {
		return p, rank
	}
	err = json.Unmarshal(body, &p)
	tools.Check(err)
	return p, rank
}
