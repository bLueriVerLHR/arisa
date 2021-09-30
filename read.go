package main

import (
	"encoding/json"
	"fmt"
)

func (rm *RecvMsg) Read(recv string) {
	flag := []int{0, 0}
	length := len(recv)
	for i := 0; i < length; i++ {
		if recv[i] == '{' {
			flag[0] = i
			break
		}
	}
	for i := length - 1; i >= 0; i-- {
		if recv[i] == '}' {
			flag[1] = i
			break
		}
	}
	err := json.Unmarshal([]byte(recv[flag[0]:flag[1]+1]), rm)
	if err != nil {
		fmt.Println(recv)
	}
	CheckError(err)
}

func (bs *BotStatus) Read(status string) {
	flag := []int{0, 0}
	length := len(status)
	for i := 0; i < length; i++ {
		if status[i] == '{' {
			flag[0] = i
			break
		}
	}
	for i := length - 1; i >= 0; i-- {
		if status[i] == '}' {
			flag[1] = i
			break
		}
	}
	err := json.Unmarshal([]byte(status[flag[0]:flag[1]+1]), bs)
	if err != nil {
		fmt.Println(status)
	}
	CheckError(err)
}

func (st *Setu) Read(request string) {
	flag := []int{0, 0}
	length := len(request)
	for i := 0; i < length; i++ {
		if request[i] == '{' {
			flag[0] = i
			break
		}
	}
	for i := length - 1; i >= 0; i-- {
		if request[i] == '}' {
			flag[1] = i
			break
		}
	}
	err := json.Unmarshal([]byte(request[flag[0]:flag[1]+1]), st)
	if err != nil {
		fmt.Println(request)
	}
	CheckError(err)
}

func (te *TouchEvent) Read(recv string) {
	flag := []int{0, 0}
	length := len(recv)
	for i := 0; i < length; i++ {
		if recv[i] == '{' {
			flag[0] = i
			break
		}
	}
	for i := length - 1; i >= 0; i-- {
		if recv[i] == '}' {
			flag[1] = i
			break
		}
	}
	err := json.Unmarshal([]byte(recv[flag[0]:flag[1]+1]), te)
	if err != nil {
		fmt.Println(recv)
	}
	CheckError(err)
}

func (mi *MsgId) Read(recv string) {
	flag := []int{0, 0}
	length := len(recv)
	for i := 0; i < length; i++ {
		if recv[i] == '{' {
			flag[0] = i
			break
		}
	}
	for i := length - 1; i >= 0; i-- {
		if recv[i] == '}' {
			flag[1] = i
			break
		}
	}
	err := json.Unmarshal([]byte(recv[flag[0]:flag[1]+1]), mi)
	if err != nil {
		fmt.Println(recv)
	}
	CheckError(err)
}

func (p *Pixiv) Read(recv string) {
	err := json.Unmarshal([]byte(recv), p)
	if err != nil {
		fmt.Println(recv)
	}
	CheckError(err)
}
