package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func QMsearch(target string) qqmusic {
	response, err := http.Get("https://c.y.qq.com/soso/fcgi-bin/client_search_cp?ct=24&qqmusic_ver=1298&new_json=1&remoteplace=txt.yqq.song&searchid=&t=0&aggr=1&cr=1&catZhida=1&lossless=0&flag_qc=0&p=1&n=1&w=\"" + target + "\"")
	CheckError(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	var qm qqmusic
	length := len(body)
	err = json.Unmarshal(body[9:length-1], &qm)
	CheckError(err)
	return qm
}
