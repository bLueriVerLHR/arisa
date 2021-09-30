package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (mc *MOOC) Read(recv string) {
	flag := []int{0, 0}
	length := len(recv)
	for i := 0; i < length; i++ {
		if recv[i] == '{' {
			flag[0] = i
			break
		}
	}
	for i := length - 1; i >= 0; i-- {
		if recv[i] == '}' {
			flag[1] = i
			break
		}
	}
	err := json.Unmarshal([]byte(recv[flag[0]:flag[1]+1]), mc)
	CheckError(err)
}

func GetMOOCinfo(lessonCode string, termID string) []string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.icourse163.org/web/j/courseBean.getLastLearnedMocTermDto.rpc?csrfKey="+csrfKey, strings.NewReader("termId="+termID))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	text, err := ioutil.ReadFile(cookiePath)
	CheckError(err)
	req.Header.Set("Host", "www.icourse163.org")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,en-US;q=0.7,ja;q=0.3")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("edu-script-token", csrfKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", "17")
	req.Header.Set("Origin", "https://www.icourse163.org")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.icourse163.org/learn/"+lessonCode+"?tid="+termID)
	req.Header.Set("Cookie", string(text))
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")
	resp, err := client.Do(req)
	CheckError(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)

	index := regexp.MustCompile(`homeworks`).FindAllIndex(body, 10000)
	strBody := string(body)
	length := len(index)
	var allHomeWork []string

	for i := 0; i < length; i++ {
		strRange := [2]int{0, 0}
		for j := 0; j < 2048; j++ {
			if strBody[index[i][0]+j] == '[' {
				strRange[0] = index[i][0] + j + 1
			}
			if strBody[index[i][0]+j] == ']' {
				strRange[1] = index[i][0] + j
				break
			}
		}
		if strRange[0] != strRange[1] {
			allHomeWork = append(allHomeWork, strBody[strRange[0]:strRange[1]])
		}
	}
	length = len(allHomeWork)
	ret := make([]string, length)
	loc, err := time.LoadLocation("Asia/Shanghai")
	CheckError(err)
	no, err := ReturnDate(time.Now().In(loc))
	CheckError(err)
	for i := 0; i < length; i++ {
		var mc MOOC
		mc.Read(allHomeWork[i])
		tmp, err := ReturnDate(time.Unix(mc.Test.EvaluateEnd/1000, 0).In(loc))
		CheckError(err)
		if tmp < no {
			continue
		}
		ret[i] = ret[i] + "作业名称：" + mc.Name + "\n"
		ret[i] = ret[i] + "作业结束时间：" + time.Unix(mc.Test.Deadline/1000, 0).In(loc).Format("2006-01-02 15:04:05") + "\n"
		ret[i] = ret[i] + "互评开始时间：" + time.Unix(mc.Test.EvaluateStart/1000, 0).In(loc).Format("2006-01-02 15:04:05") + "\n"
		ret[i] = ret[i] + "互评结束时间：" + time.Unix(mc.Test.EvaluateEnd/1000, 0).In(loc).Format("2006-01-02 15:04:05") + "\n" + "\n"
	}
	return ret
}

func ReturnDate(in time.Time) (int, error) {
	y, m, d := in.Date()
	year := strconv.Itoa(y)
	var month string
	var day string
	switch m.String() {
	case "January":
		month = "01"
	case "February":
		month = "02"
	case "March":
		month = "03"
	case "April":
		month = "04"
	case "May":
		month = "05"
	case "June":
		month = "06"
	case "July":
		month = "07"
	case "August":
		month = "08"
	case "September":
		month = "09"
	case "October":
		month = "10"
	case "November":
		month = "11"
	case "December":
		month = "12"
	}
	if d < 10 {
		day = "0" + strconv.Itoa(d)
	} else {
		day = strconv.Itoa(d)
	}
	return strconv.Atoi(year + month + day)
}

func FormatMOOC() string {
	var resp string
	lesson := []Mooc{
		{"NEU-1002822009", "1465215445", "概率论与数理统计"},
		{"NEU-1461807175", "1465298442", "数据科学基础（Matlab）"},
		{"NEU-1001638003", "1465301447", "复变函数与积分变换"},
	}
	length := len(lesson)
	for i := 0; i < length; i++ {
		resp = resp + lesson[i].className + ":\n"
		ret := GetMOOCinfo(lesson[i].lessonId, lesson[i].termId)
		size := len(ret)
		for j := 0; j < size; j++ {
			resp = resp + ret[j]
		}
	}
	return resp
}
