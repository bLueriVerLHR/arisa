package plugins

import (
	"arisa/tools"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type QQmusic struct {
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

func QMsearch(target string) QQmusic {
	response, err := http.Get("https://c.y.qq.com/soso/fcgi-bin/client_search_cp?ct=24&qqmusic_ver=1298&new_json=1&remoteplace=txt.yqq.song&searchid=&t=0&aggr=1&cr=1&catZhida=1&lossless=0&flag_qc=0&p=1&n=1&w=\"" + target + "\"")
	tools.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	tools.Check(err)
	var qm QQmusic
	length := len(body)
	err = json.Unmarshal(body[9:length-1], &qm)
	tools.Check(err)
	return qm
}
