package main

import (
	"arisa/tools"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func RespPrivateMsg(msg RecvMsg, text string) *http.Response {
	res, err := http.PostForm("http://"+uni.Conf.BotConfig.Host+":"+strconv.Itoa(uni.Conf.BotConfig.Port)+"/send_private_msg", url.Values{"user_id": {strconv.Itoa(int(msg.UserID))}, "message": {text}})
	tools.Check(err)
	return res
}

func RespGroupMsg(msg RecvMsg, text string) *http.Response {
	res, err := http.PostForm("http://"+uni.Conf.BotConfig.Host+":"+strconv.Itoa(uni.Conf.BotConfig.Port)+"/send_group_msg", url.Values{"group_id": {strconv.Itoa(msg.GroupID)}, "message": {text}})
	tools.Check(err)
	return res
}

func RespPrivateNotice(senderID int, text string) *http.Response {
	res, err := http.PostForm("http://"+uni.Conf.BotConfig.Host+":"+strconv.Itoa(uni.Conf.BotConfig.Port)+"/send_private_msg", url.Values{"user_id": {strconv.Itoa(senderID)}, "message": {text}})
	tools.Check(err)
	return res
}

func RespGroupNotice(groupID int, text string) *http.Response {
	res, err := http.PostForm("http://"+uni.Conf.BotConfig.Host+":"+strconv.Itoa(uni.Conf.BotConfig.Port)+"/send_group_msg", url.Values{"group_id": {strconv.Itoa(groupID)}, "message": {text}})
	tools.Check(err)
	return res
}

func DeleteLater(DelId int, SetTime int64) *http.Response {
	time.Sleep(time.Duration(SetTime) * time.Second)
	res, err := http.PostForm("http://"+uni.Conf.BotConfig.Host+":"+strconv.Itoa(uni.Conf.BotConfig.Port)+"/delete_msg", url.Values{"message_id": {strconv.Itoa(DelId)}})
	tools.Check(err)
	return res
}
