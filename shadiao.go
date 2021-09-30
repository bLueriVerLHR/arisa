package main

import (
	"io/ioutil"
	"net/http"
)

func chp() string {
	response, err := http.Get("https://chp.shadiao.app/api.php")
	CheckError(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	return string(body)
}

func nmsl() string {
	response, err := http.Get("https://nmsl.shadiao.app/api.php?level=min&lang=zh_cn")
	CheckError(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	return string(body)
}
