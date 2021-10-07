package main

import (
	"arisa/tools"
	"fmt"
	"regexp"
	"strconv"
)

func Settings(msg RecvMsg) {
	if regexp.MustCompile(`set`).Match([]byte(msg.Message)) {
		if reg := regexp.MustCompile(`quo=[^\s]{1,}`); reg.Match([]byte(msg.Message)) {
			value := reg.Find([]byte(msg.Message))
			if len(value) == 0 {
				RespPrivateMsg(msg, "参数不够")
				return
			}
			if regexp.MustCompile(`hy：`).Match([]byte(msg.Message)) {
				uni.SomeS.Hy = append(uni.SomeS.Hy, string(value)[4:])
			} else {
				uni.SomeS.Ss = append(uni.SomeS.Ss, string(value)[4:])
			}
			RespPrivateMsg(msg, "添加成功")
		} else if regexp.MustCompile(`gtalk`).Match([]byte(msg.Message)) {
			valueg := regexp.MustCompile(`g=[^\s]{1,}`).Find([]byte(msg.Message))
			valuem := regexp.MustCompile(`m=[^\s]{1,}`).Find([]byte(msg.Message))
			if len(valueg) == 0 || len(valuem) == 0 {
				RespPrivateMsg(msg, "参数不够")
				return
			}
			uni.Grep.Greps = append(uni.Grep.Greps, struct {
				Regexp  string "yaml:\"regexp\""
				Message string "yaml:\"message\""
			}{
				string(valueg)[2:],
				string(valuem)[2:],
			})
			RespPrivateMsg(msg, "添加成功")
		} else if regexp.MustCompile(`add`).Match([]byte(msg.Message)) {
			valueb := regexp.MustCompile(`bad=[^\s]{1,}`).Find([]byte(msg.Message))
			if len(valueb) != 0 {
				qq, err := strconv.Atoi(string(valueb)[4:])
				tools.Check(err)
				uni.Conf.BadGirl.White = append(uni.Conf.BadGirl.White, qq)
				RespPrivateMsg(msg, "添加成功")
			}
			values := regexp.MustCompile(`setu=[^\s]{1,}`).Find([]byte(msg.Message))
			if len(values) != 0 {
				qq, err := strconv.Atoi(string(values)[5:])
				tools.Check(err)
				uni.Conf.Setu.White = append(uni.Conf.Setu.White, qq)
				RespPrivateMsg(msg, "添加成功")
			}
			valuep := regexp.MustCompile(`pixiv=[^\s]{1,}`).Find([]byte(msg.Message))
			if len(valuep) != 0 {
				qq, err := strconv.Atoi(string(valuep)[6:])
				tools.Check(err)
				uni.Conf.Pixiv.White = append(uni.Conf.Pixiv.White, qq)
				RespPrivateMsg(msg, "添加成功")
			}
			valuej := regexp.MustCompile(`jsp=[^\s]{1,}`).Find([]byte(msg.Message))
			if len(valuej) != 0 {
				qq, err := strconv.Atoi(string(valuej)[4:])
				tools.Check(err)
				uni.Conf.JSpermit = append(uni.Conf.JSpermit, qq)
				RespPrivateMsg(msg, "添加成功")
			}
		}
	}
	uni.Save()
	if regexp.MustCompile(`list`).Match([]byte(msg.Message)) {
		if regexp.MustCompile(`quo`).Match([]byte(msg.Message)) {
			var resp string
			length := len(uni.SomeS.Hy)
			for i := 0; i < length; i++ {
				resp += strconv.Itoa(i) + ". " + uni.SomeS.Hy[i] + "\n"
			}
			length = len(uni.SomeS.Ss)
			for i := 0; i < length; i++ {
				resp += strconv.Itoa(i) + ". " + uni.SomeS.Ss[i] + "\n"
			}
			RespPrivateMsg(msg, resp)
		}
		if regexp.MustCompile(`gtalk`).Match([]byte(msg.Message)) {
			var resp string
			length := len(uni.Grep.Greps)
			for i := 0; i < length; i++ {
				resp += strconv.Itoa(i) + "." + "\n" + "regexp: \n" + uni.Grep.Greps[i].Regexp + "\n" + "message: \n" + uni.Grep.Greps[i].Message + "\n"
			}
			RespPrivateMsg(msg, resp)
		}
		if regexp.MustCompile(`conf`).Match([]byte(msg.Message)) {
			resp := fmt.Sprintln(uni.Conf)
			RespPrivateMsg(msg, resp)
		}
	}
	if regexp.MustCompile(`del`).Match([]byte(msg.Message)) {
		if reg := regexp.MustCompile(`quo=[0-9]{1,}`); reg.Match([]byte(msg.Message)) {
			value := reg.Find([]byte(msg.Message))
			if len(value) == 0 {
				RespPrivateMsg(msg, "参数不够")
				return
			}
			num, err := strconv.Atoi(string(value)[4:])
			tools.Check(err)
			if regexp.MustCompile(`hy：`).Match([]byte(msg.Message)) {
				uni.SomeS.Hy = append(uni.SomeS.Hy[:num], uni.SomeS.Hy[num+1:]...)
			} else {
				uni.SomeS.Ss = append(uni.SomeS.Ss[:num], uni.SomeS.Ss[num+1:]...)
			}
			RespPrivateMsg(msg, "删除成功")
		} else if reg := regexp.MustCompile(`gtalk=[0-9]{1,}`); reg.Match([]byte(msg.Message)) {
			value := reg.Find([]byte(msg.Message))
			if len(value) == 0 {
				RespPrivateMsg(msg, "参数不够")
				return
			}
			num, err := strconv.Atoi(string(value)[6:])
			tools.Check(err)
			uni.Grep.Greps = append(uni.Grep.Greps[:num], uni.Grep.Greps[num+1:]...)
		} else if reg := regexp.MustCompile(`bad=[0-9]{1,}}`); reg.Match([]byte(msg.Message)) {
			valueb := reg.Find([]byte(msg.Message))
			if len(valueb) != 0 {
				num, err := strconv.Atoi(string(valueb)[4:])
				tools.Check(err)
				uni.Conf.BadGirl.White = append(uni.Conf.BadGirl.White[:num], uni.Conf.BadGirl.White[num+1:]...)
				RespPrivateMsg(msg, "添加成功")
				return
			}
			RespPrivateMsg(msg, "参数不够")
			return
		} else if reg := regexp.MustCompile(`setu=[0-9]{1,}`); reg.Match([]byte(msg.Message)) {
			values := reg.Find([]byte(msg.Message))
			if len(values) != 0 {
				num, err := strconv.Atoi(string(values)[5:])
				tools.Check(err)
				uni.Conf.BadGirl.White = append(uni.Conf.Setu.White[:num], uni.Conf.Setu.White[num+1:]...)
				RespPrivateMsg(msg, "添加成功")
				return
			}
			RespPrivateMsg(msg, "参数不够")
			return
		} else if reg := regexp.MustCompile(`pixiv=[0-9]{1,}`); reg.Match([]byte(msg.Message)) {
			valuep := reg.Find([]byte(msg.Message))
			if len(valuep) != 0 {
				num, err := strconv.Atoi(string(valuep)[6:])
				tools.Check(err)
				uni.Conf.BadGirl.White = append(uni.Conf.Pixiv.White[:num], uni.Conf.Pixiv.White[num+1:]...)
				RespPrivateMsg(msg, "添加成功")
				return
			}
			RespPrivateMsg(msg, "参数不够")
			return
		} else if reg := regexp.MustCompile(`jsp=[0-9]{1,}`); reg.Match([]byte(msg.Message)) {
			valuej := reg.Find([]byte(msg.Message))
			if len(valuej) != 0 {
				num, err := strconv.Atoi(string(valuej)[4:])
				tools.Check(err)
				uni.Conf.BadGirl.White = append(uni.Conf.JSpermit[:num], uni.Conf.JSpermit[num+1:]...)
				RespPrivateMsg(msg, "添加成功")
				return
			}
			RespPrivateMsg(msg, "参数不够")
			return
		}
	}
}
