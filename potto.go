package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/robertkrimen/otto"
)

func RunJS(msg RecvMsg) {
	if !Permitted(msg.UserID) {
		if msg.MessageType == "group" {
			RespGroupMsg(msg, "Not permitted.")
		} else if msg.MessageType == "private" {
			RespPrivateMsg(msg, "Not permitted.")
		}
		return
	}
	res := regexp.MustCompile(`<--.{1,}-->`).Find([]byte(msg.Message))
	script := msg.Message[len(res):]
	script = HTMLescap2Str(script)
	if len(res) == 0 {
		if msg.MessageType == "group" {
			RespGroupMsg(msg, "No value return, and the vm won't run.")
		} else if msg.MessageType == "private" {
			RespPrivateMsg(msg, "No value return, and the vm won't run.")
		}
		return
	}
	res = res[3 : len(res)-3]
	values := regexp.MustCompile(`[0-9a-zA-Z_\$=]{1,}`).FindAllString(string(res), 100)
	vm := otto.New()
	var resp string
	output, err := vm.Run(script)
	if err != nil {
		if msg.MessageType == "group" {
			RespGroupMsg(msg, fmt.Sprintln("ott: ", err))
		} else if msg.MessageType == "private" {
			RespPrivateMsg(msg, fmt.Sprintln("ott: ", err))
		}
		return
	}
	length := len(values)
	resp = "Output: " + output.String() + "\n"
	for i := 0; i < length; i++ {
		res, err := vm.Get(values[i])
		if err != nil {
			if msg.MessageType == "group" {
				RespGroupMsg(msg, fmt.Sprintln("ott: ", err))
			} else if msg.MessageType == "private" {
				RespPrivateMsg(msg, fmt.Sprintln("ott: ", err))
			}
			return
		}
		resp += values[i] + ": " + res.String() + "\n"
	}
	if msg.MessageType == "group" {
		RespGroupMsg(msg, resp)
	} else if msg.MessageType == "private" {
		RespPrivateMsg(msg, resp)
	}
}

func Permitted(id int64) bool {
	length := len(uni.Conf.JSpermit)
	for i := 0; i < length; i++ {
		if int(id) == uni.Conf.JSpermit[i] {
			return true
		}
	}
	return false
}

func HTMLescap2Str(input string) string {
	input = strings.ReplaceAll(input, "&#91;", "[")
	input = strings.ReplaceAll(input, "&#93;", "]")
	input = strings.ReplaceAll(input, "&#60;", "<")
	input = strings.ReplaceAll(input, "&#62;", ">")
	input = strings.ReplaceAll(input, "&#38;", "&")
	input = strings.ReplaceAll(input, "&lt;", "<")
	input = strings.ReplaceAll(input, "&gt;", ">")
	input = strings.ReplaceAll(input, "&amp;", "&")
	return input
}
