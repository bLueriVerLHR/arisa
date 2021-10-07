package main

import (
	"arisa/plugins"
	"arisa/tools"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

type TouchEvent struct {
	GroupID    int    `json:"group_id"`
	NoticeType string `json:"notice_type"`
	PostType   string `json:"post_type"`
	SelfID     int64  `json:"self_id"`
	SenderID   int64  `json:"sender_id"`
	SubType    string `json:"sub_type"`
	TargetID   int    `json:"target_id"`
	Time       int    `json:"time"`
	UserID     int64  `json:"user_id"`
}

type Recall struct {
	GroupID    int    `json:"group_id"`
	MessageID  int    `json:"message_id"`
	NoticeType string `json:"notice_type"`
	OperatorID int64  `json:"operator_id"`
	PostType   string `json:"post_type"`
	SelfID     int64  `json:"self_id"`
	Time       int    `json:"time"`
	UserID     int64  `json:"user_id"`
}

func GrepTalks(msg RecvMsg) bool {
	Gr := uni.Grep.Greps
	length := len(Gr)
	for i := 0; i < length; i++ {
		if regexp.MustCompile(Gr[i].Regexp).Match([]byte(msg.RawMessage)) {
			RespGroupMsg(msg, Gr[i].Message)
			return true
		}
	}
	return false
}

func RepeatSentence(msg RecvMsg) {
	if msg.RawMessage == "收到" {
		return
	}
	rand.Seed(time.Now().Unix())
	if msg.RawMessage != uni.LastMsg {
		uni.LastMsg = msg.RawMessage
		uni.Repeat = 0
	} else {
		uni.Repeat++
		if uni.Repeat >= 3 {
			if rand.Int()%4 == 0 {
				RespGroupMsg(msg, uni.LastMsg)
			}
		}
	}
}

func RespPoke(msg string) {
	var te TouchEvent
	err := json.Unmarshal([]byte(msg), &te)
	tools.Check(err)
	if te.SelfID == int64(te.TargetID) {
		if te.GroupID == 0 {
			RespPrivateNotice(int(te.SenderID), "( •̀ ω •́ )y塔诺西！")
			RespPrivateNotice(int(te.SenderID), "[CQ:poke,qq="+strconv.Itoa(int(te.SenderID))+"]")
		} else {
			RespGroupNotice(te.GroupID, "(●ˇ∀ˇ●)塔诺西！")
			RespGroupNotice(te.GroupID, "[CQ:poke,qq="+strconv.Itoa(int(te.SenderID))+"]")
		}
	}
}

func AntiRecall(msg string) {
	var rc Recall
	err := json.Unmarshal([]byte(msg), &rc)
	tools.Check(err)
	if rc.OperatorID == rc.SelfID {
		return
	}
	RespGroupNotice(rc.GroupID, "啊这？撤回了啥？")
}

func SomeSay(msg string) string {
	rand.Seed(time.Now().Unix())
	if tools.Grep(`hy：`, msg) {
		return uni.SomeS.Hy[rand.Int()%len(uni.SomeS.Hy)]
	} else if tools.Grep(`ss：`, msg) {
		switch rand.Int() % 3 {
		case 1:
			return "舔狗：" + plugins.XW2dtg()
		case 2:
			return "ss：" + plugins.XW2dyy()
		default:
			return uni.SomeS.Ss[rand.Int()%len(uni.SomeS.Ss)]
		}
	}
	return ""
}

func AntiMOOC(msg RecvMsg) {
	if value, ok := uni.Operating[msg.UserID]; ok && value {
		if regexp.MustCompile(`停下|stop|Stop|Quit|q|quit|Exit|exit`).Match([]byte(msg.Message)) {
			uni.Operating[msg.UserID] = false
			RespPrivateMsg(msg, "MOOC登录停止")
			return
		} else {
			RespPrivateMsg(msg, "MOOC登录正在运行，如果需要退出，请输入mooc.exit")
			return
		}
	}
	if tools.Grep(`mooc -i`, msg.Message) {
		uni.Operating[msg.UserID] = true
		var thisCookie = MoocLoop(msg)
		uni.Operating[msg.UserID] = false
		if thisCookie == "" {
			RespPrivateMsg(msg, "登录失败")
			return
		}
		RespPrivateMsg(msg, "登录成功！")
		if usr := HasUsr(msg.UserID); usr != -1 {
			uni.MoocA.List[usr].Cookie = thisCookie
			return
		}
		uni.MoocA.List = append(uni.MoocA.List, struct {
			Qq     int64  "json:\"qq\""
			Cookie string "json:\"cookie\""
			Lesson []struct {
				Name    string "json:\"name\""
				TermID  string "json:\"termId\""
				ClassID string "json:\"classId\""
			} "json:\"lesson\""
		}{
			msg.UserID,
			thisCookie,
			[]struct {
				Name    string "json:\"name\""
				TermID  string "json:\"termId\""
				ClassID string "json:\"classId\""
			}{{
				"occupy",
				"0",
				"0",
			},
			},
		},
		)
	} else if msg.RawMessage == "mooc -u" {
		usr := HasUsr(msg.UserID)
		if usr == -1 {
			RespPrivateMsg(msg, "Cookie可能没加载哦")
			return
		}
		now := tools.ReturnIntDate(tools.Time2China(time.Now()))
		retList := plugins.GetLessonList(uni.MoocA.List[usr].Cookie)
		if len(retList.Result.Result) == 0 {
			RespPrivateMsg(msg, "Cookie可能错了，或过期了哦")
			return
		}
		length := len(retList.Result.Result)
		var respList string
		for i := 0; i < length; i++ {
			if HasLesson(retList.Result.Result[i].Name, usr) {
				continue
			}
			if !(tools.ReturnIntDate(plugins.MOOCtime2China(retList.Result.Result[i].TermPanel.EndTime)) < now) {
				uni.MoocA.List[usr].Lesson = append(uni.MoocA.List[usr].Lesson, struct {
					Name    string "json:\"name\""
					TermID  string "json:\"termId\""
					ClassID string "json:\"classId\""
				}{
					retList.Result.Result[i].Name,
					strconv.Itoa(retList.Result.Result[i].TermPanel.ID),
					strconv.Itoa(retList.Result.Result[i].ID),
				})
				respList = respList + retList.Result.Result[i].Name + "\n"
			}
		}
		if respList == "" {
			RespPrivateMsg(msg, "所有课程均已经加载完毕！")
		} else {
			RespPrivateMsg(msg, "本次更新作业如下：\n\n"+respList)
		}
		uni.Save()
	} else if msg.RawMessage == "mooc -a" || msg.RawMessage == "mooc" {
		var resp string
		usr := HasUsr(msg.UserID)
		if usr == -1 {
			RespPrivateMsg(msg, "Cookie可能没加载哦")
			return
		}
		length := len(uni.MoocA.List[usr].Lesson)
		if length == 1 {
			RespPrivateMsg(msg, "没有课程哦")
			return
		}
		for i := 1; i < length; i++ {
			les := plugins.GetLessonInfo(uni.MoocA.List[usr].Cookie, uni.MoocA.List[usr].Lesson[i].TermID, uni.MoocA.List[usr].Lesson[i].ClassID)
			if len(les.Result.MocTermDto.CourseName) == 0 {
				RespPrivateMsg(msg, "Cookie可能错了，或过期了哦")
				return
			}
			resp = resp + "<" + les.Result.MocTermDto.CourseName + ">:\n"
			length := len(les.Result.MocTermDto.Chapters)
			for j := 0; j < length; j++ {
				hwl := len(les.Result.MocTermDto.Chapters[j].Homeworks)
				if hwl != 0 {
					for k := 0; k < hwl; k++ {
						mct := plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.EvaluateEnd)
						nt := tools.Time2China(time.Now())
						if tools.ReturnIntDate(nt) > tools.ReturnIntDate(mct) {
							continue
						}
						resp = resp + "作业名称：" + les.Result.MocTermDto.Chapters[j].Homeworks[k].Name + "\n"
						resp = resp + "作业开始：" + plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.Deadline).Format("2006-01-02 15:04:05") + "\n"
						resp = resp + "互评开始：" + plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.EvaluateStart).Format("2006-01-02 15:04:05") + "\n"
						resp = resp + "互评结束：" + plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.EvaluateEnd).Format("2006-01-02 15:04:05") + "\n" + "\n"
					}
				}
			}
			RespPrivateMsg(msg, resp)
			resp = ""
		}
	} else if msg.RawMessage == "mooc -l" {
		usr := HasUsr(msg.UserID)
		if usr == -1 {
			RespPrivateMsg(msg, "Cookie可能没加载哦")
			return
		}
		length := len(uni.MoocA.List[usr].Lesson)
		if length == 1 {
			RespPrivateMsg(msg, "没有课程哦")
			return
		}
		var respList string
		for i := 1; i < length; i++ {
			respList = respList + strconv.Itoa(i) + ". " + uni.MoocA.List[usr].Lesson[i].Name + "\n"
		}
		RespPrivateMsg(msg, respList)
	} else if tools.Grep(`^mooc\s-n\s[0-9]{1,}$`, msg.Message) {
		usr := HasUsr(msg.UserID)
		if usr == -1 {
			RespPrivateMsg(msg, "Cookie可能没加载哦")
			return
		}
		length := len(uni.MoocA.List[usr].Lesson)
		if length == 1 {
			RespPrivateMsg(msg, "没有课程哦")
			return
		}
		Strnum := msg.Message[8:]
		num, err := strconv.Atoi(Strnum)
		tools.Check(err)
		if num > length || num < 1 {
			RespPrivateMsg(msg, "无效序号")
			return
		}
		var resp string
		les := plugins.GetLessonInfo(uni.MoocA.List[usr].Cookie, uni.MoocA.List[usr].Lesson[num].TermID, uni.MoocA.List[usr].Lesson[num].ClassID)
		if len(les.Result.MocTermDto.CourseName) == 0 {
			RespPrivateMsg(msg, "Cookie可能错了，或过期了哦")
			return
		}
		resp = resp + "<" + les.Result.MocTermDto.CourseName + ">:\n"
		length = len(les.Result.MocTermDto.Chapters)
		for j := 0; j < length; j++ {
			hwl := len(les.Result.MocTermDto.Chapters[j].Homeworks)
			if hwl != 0 {
				for k := 0; k < hwl; k++ {
					mct := plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.EvaluateEnd)
					nt := tools.Time2China(time.Now())
					if tools.ReturnIntDate(nt) > tools.ReturnIntDate(mct) {
						continue
					}
					resp = resp + "作业名称：" + les.Result.MocTermDto.Chapters[j].Homeworks[k].Name + "\n"
					resp = resp + "作业开始：" + plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.Deadline).Format("2006-01-02 15:04:05") + "\n"
					resp = resp + "互评开始：" + plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.EvaluateStart).Format("2006-01-02 15:04:05") + "\n"
					resp = resp + "互评结束：" + plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.EvaluateEnd).Format("2006-01-02 15:04:05") + "\n" + "\n"
				}
			}
		}
		RespPrivateMsg(msg, resp)
	} else if msg.Message == "mooc -c" {
		usr := HasUsr(msg.UserID)
		if usr == -1 {
			RespPrivateMsg(msg, "Cookie可能没加载哦")
			return
		}
		uni.MoocA.List[usr].Lesson = []struct {
			Name    string "json:\"name\""
			TermID  string "json:\"termId\""
			ClassID string "json:\"classId\""
		}{{
			"occupy",
			"0",
			"0",
		}}
	} else if msg.Message == "mooc -s" {
		if HasUsr(msg.UserID) != -1 {
			RespPrivateMsg(msg, "已经订阅了哦！")
			return
		}
		uni.Slist.Mooc = append(uni.Slist.Mooc, msg.UserID)
		uni.Save()
	}
}

func MoocLoop(msg RecvMsg) string {
	var thisCookie string
	for uni.Operating[msg.UserID] {
		QR := plugins.MOOCLoginQRcode()
		res := RespPrivateMsg(msg, "[CQ:image,file="+QR.Result.CodeURL+"]")
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		tools.Check(err)
		var del MsgId
		err = json.Unmarshal(body, &del)
		tools.Check(err)
		go DeleteLater(del.Data.MessageID, 60)
		tmp, flag, cookies := plugins.CheckingQRcode(QR.Result.PollKey)
		length := len(cookies)
		var STUDY_WTR string
		for i := 0; i < length; i++ {
			if cookies[i].Name == "STUDY_WTR" {
				STUDY_WTR = cookies[i].Value
				break
			}
		}
		if flag {
			cookies, ref := plugins.MocMobChangeCookie(tmp, STUDY_WTR)
			length := len(cookies)
			for i := 0; i < length; i++ {
				thisCookie += cookies[i].Raw + ";"
			}
			cookies = plugins.GetNeteaseWDAuid(thisCookie, ref)
			length = len(cookies)
			for i := 0; i < length; i++ {
				thisCookie += cookies[i].Raw + ";"
			}
			return thisCookie
		}
	}
	return ""
}

func MOOClock(id int64) {
	var resp string
	usr := HasUsr(id)
	if usr == -1 {
		MOOCClockResp(strconv.FormatInt(id, 10), "Cookie可能没加载哦")
		return
	}
	length := len(uni.MoocA.List[usr].Lesson)
	if length == 1 {
		MOOCClockResp(strconv.FormatInt(id, 10), "没有课程哦")
		return
	}
	for i := 1; i < length; i++ {
		les := plugins.GetLessonInfo(uni.MoocA.List[usr].Cookie, uni.MoocA.List[usr].Lesson[i].TermID, uni.MoocA.List[usr].Lesson[i].ClassID)
		if len(les.Result.MocTermDto.CourseName) == 0 {
			MOOCClockResp(strconv.FormatInt(id, 10), "Cookie可能错了，或过期了哦")
			return
		}
		resp = resp + "<" + les.Result.MocTermDto.CourseName + ">:\n"
		length := len(les.Result.MocTermDto.Chapters)
		for j := 0; j < length; j++ {
			hwl := len(les.Result.MocTermDto.Chapters[j].Homeworks)
			if hwl != 0 {
				for k := 0; k < hwl; k++ {
					mct := plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.EvaluateEnd)
					nt := tools.Time2China(time.Now())
					if tools.ReturnIntDate(nt) > tools.ReturnIntDate(mct) {
						continue
					}
					resp = resp + "作业名称：" + les.Result.MocTermDto.Chapters[j].Homeworks[k].Name + "\n"
					resp = resp + "作业开始：" + plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.Deadline).Format("2006-01-02 15:04:05") + "\n"
					resp = resp + "互评开始：" + plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.EvaluateStart).Format("2006-01-02 15:04:05") + "\n"
					resp = resp + "互评结束：" + plugins.MOOCtime2China(les.Result.MocTermDto.Chapters[j].Homeworks[k].Test.EvaluateEnd).Format("2006-01-02 15:04:05") + "\n" + "\n"
				}
			}
		}
		MOOCClockResp(strconv.FormatInt(id, 10), resp)
		resp = ""
	}
}

func MOOCClockResp(id string, text string) *http.Response {
	res, err := http.PostForm("http://"+uni.Conf.BotConfig.Host+":"+strconv.Itoa(uni.Conf.BotConfig.Port)+"/send_private_msg", url.Values{"user_id": {id}, "message": {text}})
	tools.Check(err)
	return res
}

func MOOCSubscribe() {
	if t := tools.Time2China(time.Now()).Hour(); !uni.Report && t == 8 {
		uni.Report = true
		length := len(uni.Slist.Mooc)
		for i := 0; i < length; i++ {
			MOOClock(uni.Slist.Mooc[i])
		}
	} else if t < 8 {
		uni.Report = false
	}
}

func HasUsr(ID int64) int {
	length := len(uni.MoocA.List)
	for i := 0; i < length; i++ {
		if uni.MoocA.List[i].Qq == ID {
			return i
		}
	}
	return -1
}

func HasLesson(ClassName string, usr int) bool {
	length := len(uni.MoocA.List[usr].Lesson)
	for i := 0; i < length; i++ {
		if uni.MoocA.List[usr].Lesson[i].Name == ClassName {
			return true
		}
	}
	return false
}

func IsWhite(target int, list []int) bool {
	length := len(list)
	for i := 0; i < length; i++ {
		if target == list[i] {
			return true
		}
	}
	return false
}

func RandEvent(msg RecvMsg) *http.Response {
	var rd int
	var err error
	if reg := regexp.MustCompile(`mode=[0-9^\s]{1,}`); reg.Match([]byte(msg.Message)) {
		rd, err = strconv.Atoi(string(reg.Find([]byte(msg.Message))[5:]))
		fmt.Println(rd)
		tools.Check(err)
	} else {
		rand.Seed(time.Now().Unix())
		rd = rand.Int() % 23
		fmt.Println(rd)
	}
	switch rd {
	case 0:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.DY+",title=???]")
	case 1:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.Dtstubg+",title=???]")
	case 2:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.Dtstudm+",title=???]")
	case 3:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.Dtstudm_m+",title=???]")
	case 4:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.Dtstudmmn+",title=???]")
	case 5:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.Dtstudmmn_m+",title=???]")
	case 6:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.Dtstumn+",title=???]")
	case 7:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.Dtstumn_m+",title=???]")
	case 8:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.PCbg+",title=???]")
	case 9:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.PPX+",title=???]")
	case 10:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.SEX+",title=???]")
	case 11:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.Padbg+",title=???]")
	case 12:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.TX+",title=???]")
	case 14:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.Xjj1+",title=???]")
	case 15:
		return RespGroupMsg(msg, "[CQ:share,url="+plugins.Xjj2+",title=???]")
	case 16:
		return RespGroupMsg(msg, "[CQ:image,file="+plugins.Btu()+",type=flash]")
	case 17:
		return RespGroupMsg(msg, "[CQ:image,file="+plugins.EEBT()+",type=flash]")
	case 18:
		return RespGroupMsg(msg, "[CQ:image,file="+plugins.Lolibj()+",type=flash]")
	case 19:
		return RespGroupMsg(msg, "[CQ:image,file="+plugins.Random2Dpic()+",type=flash]")
	case 20:
		return RespGroupMsg(msg, "[CQ:image,file="+plugins.XW2dbg()+",type=flash]")
	case 21:
		return RespGroupMsg(msg, "[CQ:image,file="+plugins.XW2ddm()+",type=flash]")
	case 22:
		return RespGroupMsg(msg, "[CQ:image,file="+plugins.XW2dmc()+",type=flash]")
	}
	return nil
}

func BcyCPN(msg RecvMsg) {
	var err error
	var res []string
	if regexp.MustCompile(`小说|文|粮`).Match([]byte(msg.Message)) {
		var st string
		if regexp.MustCompile(`日|天`).Match([]byte(msg.Message)) {
			st = "lastday"
		} else if regexp.MustCompile(`新人`).Match([]byte(msg.Message)) {
			st = "newPeople"
		} else {
			st = "week"
		}
		var num int
		reg := regexp.MustCompile(`rank=[0-9^\s]{1,}`).Find([]byte(msg.Message))
		if len(reg) != 0 {
			num, err = strconv.Atoi(string(reg)[5:])
			tools.Check(err)
			if num < 1 || num > 100 {
				num = 1
			}
		}
		var date int
		reg = regexp.MustCompile(`date=[0-9^\s]{8}`).Find([]byte(msg.Message))
		if len(reg) != 0 {
			date, err = strconv.Atoi(string(reg)[5:])
			tools.Check(err)
		} else {
			date = tools.ReturnIntDate(tools.Time2China(time.Now()))
		}
		if num != 0 {
			res = plugins.Bcy("novel", strconv.Itoa(num/20+1), st, strconv.Itoa(date), (num-1)%20)
		} else {
			rand.Seed(time.Now().Unix())
			num = rand.Int()%100 + 1
			res = plugins.Bcy("novel", strconv.Itoa(num/20+1), st, strconv.Itoa(date), (num-1)%20)
		}
		if len(res) != 0 {
			RespGroupMsg(msg, res[0])
		} else {
			RespGroupMsg(msg, "似乎获取失败了")
		}
	} else if regexp.MustCompile(`图`).Match([]byte(msg.Message)) {
		var st string
		if regexp.MustCompile(`日|天`).Match([]byte(msg.Message)) {
			st = "lastday"
		} else if regexp.MustCompile(`新人`).Match([]byte(msg.Message)) {
			st = "newPeople"
		} else {
			st = "week"
		}
		var num int
		reg := regexp.MustCompile(`rank=[0-9^\s]{1,}`).Find([]byte(msg.Message))
		if len(reg) != 0 {
			num, err = strconv.Atoi(string(reg)[5:])
			tools.Check(err)
			if num < 1 || num > 100 {
				num = 1
			}
		}
		var date int
		reg = regexp.MustCompile(`date=[0-9^\s]{8}`).Find([]byte(msg.Message))
		if len(reg) != 0 {
			date, err = strconv.Atoi(string(reg)[5:])
			tools.Check(err)
		} else {
			date = tools.ReturnIntDate(tools.Time2China(time.Now()))
		}
		if num != 0 {
			res = plugins.Bcy("illust", strconv.Itoa(num/20+1), st, strconv.Itoa(date), (num-1)%20)
		} else {
			rand.Seed(time.Now().Unix())
			num = rand.Int()%100 + 1
			res = plugins.Bcy("illust", strconv.Itoa(num/20+1), st, strconv.Itoa(date), (num-1)%20)
		}
		if length := len(res); length != 0 {
			RespGroupMsg(msg, res[0])
			for i := 1; i < length; i++ {
				RespGroupMsg(msg, "[CQ:image,file="+res[i]+"]")
			}
		} else {
			RespGroupMsg(msg, "似乎获取失败了")
		}
	} else if regexp.MustCompile(`coser|cos`).Match([]byte(msg.Message)) {
		var st string
		if regexp.MustCompile(`日|天`).Match([]byte(msg.Message)) {
			st = "lastday"
		} else if regexp.MustCompile(`新人`).Match([]byte(msg.Message)) {
			st = "newPeople"
		} else {
			st = "week"
		}
		var num int
		reg := regexp.MustCompile(`rank=[0-9^\s]{1,}`).Find([]byte(msg.Message))
		if len(reg) != 0 {
			num, err = strconv.Atoi(string(reg)[5:])
			tools.Check(err)
			if num < 1 || num > 100 {
				num = 1
			}
		}
		var date int
		reg = regexp.MustCompile(`date=[0-9^\s]{8}`).Find([]byte(msg.Message))
		if len(reg) != 0 {
			date, err = strconv.Atoi(string(reg)[5:])
			tools.Check(err)
		} else {
			date = tools.ReturnIntDate(tools.Time2China(time.Now()))
		}
		if num != 0 {
			res = plugins.Bcy("cos", strconv.Itoa(num/20+1), st, strconv.Itoa(date), (num-1)%20)
		} else {
			rand.Seed(time.Now().Unix())
			num = rand.Int()%100 + 1
			res = plugins.Bcy("cos", strconv.Itoa(num/20+1), st, strconv.Itoa(date), (num-1)%20)
		}
		if length := len(res); length != 0 {
			RespGroupMsg(msg, res[0])
			for i := 1; i < length; i++ {
				RespGroupMsg(msg, "[CQ:image,file="+res[i]+"]")
			}
		} else {
			RespGroupMsg(msg, "似乎获取失败了")
		}
	}
}
