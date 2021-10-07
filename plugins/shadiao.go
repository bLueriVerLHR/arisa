package plugins

import (
	"arisa/tools"
	"io/ioutil"
	"net/http"
)

func Chp() string {
	response, err := http.Get("https://chp.shadiao.app/api.php")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	return string(body)
}

func Nmsl() string {
	response, err := http.Get("https://nmsl.shadiao.app/api.php?level=min&lang=zh_cn")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	return string(body)
}
