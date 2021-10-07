package plugins

import (
	"arisa/tools"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	Xjj1        = "https://api.nmb.show/xiaojiejie1.php"
	Xjj2        = "https://api.nmb.show/xiaojiejie2.php"
	acg1985     = "https://api.nmb.show/1985acg.php"
	Dtstudm     = "http://api.btstu.cn/sjbz/?lx=dongman"
	Dtstumn     = "http://api.btstu.cn/sjbz/?lx=meizi"
	Dtstudmmn   = "http://api.btstu.cn/sjbz/?lx=suiji"
	Dtstudm_m   = "http://api.btstu.cn/sjbz/?lx=m_dongman"
	Dtstumn_m   = "http://api.btstu.cn/sjbz/?lx=m_meizi"
	Dtstudmmn_m = "http://api.btstu.cn/sjbz/?m_lx=suiji"
	Dtstubg     = "http://api.btstu.cn/sjbz/zsy.php"
)

type Lolibjpic struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

type XW2dpic struct {
	Code   string `json:"code"`
	Imgurl string `json:"imgurl"`
	Width  string `json:"width"`
	Height string `json:"height"`
}
type EEBTpic struct {
	Error  int    `json:"error"`
	Result int    `json:"result"`
	Img    string `json:"img"`
}

type Random2D struct {
	Code   string `json:"code"`
	Imgurl string `json:"imgurl"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

type Btupic struct {
	Code int    `json:"code"`
	URL  string `json:"url"`
}

func Lolibj() string {
	response, err := http.Get("https://api.loli.bj/api/?type=json")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var r2 Lolibjpic
	err = json.Unmarshal(body, &r2)
	tools.Check(err)
	r2.URL = strings.ReplaceAll(r2.URL, "\\", "")
	return r2.URL
}

func XW2ddm() string {
	response, err := http.Get("https://api.ixiaowai.cn/api/api.php?return=json")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var r2 XW2dpic
	err = json.Unmarshal(body, &r2)
	tools.Check(err)
	r2.Imgurl = strings.ReplaceAll(r2.Imgurl, "\\", "")
	return r2.Imgurl
}

func XW2dmc() string {
	response, err := http.Get("https://api.ixiaowai.cn/mcapi/mcapi.php?return=json")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var r2 XW2dpic
	err = json.Unmarshal(body, &r2)
	tools.Check(err)
	r2.Imgurl = strings.ReplaceAll(r2.Imgurl, "\\", "")
	return r2.Imgurl
}

func XW2dbg() string {
	response, err := http.Get("https://api.ixiaowai.cn/gqapi/gqapi.php?return=json")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var r2 XW2dpic
	err = json.Unmarshal(body, &r2)
	tools.Check(err)
	r2.Imgurl = strings.ReplaceAll(r2.Imgurl, "\\", "")
	return r2.Imgurl
}

func XW2dyy() string {
	response, err := http.Get("https://api.ixiaowai.cn/api/ylapi.php")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	return string(body)
}

func XW2dtg() string {
	response, err := http.Get("https://api.ixiaowai.cn/tgrj/index.php")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	return string(body)
}

func EEBT() string {
	rand.Seed(time.Now().Unix())
	response, err := http.Get("http://pic.eebt.com/api.php?mom=json")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var r2 EEBTpic
	err = json.Unmarshal(body, &r2)
	tools.Check(err)
	r2.Img = strings.ReplaceAll(r2.Img, "\\", "")
	return r2.Img
}

func Random2Dpic() string {
	response, err := http.Get("https://www.dmoe.cc/random.php?return=json")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var r2 Random2D
	err = json.Unmarshal(body, &r2)
	tools.Check(err)
	r2.Imgurl = strings.ReplaceAll(r2.Imgurl, "\\", "")
	return r2.Imgurl
}

func Btu() string {
	rand.Seed(time.Now().Unix())
	response, err := http.Get("http://img.btu.pp.ua/random/api.php?type=json&mode=" + strconv.Itoa(rand.Int()%2+1))
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var r2 Btupic
	err = json.Unmarshal(body, &r2)
	tools.Check(err)
	r2.URL = strings.ReplaceAll(r2.URL, "\\", "")
	return r2.URL
}
