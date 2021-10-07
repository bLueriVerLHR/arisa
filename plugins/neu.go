package plugins

import (
	"arisa/tools"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Temperature(num string, Uid string, Password string) bool {
	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}

	req, err := http.NewRequest("GET", "https://e-report.neu.edu.cn/login/neupass", nil)
	tools.Check(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(req)
	tools.Check(err)
	xsrf := resp.Cookies()

	req, err = http.NewRequest("GET", "https://pass.neu.edu.cn/tpass/login?service=https://e-report.neu.edu.cn/login/neupass/callback", nil)
	tools.Check(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	resp, err = client.Do(req)
	tools.Check(err)
	tpass := resp.Cookies()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	tools.Check(err)
	defer resp.Body.Close()
	lt, _ := doc.Find("input[name*=lt]").Attr("value")
	execution, _ := doc.Find("input[name*=execution]").Attr("value")
	_eventId, _ := doc.Find("input[name*=_eventId]").Attr("value")
	Action, _ := doc.Find("form").Attr("action")
	ul := len(Uid)
	pl := len(Password)
	rsa := Uid + Password + lt

	req, err = http.NewRequest("POST", "https://pass.neu.edu.cn"+Action, strings.NewReader("ul="+strconv.Itoa(ul)+"&pl="+strconv.Itoa(pl)+"&execution="+execution+"&rsa="+rsa+"&lt="+lt+"&_eventId="+_eventId+"&pd="+Password+"&un="+Uid))
	tools.Check(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "pass.neu.edu.cn")
	req.Header.Set("Origin", "https://pass.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://pass.neu.edu.cn/tpass/login?service=https://e-report.neu.edu.cn/login/neupass/callback")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	length := len(tpass)
	for i := 0; i < length; i++ {
		req.AddCookie(tpass[i])
	}
	resp, err = client.Do(req)
	tools.Check(err)
	tpass = resp.Cookies()
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	tools.Check(err)
	jump, _ := doc.Find("a").Attr("href")

	req, err = http.NewRequest("GET", jump, nil)
	tools.Check(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://pass.neu.edu.cn/")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	length = len(tpass)
	for i := 0; i < length; i++ {
		req.AddCookie(tpass[i])
	}
	length = len(xsrf)
	for i := 0; i < length; i++ {
		req.AddCookie(xsrf[i])
	}
	resp, err = client.Do(req)
	tools.Check(err)

	req, err = http.NewRequest("GET", "https://e-report.neu.edu.cn/login/neupass/callback", nil)
	tools.Check(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://pass.neu.edu.cn/")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	length = len(tpass)
	for i := 0; i < length; i++ {
		req.AddCookie(tpass[i])
	}
	length = len(xsrf)
	for i := 0; i < length; i++ {
		if xsrf[i].Name == "PHPSESSID" {
			req.AddCookie(resp.Cookies()[0])
		} else {
			req.AddCookie(xsrf[i])
		}
	}
	resp, err = client.Do(req)
	tools.Check(err)
	xsrf = resp.Cookies()

	req, err = http.NewRequest("GET", "https://e-report.neu.edu.cn/inspection/items/"+num+"/records/create", nil)
	tools.Check(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://pass.neu.edu.cn/")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	for i := 0; i < length; i++ {
		req.AddCookie(resp.Cookies()[i])
	}
	resp, err = client.Do(req)
	tools.Check(err)
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	tools.Check(err)
	_token, exist := doc.Find("input[name*=_token]").Attr("value")
	if exist {
		req, err = http.NewRequest("POST", "https://e-report.neu.edu.cn/inspection/items/"+num+"/records", strings.NewReader("_token="+_token+"&temperature=36.5&suspicious_respiratory_symptoms=0&symptom_descriptions="))
		tools.Check(err)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
		req.Header.Set("Host", "e-report.neu.edu.cn")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Referer", "https://e-report.neu.edu.cn/inspection/items/"+num+"/records/create")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Connection", "keep-alive")
		for i := 0; i < length; i++ {
			if xsrf[i].Name == "laravel3_session" {
				req.AddCookie(resp.Cookies()[1])
			} else if xsrf[i].Name == "XSRF-TOKEN" {
				req.AddCookie(resp.Cookies()[0])
			} else {
				req.AddCookie(xsrf[i])
			}
		}
		resp, err = client.Do(req)
		tools.Check(err)
	}
	return resp.StatusCode == 200
}

func Health(Uid string, Password string) bool {
	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}

	req, _ := http.NewRequest("GET", "https://e-report.neu.edu.cn/login/neupass", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	resp, _ := client.Do(req)
	xsrf := resp.Cookies()

	req, _ = http.NewRequest("GET", "https://pass.neu.edu.cn/tpass/login?service=https://e-report.neu.edu.cn/login/neupass/callback", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	resp, _ = client.Do(req)
	tpass := resp.Cookies()
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()
	lt, _ := doc.Find("input[name*=lt]").Attr("value")
	execution, _ := doc.Find("input[name*=execution]").Attr("value")
	_eventId, _ := doc.Find("input[name*=_eventId]").Attr("value")
	Action, _ := doc.Find("form").Attr("action")
	ul := len(Uid)
	pl := len(Password)
	rsa := Uid + Password + lt

	req, _ = http.NewRequest("POST", "https://pass.neu.edu.cn"+Action, strings.NewReader("ul="+strconv.Itoa(ul)+"&pl="+strconv.Itoa(pl)+"&execution="+execution+"&rsa="+rsa+"&lt="+lt+"&_eventId="+_eventId+"&pd="+Password+"&un="+Uid))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "pass.neu.edu.cn")
	req.Header.Set("Origin", "https://pass.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://pass.neu.edu.cn/tpass/login?service=https://e-report.neu.edu.cn/login/neupass/callback")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	length := len(tpass)
	for i := 0; i < length; i++ {
		req.AddCookie(tpass[i])
	}
	resp, _ = client.Do(req)
	tpass = resp.Cookies()
	doc, _ = goquery.NewDocumentFromReader(resp.Body)
	jump, _ := doc.Find("a").Attr("href")

	req, _ = http.NewRequest("GET", jump, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://pass.neu.edu.cn/")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	length = len(tpass)
	for i := 0; i < length; i++ {
		req.AddCookie(tpass[i])
	}
	length = len(xsrf)
	for i := 0; i < length; i++ {
		req.AddCookie(xsrf[i])
	}
	resp, _ = client.Do(req)

	req, _ = http.NewRequest("GET", "https://e-report.neu.edu.cn/login/neupass/callback", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://pass.neu.edu.cn/")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	length = len(tpass)
	for i := 0; i < length; i++ {
		req.AddCookie(tpass[i])
	}
	length = len(xsrf)
	for i := 0; i < length; i++ {
		if xsrf[i].Name == "PHPSESSID" {
			req.AddCookie(resp.Cookies()[0])
		} else {
			req.AddCookie(xsrf[i])
		}
	}
	xsrf = req.Cookies()
	resp, _ = client.Do(req)
	req, _ = http.NewRequest("GET", "https://e-report.neu.edu.cn/notes/create", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,* /*;q=0.8")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://pass.neu.edu.cn/")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	length = len(resp.Cookies())
	for i := 0; i < length; i++ {
		if resp.Cookies()[i].Name == "PHPSESSID" {
			length2 := len(xsrf)
			for j := 0; j < length2; j++ {
				if xsrf[j].Name == "PHPSESSID" {
					req.AddCookie(xsrf[j])
				}
			}
		} else {
			req.AddCookie(resp.Cookies()[i])
		}
	}
	xsrf = req.Cookies()
	resp, _ = client.Do(req)
	doc, _ = goquery.NewDocumentFromReader(resp.Body)
	str := doc.Find("a[id*=navbarDropdown]").Text()
	res := regexp.MustCompile(`当前用户：[^\s]{1,}`).Find([]byte(str))

	req, _ = http.NewRequest("GET", "https://e-report.neu.edu.cn/api/profiles/"+Uid+"?xingming="+string(res[15:]), nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,* /*;q=0.8")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://e-report.neu.edu.cn/notes/create")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "keep-alive")
	for i := 0; i < length; i++ {
		if xsrf[i].Name == "laravel3_session" {
			req.AddCookie(resp.Cookies()[1])
		} else if xsrf[i].Name == "XSRF-TOKEN" {
			req.AddCookie(resp.Cookies()[0])
		} else {
			req.AddCookie(xsrf[i])
		}
	}
	resp, _ = client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	req, _ = http.NewRequest("GET", "https://e-report.neu.edu.cn/notes", strings.NewReader(string(body)))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	req.Header.Set("Host", "e-report.neu.edu.cn")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,* /*;q=0.8")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://e-report.neu.edu.cn/notes/create")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "keep-alive")
	for i := 0; i < length; i++ {
		if xsrf[i].Name == "laravel3_session" {
			req.AddCookie(resp.Cookies()[1])
		} else if xsrf[i].Name == "XSRF-TOKEN" {
			req.AddCookie(resp.Cookies()[0])
		} else {
			req.AddCookie(xsrf[i])
		}
	}
	resp, _ = client.Do(req)
	return resp.StatusCode == 200
}
