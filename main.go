package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	HttpResponseHeader = "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n"
	ReportAddress      = ":5710"
	BOT_URL            = "http://0.0.0.0:5700/"
	SETU_URL           = "https://api.lolicon.app/setu/v2"
	JokePath           = "./jokes"
	cookiePath         = "./cookie"
	csrfKey            = "ec59fdc57f434187bc5b293d9a69f7d5"
	version            = "1"
	gallery            = "https://gallery.nyadora.moe/"
	setuAllow          = "740228215"
)

type grep struct {
	reg  string
	resp string
}

type Mooc struct {
	lessonId  string
	termId    string
	className string
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

type Setu struct {
	Error string `json:"error"`
	Data  []struct {
		Pid        int      `json:"pid"`
		P          int      `json:"p"`
		R18        bool     `json:"r18"`
		UID        int      `json:"uid"`
		Title      string   `json:"title"`
		Author     string   `json:"author"`
		Tags       []string `json:"tags"`
		Width      int      `json:"width"`
		Height     int      `json:"height"`
		Ext        string   `json:"ext"`
		UploadDate int64    `json:"uploadDate"`
		Urls       struct {
			Original string `json:"original"`
		} `json:"urls"`
	} `json:"data"`
}

type MOOC struct {
	ID              int         `json:"id"`
	GmtCreate       int64       `json:"gmtCreate"`
	GmtModified     int64       `json:"gmtModified"`
	Name            string      `json:"name"`
	Position        int         `json:"position"`
	TermID          int         `json:"termId"`
	ChapterID       int         `json:"chapterId"`
	ContentType     int         `json:"contentType"`
	ContentID       int         `json:"contentId"`
	IsTestChecked   bool        `json:"isTestChecked"`
	Units           interface{} `json:"units"`
	ReleaseTime     int64       `json:"releaseTime"`
	ViewStatus      int         `json:"viewStatus"`
	TestDraftStatus int         `json:"testDraftStatus"`
	Test            struct {
		ID                       int         `json:"id"`
		ReleaseTime              int64       `json:"releaseTime"`
		Type                     int         `json:"type"`
		Name                     string      `json:"name"`
		Deadline                 int64       `json:"deadline"`
		TestTime                 interface{} `json:"testTime"`
		Trytime                  interface{} `json:"trytime"`
		UsedTryCount             int         `json:"usedTryCount"`
		EvaluateJudgeType        int         `json:"evaluateJudgeType"`
		EvaluateNeedTrain        int         `json:"evaluateNeedTrain"`
		EvaluateStart            int64       `json:"evaluateStart"`
		EvaluateEnd              int64       `json:"evaluateEnd"`
		EvaluateScoreReleaseTime int64       `json:"evaluateScoreReleaseTime"`
		ScorePubStatus           int         `json:"scorePubStatus"`
		EnableEvaluation         bool        `json:"enableEvaluation"`
		UserScore                interface{} `json:"userScore"`
		TotalScore               float64     `json:"totalScore"`
		BonusScore               interface{} `json:"bonusScore"`
		ExamID                   int         `json:"examId"`
	} `json:"test"`
}

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

type MsgId struct {
	Data struct {
		MessageID int `json:"message_id"`
	} `json:"data"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}

type Pixiv struct {
	Code int `json:"code"`
	Data []struct {
		ID        int         `json:"id"`
		Source    string      `json:"source"`
		URL       string      `json:"url"`
		Thumb     string      `json:"thumb"`
		Rank      int         `json:"rank"`
		Date      string      `json:"date"`
		CreatedAt string      `json:"created_at"`
		UpdatedAt interface{} `json:"updated_at"`
		Title     string      `json:"title"`
	} `json:"data"`
}

type qqmusic struct {
	Code int `json:"code"`
	Data struct {
		Keyword  string        `json:"keyword"`
		Priority int           `json:"priority"`
		Qc       []interface{} `json:"qc"`
		Semantic struct {
			Curnum   int           `json:"curnum"`
			Curpage  int           `json:"curpage"`
			List     []interface{} `json:"list"`
			Totalnum int           `json:"totalnum"`
		} `json:"semantic"`
		Song struct {
			Curnum  int `json:"curnum"`
			Curpage int `json:"curpage"`
			List    []struct {
				Act    int `json:"act"`
				Action struct {
					Alert    int `json:"alert"`
					Icons    int `json:"icons"`
					Msgdown  int `json:"msgdown"`
					Msgfav   int `json:"msgfav"`
					Msgid    int `json:"msgid"`
					Msgpay   int `json:"msgpay"`
					Msgshare int `json:"msgshare"`
					Switch   int `json:"switch"`
				} `json:"action"`
				Album struct {
					ID             int    `json:"id"`
					Mid            string `json:"mid"`
					Name           string `json:"name"`
					Pmid           string `json:"pmid"`
					Subtitle       string `json:"subtitle"`
					TimePublic     string `json:"time_public"`
					Title          string `json:"title"`
					TitleHighlight string `json:"title_highlight"`
				} `json:"album"`
				Bpm         int    `json:"bpm"`
				Content     string `json:"content"`
				Desc        string `json:"desc"`
				DescHilight string `json:"desc_hilight"`
				Docid       string `json:"docid"`
				Eq          int    `json:"eq"`
				Es          string `json:"es"`
				File        struct {
					B30S          int    `json:"b_30s"`
					E30S          int    `json:"e_30s"`
					HiresBitdepth int    `json:"hires_bitdepth"`
					HiresSample   int    `json:"hires_sample"`
					MediaMid      string `json:"media_mid"`
					Size128       int    `json:"size_128"`
					Size128Mp3    int    `json:"size_128mp3"`
					Size192Aac    int    `json:"size_192aac"`
					Size192Ogg    int    `json:"size_192ogg"`
					Size24Aac     int    `json:"size_24aac"`
					Size320       int    `json:"size_320"`
					Size320Mp3    int    `json:"size_320mp3"`
					Size48Aac     int    `json:"size_48aac"`
					Size96Aac     int    `json:"size_96aac"`
					Size96Ogg     int    `json:"size_96ogg"`
					SizeAac       int    `json:"size_aac"`
					SizeApe       int    `json:"size_ape"`
					SizeDts       int    `json:"size_dts"`
					SizeFlac      int    `json:"size_flac"`
					SizeHires     int    `json:"size_hires"`
					SizeOgg       int    `json:"size_ogg"`
					SizeTry       int    `json:"size_try"`
					StrMediaMid   string `json:"strMediaMid"`
					TryBegin      int    `json:"try_begin"`
					TryEnd        int    `json:"try_end"`
					URL           string `json:"url"`
				} `json:"file"`
				Fnote      int           `json:"fnote"`
				Genre      int           `json:"genre"`
				Grp        []interface{} `json:"grp"`
				Href3      string        `json:"href3"`
				ID         int           `json:"id"`
				IndexAlbum int           `json:"index_album"`
				IndexCd    int           `json:"index_cd"`
				Interval   int           `json:"interval"`
				Isonly     int           `json:"isonly"`
				Ksong      struct {
					ID  int    `json:"id"`
					Mid string `json:"mid"`
				} `json:"ksong"`
				Label        string `json:"label"`
				Language     int    `json:"language"`
				Lyric        string `json:"lyric"`
				LyricHilight string `json:"lyric_hilight"`
				Mid          string `json:"mid"`
				Mv           struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					Title string `json:"title"`
					Vid   string `json:"vid"`
					Vt    int    `json:"vt"`
				} `json:"mv"`
				Name      string `json:"name"`
				NewStatus int    `json:"newStatus"`
				Ov        int    `json:"ov"`
				Pay       struct {
					PayDown    int `json:"pay_down"`
					PayMonth   int `json:"pay_month"`
					PayPlay    int `json:"pay_play"`
					PayStatus  int `json:"pay_status"`
					PriceAlbum int `json:"price_album"`
					PriceTrack int `json:"price_track"`
					TimeFree   int `json:"time_free"`
				} `json:"pay"`
				Protect int `json:"protect"`
				Sa      int `json:"sa"`
				Singer  []struct {
					ID             int    `json:"id"`
					Mid            string `json:"mid"`
					Name           string `json:"name"`
					Pmid           string `json:"pmid"`
					Title          string `json:"title"`
					TitleHighlight string `json:"title_highlight"`
					Type           int    `json:"type"`
					Uin            int    `json:"uin"`
				} `json:"singer"`
				Status       int    `json:"status"`
				Subtitle     string `json:"subtitle"`
				Tag          int    `json:"tag"`
				Tid          int    `json:"tid"`
				TimePublic   string `json:"time_public"`
				Title        string `json:"title"`
				TitleHilight string `json:"title_hilight"`
				Type         int    `json:"type"`
				URL          string `json:"url"`
				Version      int    `json:"version"`
				Volume       struct {
					Gain float64 `json:"gain"`
					Lra  float64 `json:"lra"`
					Peak float64 `json:"peak"`
				} `json:"volume"`
			} `json:"list"`
			Totalnum int `json:"totalnum"`
		} `json:"song"`
		Tab       int           `json:"tab"`
		Taglist   []interface{} `json:"taglist"`
		Totaltime int           `json:"totaltime"`
		Zhida     struct {
			Type        int `json:"type"`
			ZhidaSinger struct {
				AlbumNum          int           `json:"albumNum"`
				Hotalbum          []interface{} `json:"hotalbum"`
				Hotsong           []interface{} `json:"hotsong"`
				MvNum             int           `json:"mvNum"`
				SingerID          int           `json:"singerID"`
				SingerMID         string        `json:"singerMID"`
				SingerName        string        `json:"singerName"`
				SingerPic         string        `json:"singerPic"`
				SingernameHilight string        `json:"singername_hilight"`
				SongNum           int           `json:"songNum"`
			} `json:"zhida_singer"`
		} `json:"zhida"`
	} `json:"data"`
	Message string `json:"message"`
	Notice  string `json:"notice"`
	Subcode int    `json:"subcode"`
	Time    int    `json:"time"`
	Tips    string `json:"tips"`
}

type OneLang struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

var repeat int
var lastMsg string

func HandleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [2048]byte
	n, err := reader.Read(buf[:])
	CheckError(err)

	recv := string(buf[:n])

	// fmt.Println(recv)

	Type := FindTargetAttibute(recv, "post_type")
	if Type == "meta_event" {
		go HandleMetaInfo(recv)
	} else if Type == "message" {
		go HandleMsg(recv)
	} else if Type == "notice" {
		subType := FindTargetAttibute(recv, "sub_type")
		if subType == "poke" {
			go HandlePoke(recv)
		}
	} else {
		fmt.Println(recv)
	}
	_, err = conn.Write([]byte(HttpResponseHeader))
	CheckError(err)
}

func main() {
	ls, err := net.Listen("tcp", ReportAddress)
	CheckError(err)
	for {
		conn, err := ls.Accept()
		CheckError(err)
		go HandleConn(conn)
	}
}
