package plugins

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type IsoyuNews struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []struct {
		SourceID        string `json:"sourceId"`
		Template        string `json:"template,omitempty"`
		RiskLevel       int    `json:"riskLevel"`
		UpTimes         int    `json:"upTimes"`
		Lmodify         string `json:"lmodify"`
		Source          string `json:"source"`
		Postid          string `json:"postid"`
		Title           string `json:"title"`
		Mtime           string `json:"mtime"`
		HasImg          int    `json:"hasImg,omitempty"`
		Topicid         string `json:"topicid"`
		TopicBackground string `json:"topic_background,omitempty"`
		Digest          string `json:"digest"`
		Boardid         string `json:"boardid"`
		Alias           string `json:"alias,omitempty"`
		HasAD           int    `json:"hasAD,omitempty"`
		Imgsrc          string `json:"imgsrc"`
		Ptime           string `json:"ptime"`
		Daynum          string `json:"daynum"`
		ExtraShowFields string `json:"extraShowFields"`
		HasHead         int    `json:"hasHead,omitempty"`
		Order           int    `json:"order,omitempty"`
		Votecount       int    `json:"votecount"`
		HasCover        bool   `json:"hasCover,omitempty"`
		Docid           string `json:"docid"`
		Tname           string `json:"tname,omitempty"`
		URL3W           string `json:"url_3w"`
		Priority        int    `json:"priority"`
		DownTimes       int    `json:"downTimes"`
		URL             string `json:"url"`
		Quality         int    `json:"quality"`
		CommentStatus   int    `json:"commentStatus"`
		Ads             []struct {
			Subtitle string `json:"subtitle"`
			SkipType string `json:"skipType"`
			SkipID   string `json:"skipID"`
			Tag      string `json:"tag"`
			Title    string `json:"title"`
			Imgsrc   string `json:"imgsrc"`
			URL      string `json:"url"`
		} `json:"ads,omitempty"`
		Ename      string `json:"ename,omitempty"`
		ReplyCount int    `json:"replyCount"`
		Ltitle     string `json:"ltitle"`
		HasIcon    bool   `json:"hasIcon,omitempty"`
		Subtitle   string `json:"subtitle"`
		Cid        string `json:"cid,omitempty"`
	} `json:"data"`
}

type Local struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []struct {
		Imgsrc3Gtype string `json:"imgsrc3gtype"`
		Stitle       string `json:"stitle"`
		Docid        string `json:"docid"`
		Digest       string `json:"digest"`
		Source       string `json:"source"`
		Imgsrc       string `json:"imgsrc"`
		Title        string `json:"title"`
		Priority     int    `json:"priority"`
		Ptime        string `json:"ptime"`
		HasImg       int    `json:"hasImg,omitempty"`
		URL          string `json:"url"`
		CommentCount int    `json:"commentCount"`
		SkipURL      string `json:"skipURL,omitempty"`
		SkipType     string `json:"skipType,omitempty"`
		Modelmode    string `json:"modelmode,omitempty"`
	} `json:"data"`
}

type Iscos struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []struct {
		Postid       string   `json:"postid"`
		Desc         string   `json:"desc"`
		Pvnum        string   `json:"pvnum"`
		Createdate   string   `json:"createdate"`
		Scover       string   `json:"scover"`
		Setname      string   `json:"setname"`
		Cover        string   `json:"cover"`
		Pics         []string `json:"pics"`
		Clientcover1 string   `json:"clientcover1"`
		Replynum     string   `json:"replynum"`
		Topicname    string   `json:"topicname"`
		Setid        string   `json:"setid"`
		Seturl       string   `json:"seturl"`
		Datetime     string   `json:"datetime"`
		Clientcover  string   `json:"clientcover"`
		Imgsum       string   `json:"imgsum"`
		Tcover       string   `json:"tcover"`
	} `json:"data"`
}

func IsoyuN(num string) string {
	resp, _ := http.Get("https://api.isoyu.com/api/News/new_list?page=0&type=" + num)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var news IsoyuNews
	_ = json.Unmarshal(body, &news)
	var res = "新闻来啦！\n"
	var length = len(news.Data)
	for i := 0; i < length; i++ {
		res += strconv.Itoa(i+1) + ". " + news.Data[i].Title + "\n" + news.Data[i].URL + "\n"
	}
	return res
}

func IsoyuLocal() string {
	resp, _ := http.Get("https://api.isoyu.com/api/News/local_news?name=%E8%BE%BD%E5%AE%81%E7%9C%81_%E6%B2%88%E9%98%B3%E5%B8%82&page=0")
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var news Local
	_ = json.Unmarshal(body, &news)
	var res = "新闻来啦！\n"
	var length = len(news.Data)
	for i := 0; i < length; i++ {
		res += strconv.Itoa(i+1) + ". " + news.Data[i].Title + "\n" + news.Data[i].URL + "\n"
	}
	return res

}

func IsoyuCos() []string {
	rand.Seed(time.Now().Unix())
	ord := rand.Int() % 192
	resp, _ := http.Get("https://api.isoyu.com/api/picture/index?page=" + strconv.Itoa(ord))
	ord = rand.Int() % 10
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var cos Iscos
	var res []string
	_ = json.Unmarshal(body, &cos)
	length := len(cos.Data)
	if length != 0 {
		res = append(res, cos.Data[ord].Setname)
		res = append(res, cos.Data[ord].Pics...)
	}
	return res
}
