package main

import (
	"regexp"
	"strconv"
)

func FindTargetAttibute(msgJSON string, target string) string {
	index := regexp.MustCompile(target).FindAllIndex([]byte(msgJSON), 1000)
	counter := 0
	Range := []int{0, 0}
	if len(index) == 0 {
		return ""
	}
	for i := index[0][1] + 1; i < index[0][1]+50; i++ {
		if msgJSON[i] == '"' {
			counter++
			if counter == 1 {
				Range[0] = i
			} else if counter == 2 {
				Range[1] = i
				break
			}
		}
	}
	return msgJSON[Range[0]+1 : Range[1]]
}

func Greps(msg RecvMsg) ([]grep, int) {
	tmp := []grep{
		{"摸摸Arisa|Arisa摸摸头|Arisa摸摸", "mufumufu，好舒服[CQ:face,id=66][CQ:face,id=305]"},
		{"Arisa我爱你|我爱你Arisa|Arisa.*爱你|爱你.*Arisa", "莎莎也爱你哦[CQ:face,id=66][CQ:face,id=319]"},
		{"^Arisa$", "我在呢[CQ:face,id=66]"},
		{"嘉然|嘉人|嘉心糖", "为了[CQ:at,qq=" + strconv.FormatInt(msg.UserID, 10) + "]我要听猫中毒啦！！！"},
		{"Arisa你是谁|Arisa是谁|3476173182.*是谁|Arisa你是.*啥|3476173182.*你是.*啥", "我是本群最可爱的小bot[CQ:face,id=66]"},
		{"Arisa你是神吗？", "[CQ:image,file=e90fda80008cdceeba94a257c0062fa7.image]"},
		{"(3476173182|Arisa).*(tm|nm|鸡吧|寄吧|SB|傻逼|沙雕|傻子|几把谁|勾八|勾⑧|你🐎没了|你马没了|你妈没了|草泥马|操你妈|草你妈|你妈).*", "不要骂人啦！[CQ:image,file=065d0c00e49e8d18c077c6e8d47c60c6.image]"},
		{"gay群|是gay", "？我不懂"},
		{"我们班卷王是谁", "每个人都觉得自己不卷，但每个人都很卷。--鲁迅没有说过"},
		{"Arisa.*么么哒", "么么哒[CQ:face,id=305]"},
		{"Arisa.*(憨批|憨憨|憨比|逗比|逗逼)", "我憨憨的不可爱吗？[CQ:image,file=4eb27bee1d1ab28fa16949c81f3278ec.image]"},
		{"喜欢开车是吧.", "快拿刺雷捅TA，板载！！！"},
		{"不醉不归", "兄弟萌，我先干了！\n[CQ:image,file=86064829b5691359a4af684c1780c84f.image]\n[CQ:image,file=90e2057eb6e88322486ed91d66d843f8.image]"},
		{"好可爱", "赞同！[CQ:image,file=b387314685f2703330d8c3e9b7dde4ac.image]"},
		{"什么鬼", "就是就是，什么鬼哦！[CQ:image,file=e1f48c1ed1bc7cbede1680b3d37d97ec.image]"},
		{"抽卡.*双黄蛋|双黄蛋|人生第一次.*抽卡.*抽出.*个(SSR|六星|五星)", "[CQ:image,file=60d331653b0e786fff91533f60f974aa.image]"},
		{"Arisa.*骂我", "[CQ:image,file=1afd32a9ad8a15903b324fd22fceb2b6.image]"},
		{"原地去世", "[CQ:image,file=96d5f5b146b3378c92a6ea3350890790.image]"},
		{"跑快快", "[CQ:image,file=4eb5c42525c0f5a895d606770241add8.image]"},
		{"好冷的笑话|冷笑话", "[CQ:image,file=93cbed2aadf68580067434770bf989c7.image]"},
		{"惊.*了|震惊", "[CQ:image,file=dc3a0c079512991e4a4229acf87f9d95.image]"},
		{"累了", "累了[CQ:image,file=8ff10d22629e26f3571560389bd3aeb5.image]"},
		{"Arisa.*你.*(矮|平)", "[CQ:image,file=63d76651e63e467616984317e31ee5c3.image]"},
		{"咸鱼|躺平", "[CQ:image,file=6bd1337ee62f52cc5de813b7e9f3ba4b.image]"},
		{"Arisa.*(高数|线代|高等数学|线性代数|复变函数|大[物雾]|[Mm]at[Ll]ab|C|C.*|能帮我|写.*(作业|论文|报告))吗", "现在不能！\n[CQ:image,file=d7c77fc43e6482b350c3f76d84f869c7.image]"},
		{"Arisa健康上报", "不用担心错过健康信息上报的时间了\n早上填体温网址：\nhttps://e-report.neu.edu.cn/inspection/items/1/records/create\n中午填体温网址：\nhttps://e-report.neu.edu.cn/inspection/items/2/records/create\n下午填体温网址：\nhttps://e-report.neu.edu.cn/inspection/items/3/records/create"},
		{"好(色|涩|铯)哦|hso", "什么什么？让我康康！[CQ:image,file=428e30d054f13d752ff11abec7eb55f8.image]"},
		{"我是(loli控|萝莉控)", "[CQ:image,file=32f4e6112ad890f26321fb18e7e19706.image]"},
		{"懂得都懂", "[CQ:image,file=9f1e2333d70ef35f4e56c1f296eb4b4a.image]"},
		{"我可爱吗", "[CQ:image,file=edc3064dd71eabcc0e7542d7e262fd70.image]"},
		{"塔诺西", "[CQ:image,file=ca47a9b54b4e8238f1e3314b14c187e2.image]"},
		{"苦露西", "[CQ:image,file=8345391b3adb3a9decb32fa989965c7b.image]"},
		{"吃瓜吃瓜", "[CQ:image,file=51bf25bf889494903a903a2da000005f.image]"},
		{"无聊啊啊*", "[CQ:image,file=2b6c909efae2d84f6fff2ff4ca2e46b0.image]"},
		{"114514|哼.*啊啊*", "[CQ:image,file=f0f83d53ee4766a723d38b825fe1304b.image]"},
		{"无语子|绝绝子", "[CQ:image,file=cb819441ad5dd48486f2a556be08b243.image]"},
		{"我(牛逼|厉害|nb)吧", "[CQ:image,file=c94852db26a502fe226102ee5b9706fc.image]"},
		{"细(说|🔒)", "细🔒"},
		{"Arisa.*(口我|揍我|转.圈|透我)|魅魔|(透|草)Arisa", "[CQ:image,file=f857a51cd2c0f3d7299993aed15a15e1.image]"},
		{"二次元.*爬|Arisa.*爬|能不能爬", "😅急了是吧？是不是觉得户口本上只有一页还多了？是个人就给我好好说话，别成天爬来爬去好🐴"},
		{".*急了.*急了", "🤣词穷了是吧？[CQ:at,qq=" + strconv.FormatInt(msg.UserID, 10) + "]"},
		{"😅", "e^{💧ln😄}"},
		{"男人.*默契", "[CQ:share,url=https://www.bilibili.com/video/BV1Uy4y187WB,title=男人的默契]"},
		{"打五把csgo", "[CQ:share,url=https://www.bilibili.com/video/BV1AK411g7xc,title=IndiHome]"},
		{"抬走", "     🕺🏿 🕺🏿 🕺🏿\n🕺🏿      ⚰️ \n     🕺🏿 🕺🏿 🕺🏿"},
		{"黑人抬棺", "[CQ:share,url=https://www.bilibili.com/video/BV1NZ4y1j7nw,title=黑人抬棺]"},
		{"Arisa --version", version},
		{"Arisa --blog", "[CQ:share,url=https://blueriverlhr.github.io/,title=蓝川的世界]"},
	}
	return tmp, len(tmp)
}
