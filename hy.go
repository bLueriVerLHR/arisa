package main

func hy(flag int, msg RecvMsg) {
	switch flag {
	case 0:
		respGroupMsg(msg, "hy：听说我这个专业前景不好")
	case 1:
		respGroupMsg(msg, "hy：我太菜了")
	case 2:
		respGroupMsg(msg, "hy：快来人骂我")
	case 3:
		respGroupMsg(msg, "hy：我不需要老婆")
	case 4:
		respGroupMsg(msg, "hy：我是抖m")
	case 5:
		respGroupMsg(msg, "sy：和四个妹妹\nhy：都约好了")
	case 6:
		respGroupMsg(msg, "hy：漂亮的我都喜欢")
	case 7:
		respGroupMsg(msg, "hy：用pg拉拢导员和教学办")
	case 8:
		respGroupMsg(msg, "hy：我真的啥都不知道")
	case 9:
		respGroupMsg(msg, "hy：nw一般在我腿上")
	case 10:
		respGroupMsg(msg, "hy：我是小丑")
	case 11:
		respGroupMsg(msg, "hy：有人压榨我真的好爽这种感觉")
	case 12:
		respGroupMsg(msg, "hy：gpa就是个屁")
	case 13:
		respGroupMsg(msg, "hy：有人压榨我真的好爽这种感觉")
	default:
		respGroupMsg(msg, "hy：听说隔壁专业有40多个绩点4.0以上的")
	}
}

func SomeoneSay(flag int, msg RecvMsg) {
	switch flag {
	case 0:
		respGroupMsg(msg, "sy：我有一宿舍")
	case 1:
		respGroupMsg(msg, "mjj：你可以把我日得喵喵叫吗？\nmjj：喵呜")
	case 2:
		respGroupMsg(msg, "mjj：昕昕快来\nmjj：让我爽爽")
	case 3:
		respGroupMsg(msg, "mjj：最近突然对屁股感兴趣了")
	case 4:
		respGroupMsg(msg, "mjj：呜呜呜呜\nmjj：呜呜呜呜\nmjj：npy越看越好看怎么办\nmjj：[CQ:face,id=9]")
	case 5:
		respGroupMsg(msg, "mjj：表白是不可能表白的\nmjj：这辈子都不可能\nmjj：我就是从这里跳下去，从这里摔死\nmjj：也不会表白")
	case 6:
		respGroupMsg(msg, "mjj：男孩子还是谈一场然后分掉好")
	case 7:
		respGroupMsg(msg, "yjjj：墨佬要亲亲学弟\nmjj：正常")
	case 8:
		respGroupMsg(msg, "mjj：好想来个大大的拥抱\nmjj：软软的那种")
	case 10:
		respGroupMsg(msg, "zggg：nw马上晋升院士了吗？\nds：我也")
	case 11:
		respGroupMsg(msg, "sy：你寄吧谁啊")
	case 12:
		respGroupMsg(msg, "ds：刚洗完澡出来\nds：凉快了")
	case 13:
		respGroupMsg(msg, "ds：我亲学妹\nds：抱着亲")
	case 14:
		respGroupMsg(msg, "ds：我直接脱下抢走\nds：管他那么多\nds：照冲不误\n……\nds：这女的什么牛马")
	case 15:
		respGroupMsg(msg, "cy：我加学妹又怎么了\ncy：只要不被发现\ncy：我还能继续跳")
	case 16:
		respGroupMsg(msg, "yjjj：你们想发的话也可以\nyjjj：我也想让我npy打打我\nyjjj：我老M了")
	case 17:
		respGroupMsg(msg, "mjj：我这么久都没对象一般是你的锅\nmjj：[CQ:face,id=178][CQ:face,id=67]\nzggg：关我屁事\nzggg：自己点球射不进去怪队友")
	case 18:
		respGroupMsg(msg, "mjj：我都要了\nmjj：正好我有两个学妹嘿嘿嘿")
	case 19:
		respGroupMsg(msg, "mjj：小姐姐已经两天没回我消息了\nmjj：好爽啊")
	case 20:
		respGroupMsg(msg, "mrr：hy被我草晕了")
	case 21:
		respGroupMsg(msg, "mjj：但我也不知道为啥她们把我删了")
	case 22:
		respGroupMsg(msg, "mjj：giegie\nmjj：导员进了我们宿舍\nmjj：在那里摸我的枪\n")
	case 23:
		respGroupMsg(msg, "zggg：女孩子身体热起来，皮肤变得有些红的话，说明她已经动情了")
	case 24:
		respGroupMsg(msg, "mjj：删屁股")
	case 25:
		respGroupMsg(msg, "baibai：什么吃饭\nbaibai：吃什么饭\nbaibai：有饭吃？")
	case 26:
		respGroupMsg(msg, "ds：与数千名女生有染")
	case 27:
		respGroupMsg(msg, "mjj：我现在专心看学妹\nmjj：yzjj麻烦等一等")
	case 28:
		respGroupMsg(msg, "mrr：我tm直接抱起yyf原地冲刺")
	case 29:
		respGroupMsg(msg, "yyf：我tm直接抱起mrr原地冲刺")
	case 30:
		respGroupMsg(msg, "mjj：我想娶一个平淡，看起来舒服的女子，她不用很漂亮，也不要涂很重的口红，擦很重的胭脂，也不需要爱慕虚荣，她要有一个温柔大方的心，希望她可以看到我满眼的爱恋，希望她可以喜欢我二到不行的笑话，如果能遇到的话，我一定会捧着她的脸对她说，遇到你真的是我一生中最幸运的事")
	case 31:
		respGroupMsg(msg, "mjj：完屁股")
	case 32:
		respGroupMsg(msg, "mjj：吃屁股")
	case 33:
		respGroupMsg(msg, "mjj：笑个屁股")
	case 34:
		respGroupMsg(msg, "小妮：我超勇的！")
	case 35:
		respGroupMsg(msg, "yjjj：朱墨，我第一次见你时9月6号那天你真的挺帅的，很有风范，红色卫衣，这一年马上就要过去了，最近我老是想起曾经种种，仿佛在昨天我还是刚下火车刚迈进校园一般。我们一起在9月12号那天去换可乐，晚上一起去团建。感觉就像是在昨天一样，第二天国庆节，一晚上兴奋得睡不着，第二天送走易凡后，我们一起顶着伞去鱼市，路过泥泞的小巷，你走散了，我回去找你，之后我差点和拉泡沫的车撞上。从鱼市出来看过花花草草，拐角的老妇在卖瓶装鱼，之后回去时你和我们错开了。在大润发我们一起舔酸奶小杯，我买的寿司没有条形码，去人工柜台结账，让你们等了好久，理发时我们一起等jk。冒雨从西门出来原来才不过2点多。留下的美好回忆，永远记在心里。以后有机会再一起出去。年关将至，心有所感，聊表抒怀，不知所言。人生不过梦一场，今宵别梦寒，爱你一万年")
	case 36:
		respGroupMsg(msg, "贞子：汪汪汪")
	case 37:
		respGroupMsg(msg, "yjjj：不瞒你说我昨天做梦梦到自己变成小姐姐了")
	default:
		respGroupMsg(msg, "dslsp")
	}
}
