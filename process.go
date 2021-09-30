package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

func HandleMsg(msg string) {
	var message RecvMsg
	message.Read(msg)
	if message.MessageType == "group" {
		GroupHeater(message)
		fmt.Println(message.GroupID, message.UserID, message.Sender.Nickname, message.RawMessage)
	}
	if message.MessageType == "private" {
		Whisper(message)
		fmt.Println(message.UserID, message.Sender.Nickname, message.RawMessage)
	}
}

func HandleMetaInfo(metaInfo string) {
	var meta BotStatus
	meta.Read(metaInfo)
	status := meta.Status.AppGood && meta.Status.AppEnabled && meta.Status.AppInitialized && meta.Status.Good && meta.Status.Online
	if !status {
		fmt.Print("Something wrong with bot.\n")
	}
}

func HandlePoke(msg string) {
	var te TouchEvent
	te.Read(msg)
	if te.SelfID == int64(te.TargetID) {
		respGroupNotice(te.GroupID, "mufumufu[CQ:image,file=729173096bca29990bb8a8d1b7bca32b.image]")
		respGroupNotice(te.GroupID, "[CQ:poke,qq="+strconv.Itoa(int(te.SenderID))+"]")
	}
}

func Whisper(msg RecvMsg) {
	Action(msg, `Arisa --help`, func() {
		respPrivateMsg(msg, "私聊的功能和群聊会大相径庭，目前还在开发中。"+`目前支持：
正则聊天
`)
	})
	gp, l := Greps(msg)
	for i := 0; i < l; i++ {
		go SimpleReply(gp[i], msg)
	}
}

func GroupHeater(msg RecvMsg) {
	Action(msg, `Arisa --help`, func() {
		respGroupMsg(msg, "https://github.com/bLueriVerLHR/class2003bot\n"+`目前支持：
正则聊天
打招呼
发**图（仅在特定群有效）
讲苏联政治冷笑话（讽刺意味十足，慎重使用）
查mooc作业（主课和学委参加的课程）
hy语录
Arisa最近听什么
学习资料
点歌
名言
夸人
诅咒
`)
	})

	rand.Seed(int64(time.Now().Second()))
	if msg.RawMessage != lastMsg {
		lastMsg = msg.RawMessage
		repeat = 0
	} else {
		repeat++
		if repeat >= 3 {
			if rand.Int()%10 == 0 {
				if regexp.MustCompile(`\?{1.}|？{1.}`).Match([]byte(lastMsg)) {
					lastMsg = "[CQ:image,file=af6d521b897dfd5b2e5f82f080dab962.image]"
				}
				respGroupMsg(msg, lastMsg)
			}
		}
	}

	gp, l := Greps(msg)
	for i := 0; i < l; i++ {
		go SimpleReply(gp[i], msg)
	}
	/* 以下开始定义稍微复杂的功能 */

	Action(msg, `早安`, func() {
		loc, err := time.LoadLocation("Asia/Shanghai")
		CheckError(err)
		now := time.Now().In(loc).Hour()
		if !(now >= 3 && now <= 10) {
			respGroupMsg(msg, "goooo?")
		} else {
			respGroupMsg(msg, "[CQ:at,qq="+strconv.FormatInt(msg.UserID, 10)+"]早安，今天也是元气满满的一天呢！")
		}
	})
	Action(msg, `早上好|おはよう|ohayo|Good Morning|ohakusa|ohayayo|oha呀(呦|哟)|oha.*desu`, func() {
		loc, err := time.LoadLocation("Asia/Shanghai")
		CheckError(err)
		now := time.Now().In(loc).Hour()
		if !(now >= 3 && now <= 10) {
			respGroupMsg(msg, "goooo?")
		} else {
			respGroupMsg(msg, "[CQ:at,qq="+strconv.FormatInt(msg.UserID, 10)+"]Ohakusa!")
		}
	})
	Action(msg, `午安|中午好`, func() {
		loc, err := time.LoadLocation("Asia/Shanghai")
		CheckError(err)
		now := time.Now().In(loc).Hour()
		if !(now >= 11 && now <= 13) {
			respGroupMsg(msg, "goooo?")
		} else {
			respGroupMsg(msg, "[CQ:at,qq="+strconv.FormatInt(msg.UserID, 10)+"]呼，稍微有点困了呢。")
		}
	})
	Action(msg, `下午好`, func() {
		loc, err := time.LoadLocation("Asia/Shanghai")
		CheckError(err)
		now := time.Now().In(loc).Hour()
		if !(now >= 13 && now <= 17) {
			respGroupMsg(msg, "goooo?")
		} else {
			respGroupMsg(msg, "[CQ:at,qq="+strconv.FormatInt(msg.UserID, 10)+"]下午也要精神满满哦！")
		}
	})
	Action(msg, `晚上好`, func() {
		loc, err := time.LoadLocation("Asia/Shanghai")
		CheckError(err)
		now := time.Now().In(loc).Hour()
		if !(now >= 18 && now <= 20) {
			respGroupMsg(msg, "goooo?")
		} else {
			respGroupMsg(msg, "[CQ:at,qq="+strconv.FormatInt(msg.UserID, 10)+"]晚上好，记得不要熬夜，要早点睡哦。")
		}
	})
	Action(msg, `晚安`, func() {
		loc, err := time.LoadLocation("Asia/Shanghai")
		CheckError(err)
		now := time.Now().In(loc).Hour()
		if !(now >= 21 || now <= 2) {
			respGroupMsg(msg, "goooo?")
		} else if rand.Int()%3 == 0 {
			respGroupMsg(msg, "[CQ:at,qq="+strconv.FormatInt(msg.UserID, 10)+"]嘻嘻，一起睡觉吧！\n当……当然只是时间上一起。")
		} else {
			respGroupMsg(msg, "[CQ:at,qq="+strconv.FormatInt(msg.UserID, 10)+"]晚安，好好睡觉，不要看手机啦。")
		}
	})
	Action(msg, `--setu`, func() {
		var del []int
		st := SetuKoKo(msg)
		if len(st.Data) == 0 {
			return
		}
		res := respGroupMsg(msg, "涩图来啦[CQ:face,id=178]："+st.Data[0].Urls.Original)
		body, err := ioutil.ReadAll(res.Body)
		CheckError(err)
		var mi MsgId
		mi.Read(string(body))
		del = append(del, mi.Data.MessageID)

		res = respGroupMsg(msg, "作者："+st.Data[0].Author)
		body, err = ioutil.ReadAll(res.Body)
		CheckError(err)
		mi.Read(string(body))
		del = append(del, mi.Data.MessageID)

		if regexp.MustCompile(`r18`).Match([]byte(msg.RawMessage)) {
			res = respGroupMsg(msg, "[CQ:image,file=eced5efe93489fa63204cccec0156f84.image]")
			body, err = ioutil.ReadAll(res.Body)
			CheckError(err)
			mi.Read(string(body))
			del = append(del, mi.Data.MessageID)
		}

		go DeleteLater(del)
	})
	Action(msg, `--pixiv`, func() {
		pic, rank := SetuRank(msg)
		if rank == "" {
			rank = "1"
		}
		if len(pic.Data) == 0 {
			respGroupMsg(msg, "这个时间没有排行信息")
			return
		}
		RankNum, err := strconv.Atoi(rank)
		CheckError(err)
		if RankNum < 0 || RankNum > 29 {
			return
		}
		picture := pic.Data[RankNum-1]
		picId := pic.Data[RankNum-1].URL
		length := len(picId)
		for i := length - 1; i > 0; i-- {
			if picId[i] == '=' {
				picId = picId[i+1:]
				break
			}
		}
		if picture.Source == "pixiv_normal" {
			respGroupMsg(msg, "Title: "+picture.Title)
			respGroupMsg(msg, "[CQ:image,file="+"https://pixiv.cat/"+picId+".jpg"+"]")
		} else {
			respGroupMsg(msg, "pid: "+picId+"\nTitle: "+picture.Title)
		}
	})
	Action(msg, `Arisa来点苏联冷笑话|苏联冷笑话|苏联.*冷笑话`, func() {
		code := rand.Int() % 290
		respGroupMsg(msg, SovietJokes(code))
	})
	Action(msg, `Arisa.*MOOC.*|Arisa.*慕课.*|Arisa.*mooc.*`, func() {
		resp := FormatMOOC()
		respGroupMsg(msg, "作业如下：\n"+resp)
	})
	Action(msg, `hy：`, func() {
		flag := rand.Int() % 25
		hy(flag, msg)
	})
	Action(msg, `ss：`, func() {
		flag := rand.Int() % 40
		SomeoneSay(flag, msg)
	})
	Action(msg, `Arisa.*在听什么.*|Arisa今天听什么|Arisa最近听什么`, func() {
		MusicList := []string{
			"1855475139",
			"1840459406",
			"1459128908",
			"1409093516",
			"1803908863",
			"1858068739",
			"442867526",
			"32102655",
			"441491167",
			"1432456852",

			// NieR:Automata
			"468490569",
			"468490577",
			"468490592",
			"468490576",
			"468490601",
			"468490574",

			// Coldplay
			"1448421768",
		}
		length := len(MusicList)
		respGroupMsg(msg, "Arisa最近在听这首歌！")
		respGroupMsg(msg, "[CQ:music,type=163,id="+MusicList[rand.Int()%length]+"]")
	})
	Action(msg, `Arisa来点.*学习资料`, func() {
		switch true {
		case regexp.MustCompile(`Arisa.*(C|c|C语言|C\+\+|c\+\+|cpp|CPP)学习资料`).Match([]byte(msg.Message)):
			respGroupMsg(msg, "[CQ:share,url=https://zh.cppreference.com/,title=cppreference]")
			respGroupMsg(msg, "[CQ:share,url=http://www.cplusplus.com/,title=The C++ Resources Network]")
			respGroupMsg(msg, "[CQ:share,url=https://www.learncpp.com/,title=The C++ Tutorial | Learn C++]")
		case regexp.MustCompile(`Arisa.*(python|Python)学习资料`).Match([]byte(msg.Message)):
			respGroupMsg(msg, "[CQ:share,url=https://python.iswbm.com/,title=Python中文指南]")
		default:
			respGroupMsg(msg, "目前Arisa不支持这些内容")
		}
	})
	Action(msg, `Arisa点歌`, func() {
		if !regexp.MustCompile(`Arisa点歌\s.{1,}`).Match([]byte(msg.RawMessage)) {
			respGroupMsg(msg, "歌名哩？看不到啊？而点歌后面要加一个空格这样才可以和前面的内容区分哦！")
			return
		}
		fmt.Println(msg.RawMessage[12:])
		qm := QMsearch(msg.RawMessage[12:])
		if qm.Code != 0 || len(qm.Data.Song.List) == 0 {
			respGroupMsg(msg, "出错啦，找不到哩。")
			return
		}
		respGroupMsg(msg, "[CQ:music,type=qq,id="+strconv.Itoa(qm.Data.Song.List[0].ID)+"]")
	})
	Action(msg, `Arisa.*名言`, func() {
		ol := OneSentence(msg.RawMessage)
		if ol.UUID == "" {
			respGroupMsg(msg, "出错啦，找不到哩。")
			return
		}
		respGroupMsg(msg, ol.Hitokoto+"\nFrom "+ol.From+"\nFrom "+ol.FromWho+"\nCreator: "+ol.Creator)
	})
	Action(msg, `(Arisa|旅人|吕人|铝人|地鼠).*夸我|Arisa.*彩虹屁`, func() {
		respGroupMsg(msg, chp())
	})
	Action(msg, `Arisa诅咒`, func() {
		respGroupMsg(msg, nmsl())
	})
}

func Action(msg RecvMsg, grep string, f func()) {
	if regexp.MustCompile(grep).Match([]byte(msg.RawMessage)) {
		go f()
	}
}

func SimpleReply(gp grep, msg RecvMsg) {
	Action(msg, gp.reg, func() {
		respGroupMsg(msg, gp.resp)
	})
}

func respPrivateMsg(msg RecvMsg, text string) *http.Response {
	res, _ := http.PostForm(BOT_URL+"send_private_msg", url.Values{"user_id": {strconv.Itoa(int(msg.UserID))}, "message": {text}})
	return res
}

func respGroupMsg(msg RecvMsg, text string) *http.Response {
	res, _ := http.PostForm(BOT_URL+"send_group_msg", url.Values{"group_id": {strconv.Itoa(msg.GroupID)}, "message": {text}})
	return res
}

func respGroupNotice(groupID int, text string) *http.Response {
	res, _ := http.PostForm(BOT_URL+"send_group_msg", url.Values{"group_id": {strconv.Itoa(groupID)}, "message": {text}})
	return res
}

func DeleteLater(del []int) {
	time.Sleep(30 * time.Second)
	for i := 0; i < len(del); i++ {
		http.PostForm(BOT_URL+"delete_msg", url.Values{"message_id": {strconv.Itoa(del[i])}})
	}
}
