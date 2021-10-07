package tools

import (
	"strconv"
	"time"
)

func ReturnIntDate(in time.Time) int {
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
	ret, err := strconv.Atoi(year + month + day)
	Check(err)
	return ret
}

func Time2China(input time.Time) time.Time {
	China, err := time.LoadLocation("Asia/Shanghai")
	Check(err)
	return input.In(China)
}
