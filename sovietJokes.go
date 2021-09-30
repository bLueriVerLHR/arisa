package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
)

func SovietJokes(code int) string {
	text, err := ioutil.ReadFile(JokePath)
	CheckError(err)
	if code == 289 {
		index := regexp.MustCompile(strconv.Itoa(code) + "、").FindIndex(text)
		return "苏联政治冷笑话：\n" + string(text)[index[0]+5:]
	} else {
		index1 := regexp.MustCompile(strconv.Itoa(code) + "、").FindIndex(text)
		for len(index1) == 0 {
			code++
			index1 = regexp.MustCompile(strconv.Itoa(code) + "、").FindIndex(text)
		}
		index2 := regexp.MustCompile(strconv.Itoa(code+1) + "、").FindIndex(text)
		for len(index2) == 0 {
			code++
			index2 = regexp.MustCompile(strconv.Itoa(code) + "、").FindIndex(text)
		}
		return "苏联政治冷笑话：\n" + string(text)[index1[0]+8:index2[0]-1]
	}
}
