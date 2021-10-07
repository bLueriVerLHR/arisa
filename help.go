package main

import "arisa/tools"

func Help(msg string) string {
	if tools.Grep(`^Arisa --help$`, msg) {
		return `Arisa created by BlueRiverLHR

用到的开源项目有：
日志：
github.com/sirupsen/logrus
yaml解析：
gopkg.in/yaml.v2
网页爬取：
github.com/PuerkitoBio/goquery

该项目基于：
go-cqhttp: https://github.com/Mrs4s/go-cqhttp
制作的插件

目前功能：

<--! 以下每一个功能都不能滥用，如果滥用导致 Arisa 封禁，该用户/群列入踢出功能白名单 -->

回应戳一戳
回应撤回
正则聊天
打招呼
语录
点歌
名言
夸人
诅咒
发涩图
pixiv榜
js解释器otto
盲盒
coser
新闻
查mooc
搜索

Admin功能

可以使用：（有空格）
Arisa --help + 功能中文名称
进一步了解信息`
	} else if tools.Grep(`^Arisa --help 回应戳一戳$`, msg) {
		return `回应戳一戳
戳一戳Arisa吧！`
	} else if tools.Grep(`^Arisa --help 回应撤回$`, msg) {
		return `回应撤回
特别友好的撤回响应机制`
	} else if tools.Grep(`^Arisa --help 正则聊天$`, msg) {
		return `正则聊天
基于正则表达式的聊天系统

可以自定义正则
但是一定要保证不会影响到正常的对话聊天
而且不能与之前的正则冲突`
	} else if tools.Grep(`^Arisa --help 打招呼$`, msg) {
		return `打招呼
基于正则表达式的按照时间回复早晚安

正则：

时间：0-4
(?i)(^早安|^早上好|^おはよう|^ohayo$|^Good Morning|^ohakusa$|^ohayayo$|^oha呀(呦|哟)$|^oha.*desu$|^早$|안녕하세요|좋은아침이에요|안녕|좋은아침)

(?i)(oyasumi|晚安|おやすみ)

时间：5-9
(?i)(^早安|^早上好|^おはよう|^ohayo$|^Good Morning|^ohakusa$|^ohayayo$|^oha呀(呦|哟)$|^oha.*desu$|^早$|안녕하세요|좋은아침이에요|안녕|좋은아침)

时间：10-13
午安|中午好

时间：14-17
下午好

时间：18-23
晚上好

(?i)(oyasumi|晚安|おやすみ)`
	} else if tools.Grep(`^Arisa --help 语录`, msg) {
		return `语录
正则：
ss：
hy：

注意，冒号为中文冒号
Arisa会返回某个人的经典语录
或者经典网络名言和舔狗日记
第二个则特定要hy的`
	} else if tools.Grep(`^Arisa --help 点歌`, msg) {
		return `点歌
正则：
Arisa点歌\s.{1,}`
	} else if tools.Grep(`^Arisa --help 名言`, msg) {
		return `名言
正则：
Arisa.*名言

支持动画，漫画，游戏，文学
原创，网络，随便，影视
诗词，网易云，哲学，抖机灵
类型的名言

默认为文学`
	} else if tools.Grep(`^Arisa --help 夸人`, msg) {
		return `夸人
正则：
Arisa.*夸.|Arisa.*彩虹屁

回复任意彩虹屁，情话`
	} else if tools.Grep(`^Arisa --help 诅咒`, msg) {
		return `夸人
正则：
Arisa诅咒

很恶毒的诅咒`
	} else if tools.Grep(`^Arisa --help 发涩图$`, msg) {
		return `发涩图
命令行风格的随机涩图
指令和属性，属性和属性之间需要加空格
来源 Pixiv

指令：
--setu

属性解释：
keyword=关键字  默认无
uid=画师id  默认无
tag=标签  默认无`
	} else if tools.Grep(`^Arisa --help pixiv榜$`, msg) {
		return `pixiv榜
命令行风格的pixiv榜查询
指令和属性，属性和属性之间需要加空格
每日11:00更新

指令：
--pixiv

属性解释：
rank=排名(1-30) 默认1
date=年月日(yyyy-mm-dd) 默认当日
type=类型

类型支持：
pixiv_normal
pixiv_male_r18
pixiv_male
yandere`
	} else if tools.Grep(`^Arisa --help js解释器otto`, msg) {
		return `js解释器otto
正则：
<--.{1,}-->

用法：
第一行 <--.{1,}--> 用于定义需要读取的变量的值

例如：
<--a,b,c,d-->
意思是获取代码里 a, b, c, d 的值
如果是对象，则会返回其 toString()

接下来就是输入 js 代码

例如：
<--a,b-->
a = 2
b = 3
a = b

Arisa会返回
Output: 3
a: 3
b: 3

其中Output是最后一行表达式的值`
	} else if tools.Grep(`^Arisa --help 盲盒`, msg) {
		return `盲盒
正则：
开盲盒

一堆神秘链接和图片，我也不清楚都有啥
有的质量不错，有的就狠拉跨`
	} else if tools.Grep(`^Arisa --help coser`, msg) {
		return `coser
正则：
bcycos
mycoser
isoyucos

bcycos返回半次元当日排行榜上的cos图`
	} else if tools.Grep(`^Arisa --help 新闻`, msg) {
		return `新闻
正则：
Arisa本地新闻
Arisa.*新闻

第二个支持搜索
头条，军事，娱乐，体育
科技，艺术，教育，要闻`
	} else if tools.Grep(`^Arisa --help 半次元`, msg) {
		return `半次元
正则：
半次元|bcy

配合：
小说|文|粮
图
cos

配合：
日|天
新人
默认：周

实例：
今天的半次元有啥粮嘛？`
	} else if tools.Grep(`^Arisa --help 搜索`, msg) {
		return `搜索
萌娘搜
百度搜
^Arisa.{1,}是什么$
^Arisa.{1,}是什么？$

第一个是用萌娘百科，其余的是用百度百科
Arisa是(什么|啥)`
	} else if tools.Grep(`^Arisa --help 查mooc`, msg) {
		return `查mooc
该功能目前仅应用于私聊
还需要打磨
目前还没有做多线程欢迎交流

目前用法：
mooc -i
获取mooc登录二维码
扫码登陆
登录成功会提示
Arisa会在一分钟后撤回该二维码
如果没能在二维码失效的时候登录
会再次重发一个
如果要停止该过程
正则如下：
停下|stop|Stop|Quit|q|quit|Exit|.exit|exit

更新课程列表
mooc -u

查询课程对应序号
mooc -l

查询所有课程的作业
仅互评时间还没过去的
mooc -a
mooc

查询某个课程
mooc -n + 序号
例如：
mooc -n 1
就是查询课程列表里第一个课程的作业`
	} else if tools.Grep(`^Arisa --help Admin功能`, msg) {
		return `Admin功能
正则：
Arisa set

无输出，可以刷新数据

配合：
quo= 语录
gtalk g= 正则 m= 返回
add bad= QQ号
	setu= 群号
	pixiv= 群号
	jsp= QQ号
list quo
	gtalk
	conf
del quo= 序号
		hy
	bad= 序号
	setu= 序号
	jsp= 序号`
	}
	return ""
}
