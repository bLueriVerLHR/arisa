package plugins

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func MyCoser(msg string) string {
	if regexp.MustCompile(`mycoser`).Match([]byte(msg)) {
		var collect string
		var mod int
		if regexp.MustCompile(`海外|国外|外国`).Match([]byte(msg)) {
			collect = "overseas"
			mod = 30
		} else {
			collect = "lists"
			mod = 60
		}
		rand.Seed(time.Now().Unix())
		resp, _ := http.PostForm("http://mycoser.com/picture/"+collect+"/p/"+strconv.Itoa(rand.Int()%mod+1), url.Values{})
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		defer resp.Body.Close()
		setter := rand.Int() % 40
		var ret string
		var exists bool
		doc.Find("div[class=workimage]>a[href*=picture]").Each(func(i int, s *goquery.Selection) {
			if setter == i {
				ret, exists = s.Attr("href")
				if exists {
					fmt.Println(ret)
				}
				return
			}
		})
		resp, _ = http.PostForm(`http://mycoser.com`+ret, url.Values{})
		doc, _ = goquery.NewDocumentFromReader(resp.Body)
		var allpic []string
		doc.Find("img[layer-src*=thumbMid]").Each(func(i int, s *goquery.Selection) {
			ret, exists = s.Attr("layer-src")
			if exists {
				allpic = append(allpic, ret)
			}
		})
		return "http://mycoser.com" + allpic[rand.Int()%len(allpic)]
	}
	return ""
}
