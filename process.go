package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"arisa/plugins"
	"arisa/tools"

	"github.com/sirupsen/logrus"
)

type MsgId struct {
	Data struct {
		MessageID int `json:"message_id"`
	} `json:"data"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}

type RecvMsg struct {
	Anonymous   interface{} `json:"anonymous"`
	Font        int         `json:"font"`
	GroupID     int         `json:"group_id"`
	Message     string      `json:"message"`
	MessageID   int         `json:"message_id"`
	MessageSeq  int         `json:"message_seq"`
	MessageType string      `json:"message_type"`
	PostType    string      `json:"post_type"`
	RawMessage  string      `json:"raw_message"`
	SelfID      int64       `json:"self_id"`
	Sender      struct {
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Level    string `json:"level"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
		Sex      string `json:"sex"`
		Title    string `json:"title"`
		UserID   int64  `json:"user_id"`
	} `json:"sender"`
	SubType string `json:"sub_type"`
	Time    int    `json:"time"`
	UserID  int64  `json:"user_id"`
}

type BotStatus struct {
	Interval      int    `json:"interval"`
	MetaEventType string `json:"meta_event_type"`
	PostType      string `json:"post_type"`
	SelfID        int64  `json:"self_id"`
	Status        struct {
		AppEnabled     bool        `json:"app_enabled"`
		AppGood        bool        `json:"app_good"`
		AppInitialized bool        `json:"app_initialized"`
		Good           bool        `json:"good"`
		Online         bool        `json:"online"`
		PluginsGood    interface{} `json:"plugins_good"`
		Stat           struct {
			PacketReceived  int `json:"packet_received"`
			PacketSent      int `json:"packet_sent"`
			PacketLost      int `json:"packet_lost"`
			MessageReceived int `json:"message_received"`
			MessageSent     int `json:"message_sent"`
			DisconnectTimes int `json:"disconnect_times"`
			LostTimes       int `json:"lost_times"`
			LastMessageTime int `json:"last_message_time"`
		} `json:"stat"`
	} `json:"status"`
	Time int `json:"time"`
}

func HandleConn(resp http.ResponseWriter, req *http.Request) {
	MOOCSubscribe()
	reader, err := ioutil.ReadAll(req.Body)
	tools.Check(err)
	if recv := string(reader); tools.Grep(`"post_type":"meta_event"`, recv) {
		go HandleMetaEvent(recv)
	} else if tools.Grep(`"post_type":"message"`, recv) {
		go HandleMsg(recv)
	} else if tools.Grep(`"post_type":"notice"`, recv) {
		go HandleNotice(recv)
	}
	resp.WriteHeader(http.StatusOK)
	tools.Check(err)
}

func HandleMetaEvent(recv string) {
	var bs BotStatus
	err := json.Unmarshal([]byte(recv), &bs)
	tools.Check(err)
	if !bs.Status.Online {
		logrus.Errorln("Bot not online!")
	}
}

func HandleMsg(recv string) {
	var rm RecvMsg
	err := json.Unmarshal([]byte(recv), &rm)
	tools.Check(err)
	if rm.MessageType == "group" {
		HandleGroupMsg(rm)
	} else if rm.MessageType == "private" {
		HandlePrivateMsg(rm)
	}
}

func HandleNotice(recv string) {
	if tools.Grep(`"sub_type":"poke"`, recv) {
		RespPoke(recv)
	} else if tools.Grep(`"notice_type":"group_recall"`, recv) {
		AntiRecall(recv)
	}
}

func HandleGroupMsg(msg RecvMsg) {
	tools.ProtectRun(func() {
		ret := Help(msg.Message)
		if ret != "" {
			RespGroupMsg(msg, ret)
		}
	})
	RepeatSentence(msg)
	if GrepTalks(msg) {
		return
	}
	tools.ProtectRun(func() {
		ret := plugins.Greeting(msg.Message)
		if ret != "" {
			RespGroupMsg(msg, "[CQ:at,qq="+strconv.FormatInt(msg.UserID, 10)+"]"+ret)
		}
	})
	tools.ProtectRun(func() {
		ret := SomeSay(msg.Message)
		if ret != "" {
			RespGroupMsg(msg, ret)
		}
	})
	tools.ProtectRun(func() {
		Action(msg, `Arisa??????`, func() {
			if !regexp.MustCompile(`Arisa??????\s.{1,}`).Match([]byte(msg.RawMessage)) {
				RespGroupMsg(msg, "?????????????????????????????????????????????????????????????????????????????????????????????????????????")
				return
			}
			qm := plugins.QMsearch(msg.RawMessage[12:])
			if qm.Code != 0 || len(qm.Data.Song.List) == 0 {
				RespGroupMsg(msg, "???????????????????????????")
				return
			}
			RespGroupMsg(msg, "[CQ:music,type=qq,id="+strconv.Itoa(qm.Data.Song.List[0].ID)+"]")
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `Arisa.*??????`, func() {
			ol := plugins.OneSentence(msg.RawMessage)
			if ol.UUID == "" {
				RespGroupMsg(msg, "???????????????????????????")
				return
			}
			RespGroupMsg(msg, ol.Hitokoto+"\nFrom "+ol.From+"\nFrom "+ol.FromWho+"\nCreator: "+ol.Creator)
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `Arisa.*???.|Arisa.*?????????`, func() {
			RespGroupMsg(msg, plugins.Chp())
		})
		Action(msg, `Arisa??????`, func() {
			if IsWhite(msg.GroupID, uni.Conf.BadGirl.White) {
				RespGroupMsg(msg, plugins.Nmsl())
			}
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `--setu`, func() {
			st := plugins.SetuKoKo(msg.Message)
			if len(st.Data) == 0 {
				return
			}
			if st.Data[0].R18 && !IsWhite(msg.GroupID, uni.Conf.Setu.White) {
				RespGroupMsg(msg, "???????????????R18???????????????????????????")
				return
			}
			RespGroupMsg(msg, "????????????[CQ:face,id=178]???"+st.Data[0].Urls.Original)
			if !st.Data[0].R18 {
				res := RespGroupMsg(msg, "[CQ:image,file="+st.Data[0].Urls.Original+"]")
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				tools.Check(err)
				var del MsgId
				err = json.Unmarshal(body, &del)
				tools.Check(err)
				go DeleteLater(del.Data.MessageID, 30)
				return
			}
			RespGroupMsg(msg, "????????????")
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `--pixiv`, func() {
			pic, rank := plugins.SetuRank(msg.Message)
			if rank == "" {
				rank = "1"
			}
			if len(pic.Data) == 0 {
				RespGroupMsg(msg, "??????????????????????????????")
				return
			}
			RankNum, err := strconv.Atoi(rank)
			tools.Check(err)
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
			if picture.Source != "pixiv_male_r18" {
				RespGroupMsg(msg, "Title: "+picture.Title)
				RespGroupMsg(msg, "[CQ:image,file="+"https://pixiv.cat/"+picId+".jpg"+"]")
			} else if IsWhite(msg.GroupID, uni.Conf.Pixiv.White) {
				RespGroupMsg(msg, "pid: "+picId+"\nTitle: "+picture.Title)
			} else {
				RespGroupMsg(msg, "???????????????R18???????????????????????????")
			}
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `<--.{1,}-->`, func() {
			RunJS(msg)
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `?????????`, func() {
			res := RandEvent(msg)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			tools.Check(err)
			var del MsgId
			err = json.Unmarshal(body, &del)
			tools.Check(err)
			go DeleteLater(del.Data.MessageID, 30)
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `mycoser`, func() {
			if ret := plugins.MyCoser(msg.Message); ret != "" {
				RespGroupMsg(msg, "[CQ:image,file="+ret+"]")
			}
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `?????????|bcy`, func() {
			BcyCPN(msg)
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `?????????`, func() {
			RespGroupMsg(msg, "[CQ:share,url=https://zh.moegirl.org.cn/"+msg.Message[9:]+",title="+msg.Message[9:]+"]")
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `?????????`, func() {
			RespGroupMsg(msg, "[CQ:share,url=https://baike.baidu.com/item/"+msg.Message[9:]+",title="+msg.Message[9:]+"]")
		})
		Action(msg, `^Arisa.{1,}?????????$`, func() {
			RespGroupMsg(msg, "[CQ:share,url=https://baike.baidu.com/item/"+msg.Message[5:len(msg.Message)-9]+",title="+msg.Message[5:len(msg.Message)-9]+"]")
		})
		Action(msg, `^Arisa.{1,}????????????$`, func() {
			RespGroupMsg(msg, "[CQ:share,url=https://baike.baidu.com/item/"+msg.Message[5:len(msg.Message)-12]+",title="+msg.Message[5:len(msg.Message)-12]+"]")
		})
		Action(msg, `^Arisa???(??????|???)`, func() {
			RespGroupMsg(msg, "[CQ:share,url=https://zh.moegirl.org.cn/%E4%BA%9A%E9%87%8C%E8%8E%8E(%E5%85%AC%E4%B8%BB%E8%BF%9E%E7%BB%93)#,title=?????????]")
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `Arisa.*??????`, func() {
			var NewType string
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "0"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "1"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "2"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "3"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "4"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "5"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "6"
			}
			if NewType == "" {
				NewType = "7"
			}
			RespGroupMsg(msg, plugins.IsoyuN(NewType))
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `Arisa????????????`, func() {
			RespGroupMsg(msg, plugins.IsoyuLocal())
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `isoyucos`, func() {
			res := plugins.IsoyuCos()
			if length := len(res); length != 0 {
				RespGroupMsg(msg, res[0])
				for i := 1; i < length; i++ {
					RespGroupMsg(msg, "[CQ:image,file="+res[i]+"]")
				}
			} else {
				RespGroupMsg(msg, "?????????????????????")
			}
		})
	})
}

func HandlePrivateMsg(msg RecvMsg) {
	tools.ProtectRun(func() {
		ret := Help(msg.Message)
		if ret != "" {
			RespPrivateMsg(msg, ret)
		}
	})
	tools.ProtectRun(func() {
		Action(msg, `mooc`, func() {
			AntiMOOC(msg)
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `<--.{1,}-->`, func() {
			RunJS(msg)
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `Arisa admin`, func() {
			if !(msg.UserID == uni.Conf.Admin) {
				return
			}
			Settings(msg)
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `Arisa??????`, func() {
			var NewType string
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "0"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "1"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "2"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "3"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "4"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "5"
			}
			if regexp.MustCompile(`??????`).Match([]byte(msg.Message)) {
				NewType = "6"
			}
			if NewType == "" {
				NewType = "7"
			}
			RespPrivateMsg(msg, plugins.IsoyuN(NewType))
		})
	})
	tools.ProtectRun(func() {
		Action(msg, `mycoser`, func() {
			if ret := plugins.MyCoser(msg.Message); ret != "" {
				RespPrivateMsg(msg, "[CQ:image,file="+ret+"]")
			}
		})
	})
}

func Action(msg RecvMsg, grep string, f func()) {
	if regexp.MustCompile(grep).Match([]byte(msg.RawMessage)) {
		f()
	}
}
