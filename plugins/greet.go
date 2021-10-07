package plugins

import (
	"arisa/tools"
	"math/rand"
	"time"
)

func Greeting(msg string) string {
	rand.Seed(time.Now().Unix())
	China, err := time.LoadLocation("Asia/Shanghai")
	tools.Check(err)
	now := time.Now().In(China).Hour()
	if now <= 4 && now >= 0 {
		if tools.Grep(`(?i)(^早安|^早上好|^おはよう|^ohayo$|^Good Morning|^ohakusa$|^ohayayo$|^oha呀(呦|哟)$|^oha.*desu$|^早$|안녕하세요|좋은아침이에요|안녕|좋은아침)`, msg) {
			switch rand.Int() % 3 {
			case 0:
				return "(～﹃～)~zZ"
			case 1:
				return "不懂就问，这就是卷王的早上吗？"
			case 2:
				return "？？？"
			}
		}
		if tools.Grep(`(?i)(oyasumi|晚安|おやすみ)`, msg) {
			switch now {
			case 1:
				return "早早睡，早早起，明天还是一条好汉！"
			case 2:
				return "WOW，卷王！"
			default:
				return "(～﹃～)~zZ"
			}
		}
	} else if now <= 5 && now >= 9 {
		if tools.Grep(`(?i)(^早安|^早上好|^おはよう|^ohayo$|^Good Morning|^ohakusa$|^ohayayo$|^oha呀(呦|哟)$|^oha.*desu$|^早$|안녕하세요|좋은아침이에요|안녕|좋은아침)`, msg) {
			switch rand.Int() % 4 {
			case 0:
				return "人为什么要起床？.jpg"
			case 1:
				return "好困……(～﹃～)~zZ"
			case 2:
				return "早睡早起身体好！"
			case 3:
				return "正在和被子大战300回合……(～﹃～)~zZ"
			}
		}
	} else if now <= 13 && now >= 10 {
		if tools.Grep(`午安|中午好`, msg) {
			return "稍微有点困了呢"
		}
	} else if now <= 17 && now >= 14 {
		if tools.Grep(`下午好`, msg) {
			return "下午也要精神满满哦！"
		}
	} else if now <= 23 && now >= 18 {
		if tools.Grep(`晚上好`, msg) {
			return "记得不要熬夜，要早点睡哦。"
		}
		if tools.Grep(`(?i)(oyasumi|晚安|おやすみ)`, msg) {
			switch now {
			case 1:
				return "早早睡，早早起，明天还是一条好汉！"
			case 2:
				return "WOW，卷王！"
			default:
				return "(～﹃～)~zZ"
			}
		}
	}
	return ""
}
