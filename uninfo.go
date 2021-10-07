package main

import (
	"arisa/tools"
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Admin   int64 `yaml:"Master"`
	BadGirl struct {
		White []int `yaml:"white"`
	} `yaml:"BadGirl"`
	BotConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Post int    `yaml:"post"`
	} `yaml:"BotConfig"`
	GrepTalk struct {
		Path string `yaml:"path"`
	} `yaml:"GrepTalk"`
	Setu struct {
		White []int `yaml:"white"`
	} `yaml:"Setu"`
	Pixiv struct {
		White []int `yaml:"white"`
	} `yaml:"Pixiv"`
	MOOC struct {
		Path string `yaml:"path"`
	} `yaml:"MOOC"`
	Quotation struct {
		Path string `yaml:"path"`
	} `yaml:"Quotation"`
	JSpermit []int `yaml:"JSpermit"`
}

type GrepTalk struct {
	Greps []struct {
		Regexp  string `yaml:"regexp"`
		Message string `yaml:"message"`
	} `yaml:"Greps"`
}

type MoocAccount struct {
	List []struct {
		Qq     int64  `json:"qq"`
		Cookie string `json:"cookie"`
		Lesson []struct {
			Name    string `json:"name"`
			TermID  string `json:"termId"`
			ClassID string `json:"classId"`
		} `json:"lesson"`
	} `json:"List"`
}

type SomeoneSay struct {
	Hy []string `json:"hy"`
	Ss []string `json:"ss"`
}

type Subscribe struct {
	Mooc   []int64 `json:"mooc"`
	Health []struct {
		Qq  int64  `json:"qq"`
		UID string `json:"uid"`
		Pwd string `json:"pwd"`
	} `json:"health"`
}

type Universal struct {
	Conf         Config
	Grep         GrepTalk
	MoocA        MoocAccount
	SomeS        SomeoneSay
	Slist        Subscribe
	Repeat       int
	LastMsg      string
	On           bool
	MoocLoopFlag bool
	Operating    map[int64]bool
	Report       bool
}

func (uni *Universal) Get() {
	// config
	yml, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yml, &uni.Conf)
	if err != nil {
		panic(err)
	}
	// mooc
	mooc, err := ioutil.ReadFile(uni.Conf.MOOC.Path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(mooc, &uni.MoocA)
	if err != nil {
		panic(err)
	}
	// grep
	greps, err := ioutil.ReadFile(uni.Conf.GrepTalk.Path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(greps, &uni.Grep)
	if err != nil {
		panic(err)
	}
	// quotation
	quo, err := ioutil.ReadFile(uni.Conf.Quotation.Path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(quo, &uni.SomeS)
	if err != nil {
		panic(err)
	}
	// subscribe
	sub, err := ioutil.ReadFile("./data/subscribe.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(sub, &uni.Slist)
	if err != nil {
		panic(err)
	}
	uni.Operating = make(map[int64]bool)
}

func (uni *Universal) Save() {
	text, err := json.Marshal(uni.MoocA)
	tools.Check(err)
	err = ioutil.WriteFile(uni.Conf.MOOC.Path, text, 0666)
	tools.Check(err)
	text, err = json.Marshal(uni.SomeS)
	tools.Check(err)
	err = ioutil.WriteFile(uni.Conf.Quotation.Path, text, 0666)
	tools.Check(err)
	text, err = yaml.Marshal(uni.Grep.Greps)
	tools.Check(err)
	text = []byte("Greps:\n" + string(text))
	err = ioutil.WriteFile(uni.Conf.GrepTalk.Path, text, 0666)
	tools.Check(err)
	text, err = yaml.Marshal(uni.Conf)
	tools.Check(err)
	err = ioutil.WriteFile("./aconfig.yml", text, 0666)
	tools.Check(err)
	text, err = json.Marshal(uni.Slist)
	tools.Check(err)
	err = ioutil.WriteFile("./data/subscribe.json", text, 0666)
	tools.Check(err)
}
