package plugins

import (
	"arisa/tools"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BcyCoser struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TopListItemInfo []struct {
			TlType     string `json:"tl_type"`
			Since      string `json:"since"`
			ItemDetail struct {
				ItemID        string `json:"item_id"`
				UID           int64  `json:"uid"`
				Uname         string `json:"uname"`
				Avatar        string `json:"avatar"`
				OdinUID       string `json:"odin_uid"`
				ValueUser     int    `json:"value_user"`
				VuDescription string `json:"vu_description"`
				FollowState   string `json:"follow_state"`
				Rights        []struct {
					ID                 int    `json:"id"`
					Rid                int    `json:"rid"`
					DisplayName        string `json:"display_name"`
					Type               int    `json:"type"`
					ExpireTime         int64  `json:"expire_time"`
					Extra              string `json:"extra"`
					Status             int    `json:"status"`
					Description        string `json:"description"`
					Link               string `json:"link"`
					LinkTitle          string `json:"link_title"`
					AvailableStartTime int    `json:"available_start_time"`
					AvailableEndTime   int    `json:"available_end_time"`
					Active             bool   `json:"active"`
					Own                bool   `json:"own"`
				} `json:"rights"`
				Ctime     int    `json:"ctime"`
				Type      string `json:"type"`
				Plain     string `json:"plain"`
				WordCount int    `json:"word_count"`
				Cover     string `json:"cover"`
				Multi     []struct {
					Path         string  `json:"path"`
					Type         string  `json:"type"`
					Mid          int     `json:"mid"`
					W            int     `json:"w"`
					H            int     `json:"h"`
					OriginalPath string  `json:"original_path"`
					Ratio        float64 `json:"ratio"`
					Format       string  `json:"format"`
					VisibleLevel string  `json:"visible_level"`
					Origin       string  `json:"origin"`
				} `json:"multi"`
				PicNum    int    `json:"pic_num"`
				Work      string `json:"work"`
				Wid       int    `json:"wid"`
				RealName  string `json:"real_name"`
				WorkCover string `json:"work_cover"`
				PostTags  []struct {
					TagID       int    `json:"tag_id"`
					TagName     string `json:"tag_name"`
					Type        string `json:"type"`
					Cover       string `json:"cover"`
					RelativeWid int    `json:"relative_wid"`
				} `json:"post_tags"`
				LikeCount  int  `json:"like_count"`
				UserLiked  bool `json:"user_liked"`
				ReplyCount int  `json:"reply_count"`
				ShareCount int  `json:"share_count"`
				Props      []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"props"`
				Replies      []interface{} `json:"replies"`
				VisibleLevel int           `json:"visible_level"`
				UserFavored  bool          `json:"user_favored"`
				ImageList    []struct {
					Path         string  `json:"path"`
					Type         string  `json:"type"`
					Mid          int     `json:"mid"`
					W            int     `json:"w"`
					H            int     `json:"h"`
					OriginalPath string  `json:"original_path"`
					Ratio        float64 `json:"ratio"`
					Format       string  `json:"format"`
					VisibleLevel string  `json:"visible_level"`
					Origin       string  `json:"origin"`
				} `json:"image_list"`
				TopListDetail struct {
					Rank        int    `json:"rank"`
					TopListName string `json:"top_list_name"`
					BcyURL      string `json:"bcy_url"`
				} `json:"top_list_detail"`
				ExtraProperties struct {
					ItemReplyDisable bool `json:"item_reply_disable"`
				} `json:"extra_properties"`
				SelectedStatus  int    `json:"selected_status"`
				SelectedComment string `json:"selected_comment"`
				EditorStatus    string `json:"editor_status"`
				PostInSet       bool   `json:"post_in_set"`
				ViewCount       int    `json:"view_count"`
				SetData         struct {
					Title        string `json:"title"`
					SetPostPrev  int    `json:"set_post_prev"`
					SetPostNext  int    `json:"set_post_next"`
					PostPos      int    `json:"post_pos"`
					Count        int    `json:"count"`
					Subscribed   bool   `json:"subscribed"`
					SubscribeNum int    `json:"subscribe_num"`
					ItemSetID    int    `json:"item_set_id"`
				} `json:"set_data"`
				Repostable  bool `json:"repostable"`
				RepostCount int  `json:"repost_count"`
				Collection  struct {
					Title        string `json:"title"`
					CollectionID int    `json:"collection_id"`
					User         struct {
						UID int64 `json:"uid"`
					} `json:"user"`
					CollectionType string `json:"collection_type"`
				} `json:"collection"`
				VisibleStatus    int    `json:"visible_status"`
				VisibleStatusMsg string `json:"visible_status_msg"`
				Mtime            int    `json:"mtime"`
			} `json:"item_detail"`
			TopListDetail struct {
				TtypeSet struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"ttype_set"`
				SubTypeSet struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"sub_type_set"`
				Stime int `json:"stime"`
				Count int `json:"count"`
				Rank  int `json:"rank"`
			} `json:"top_list_detail"`
		} `json:"top_list_item_info"`
	} `json:"data"`
}

type BcyIllust struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TopListItemInfo []struct {
			TlType     string `json:"tl_type"`
			Since      string `json:"since"`
			ItemDetail struct {
				ItemID        string        `json:"item_id"`
				UID           int64         `json:"uid"`
				Uname         string        `json:"uname"`
				Avatar        string        `json:"avatar"`
				OdinUID       string        `json:"odin_uid"`
				ValueUser     int           `json:"value_user"`
				VuDescription string        `json:"vu_description"`
				FollowState   string        `json:"follow_state"`
				Rights        []interface{} `json:"rights"`
				Ctime         int           `json:"ctime"`
				Type          string        `json:"type"`
				Plain         string        `json:"plain"`
				WordCount     int           `json:"word_count"`
				Cover         string        `json:"cover"`
				Multi         []struct {
					Path         string `json:"path"`
					Type         string `json:"type"`
					Mid          int    `json:"mid"`
					W            int    `json:"w"`
					H            int    `json:"h"`
					OriginalPath string `json:"original_path"`
					Ratio        int    `json:"ratio"`
					Format       string `json:"format"`
					VisibleLevel string `json:"visible_level"`
					Origin       string `json:"origin"`
				} `json:"multi"`
				PicNum    int    `json:"pic_num"`
				Work      string `json:"work"`
				Wid       int    `json:"wid"`
				RealName  string `json:"real_name"`
				WorkCover string `json:"work_cover"`
				PostTags  []struct {
					TagID       int    `json:"tag_id"`
					TagName     string `json:"tag_name"`
					Type        string `json:"type"`
					Cover       string `json:"cover"`
					RelativeWid int    `json:"relative_wid"`
				} `json:"post_tags"`
				LikeCount  int  `json:"like_count"`
				UserLiked  bool `json:"user_liked"`
				ReplyCount int  `json:"reply_count"`
				ShareCount int  `json:"share_count"`
				Props      []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"props"`
				Replies      []interface{} `json:"replies"`
				VisibleLevel int           `json:"visible_level"`
				UserFavored  bool          `json:"user_favored"`
				ImageList    []struct {
					Path         string  `json:"path"`
					Type         string  `json:"type"`
					Mid          int     `json:"mid"`
					W            int     `json:"w"`
					H            int     `json:"h"`
					OriginalPath string  `json:"original_path"`
					Ratio        float64 `json:"ratio"`
					Format       string  `json:"format"`
					VisibleLevel string  `json:"visible_level"`
					Origin       string  `json:"origin"`
				} `json:"image_list"`
				TopListDetail struct {
					Rank        int    `json:"rank"`
					TopListName string `json:"top_list_name"`
					BcyURL      string `json:"bcy_url"`
				} `json:"top_list_detail"`
				ExtraProperties struct {
					ItemReplyDisable bool `json:"item_reply_disable"`
				} `json:"extra_properties"`
				SelectedStatus  int    `json:"selected_status"`
				SelectedComment string `json:"selected_comment"`
				EditorStatus    string `json:"editor_status"`
				PostInSet       bool   `json:"post_in_set"`
				ViewCount       int    `json:"view_count"`
				SetData         struct {
					Title        string `json:"title"`
					SetPostPrev  int    `json:"set_post_prev"`
					SetPostNext  int    `json:"set_post_next"`
					PostPos      int    `json:"post_pos"`
					Count        int    `json:"count"`
					Subscribed   bool   `json:"subscribed"`
					SubscribeNum int    `json:"subscribe_num"`
					ItemSetID    int    `json:"item_set_id"`
				} `json:"set_data"`
				Repostable  bool `json:"repostable"`
				RepostCount int  `json:"repost_count"`
				Collection  struct {
					Title        string `json:"title"`
					CollectionID int    `json:"collection_id"`
					User         struct {
						UID int64 `json:"uid"`
					} `json:"user"`
					CollectionType string `json:"collection_type"`
				} `json:"collection"`
				VisibleStatus    int    `json:"visible_status"`
				VisibleStatusMsg string `json:"visible_status_msg"`
				Mtime            int    `json:"mtime"`
			} `json:"item_detail"`
			TopListDetail struct {
				TtypeSet struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"ttype_set"`
				SubTypeSet struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"sub_type_set"`
				Stime int `json:"stime"`
				Count int `json:"count"`
				Rank  int `json:"rank"`
			} `json:"top_list_detail"`
		} `json:"top_list_item_info"`
	} `json:"data"`
}

type BcyNovel struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TopListItemInfo []struct {
			TlType     string `json:"tl_type"`
			Since      string `json:"since"`
			ItemDetail struct {
				ItemID        string `json:"item_id"`
				UID           int    `json:"uid"`
				Uname         string `json:"uname"`
				Avatar        string `json:"avatar"`
				OdinUID       string `json:"odin_uid"`
				ValueUser     int    `json:"value_user"`
				VuDescription string `json:"vu_description"`
				FollowState   string `json:"follow_state"`
				Rights        []struct {
					ID                 int    `json:"id"`
					Rid                int    `json:"rid"`
					DisplayName        string `json:"display_name"`
					Type               int    `json:"type"`
					ExpireTime         int    `json:"expire_time"`
					Extra              string `json:"extra"`
					Status             int    `json:"status"`
					Description        string `json:"description"`
					Link               string `json:"link"`
					LinkTitle          string `json:"link_title"`
					AvailableStartTime int    `json:"available_start_time"`
					AvailableEndTime   int    `json:"available_end_time"`
					Active             bool   `json:"active"`
					Own                bool   `json:"own"`
				} `json:"rights"`
				Ctime     int           `json:"ctime"`
				Type      string        `json:"type"`
				Title     string        `json:"title"`
				Summary   string        `json:"summary"`
				Content   string        `json:"content"`
				Plain     string        `json:"plain"`
				WordCount int           `json:"word_count"`
				Cover     string        `json:"cover"`
				Multi     []interface{} `json:"multi"`
				PicNum    int           `json:"pic_num"`
				Work      string        `json:"work"`
				Wid       int           `json:"wid"`
				RealName  string        `json:"real_name"`
				WorkCover string        `json:"work_cover"`
				PostTags  []struct {
					TagID       int    `json:"tag_id"`
					TagName     string `json:"tag_name"`
					Type        string `json:"type"`
					Cover       string `json:"cover"`
					EventID     int    `json:"event_id,omitempty"`
					EventEnd    bool   `json:"event_end,omitempty"`
					RelativeWid int    `json:"relative_wid"`
				} `json:"post_tags"`
				LikeCount  int  `json:"like_count"`
				UserLiked  bool `json:"user_liked"`
				ReplyCount int  `json:"reply_count"`
				ShareCount int  `json:"share_count"`
				Props      []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"props"`
				Replies       []interface{} `json:"replies"`
				VisibleLevel  int           `json:"visible_level"`
				UserFavored   bool          `json:"user_favored"`
				ImageList     []interface{} `json:"image_list"`
				TopListDetail struct {
					Rank        int    `json:"rank"`
					TopListName string `json:"top_list_name"`
					BcyURL      string `json:"bcy_url"`
				} `json:"top_list_detail"`
				ExtraProperties struct {
					ItemReplyDisable bool `json:"item_reply_disable"`
				} `json:"extra_properties"`
				SelectedStatus  int    `json:"selected_status"`
				SelectedComment string `json:"selected_comment"`
				EditorStatus    string `json:"editor_status"`
				PostInSet       bool   `json:"post_in_set"`
				ViewCount       int    `json:"view_count"`
				SetData         struct {
					Title        string `json:"title"`
					SetPostPrev  int    `json:"set_post_prev"`
					SetPostNext  int    `json:"set_post_next"`
					PostPos      int    `json:"post_pos"`
					Count        int    `json:"count"`
					Subscribed   bool   `json:"subscribed"`
					SubscribeNum int    `json:"subscribe_num"`
					ItemSetID    int    `json:"item_set_id"`
				} `json:"set_data"`
				Repostable  bool `json:"repostable"`
				RepostCount int  `json:"repost_count"`
				Collection  struct {
					Title        string `json:"title"`
					CollectionID int    `json:"collection_id"`
					User         struct {
						UID int `json:"uid"`
					} `json:"user"`
					CollectionType string `json:"collection_type"`
				} `json:"collection"`
				VisibleStatus    int    `json:"visible_status"`
				VisibleStatusMsg string `json:"visible_status_msg"`
				Mtime            int    `json:"mtime"`
			} `json:"item_detail"`
			TopListDetail struct {
				TtypeSet struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"ttype_set"`
				SubTypeSet struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"sub_type_set"`
				Stime int    `json:"stime"`
				Count int    `json:"count"`
				Rank  int    `json:"rank"`
				Wave  string `json:"wave"`
			} `json:"top_list_detail"`
		} `json:"top_list_item_info"`
	} `json:"data"`
}

func Bcy(ttype string, p string, sub_type string, date string, rank int) []string {
	fmt.Println(ttype, p, sub_type, date, rank)
	resp, err := http.Get("https://bcy.net/apiv3/rank/list/itemInfo?p=" + p + "&ttype=" + ttype + "&sub_type=" + sub_type + "&date=" + date)
	tools.Check(err)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
	tools.Check(err)
	if ttype == "cos" {
		var image_list []string
		var ag BcyCoser
		err = json.Unmarshal(body, &ag)
		tools.Check(err)
		if rank >= len(ag.Data.TopListItemInfo) {
			rank = len(ag.Data.TopListItemInfo) - 1
		}
		if rank > 0 {
			length := len(ag.Data.TopListItemInfo[rank].ItemDetail.ImageList)
			image_list = append(image_list, "地址："+"https://bcy.net/item/detail/"+ag.Data.TopListItemInfo[rank].ItemDetail.ItemID)
			for i := 0; i < length; i++ {
				image_list = append(image_list, ag.Data.TopListItemInfo[rank].ItemDetail.ImageList[i].Path+"~tplv-banciyuan-w650.image")
			}
			return image_list
		}
	} else if ttype == "illust" {
		var image_list []string
		var ag BcyCoser
		err = json.Unmarshal(body, &ag)
		tools.Check(err)
		if rank >= len(ag.Data.TopListItemInfo) {
			rank = len(ag.Data.TopListItemInfo) - 1
		}
		if rank > 0 {
			length := len(ag.Data.TopListItemInfo[rank].ItemDetail.ImageList)
			image_list = append(image_list, "地址："+"https://bcy.net/item/detail/"+ag.Data.TopListItemInfo[rank].ItemDetail.ItemID)
			for i := 0; i < length; i++ {
				image_list = append(image_list, ag.Data.TopListItemInfo[rank].ItemDetail.ImageList[i].Path+"~tplv-banciyuan-w650.image")
			}
			return image_list
		}
	} else if ttype == "novel" {
		var description string
		var ag BcyNovel
		err = json.Unmarshal(body, &ag)
		tools.Check(err)
		if rank >= len(ag.Data.TopListItemInfo) {
			rank = len(ag.Data.TopListItemInfo) - 1
		}
		if rank > 0 {
			description = "标题：" + ag.Data.TopListItemInfo[rank].ItemDetail.Title + "\n\n摘要：" + ag.Data.TopListItemInfo[rank].ItemDetail.Content + "\n\n链接：" + "https://bcy.net/item/detail/" + ag.Data.TopListItemInfo[rank].ItemDetail.ItemID
		}
		return []string{description}
	}
	return []string{}
}
