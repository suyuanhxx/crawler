package tumblr

import (
	"regexp"
	"container/list"
	"fmt"
	. "../common"
	"io/ioutil"
	"encoding/xml"
	"strings"
)

type Video interface {
	VideoHd(url string) bool
	VideoDefault(url string) bool
}

type Match struct {
}

type TumblrCrawler struct {
	Queue *list.List
}

type Post struct {
	PhotoUrl []string `xml:"photo-url"`
}

type TumblrResponse struct {
	XMLName xml.Name `xml:"tumblr"`
	Posts   []Post   `xml:"posts>post"`
}

var hdPattern = regexp.MustCompile(`'.*"hdUrl":("([^\s,]*)"|false),`)

var defaultPattern = regexp.MustCompile(`.*src="(\S*)" `)

var MEDIA_NUM = 50
var START = 0

func (m *Match) VideoHd(url string) bool {
	return hdPattern.Match([]byte(url))
}

func (m *Match) VideoDefault(url string) bool {
	return defaultPattern.Match([]byte(url))
}

func New() (*TumblrCrawler) {
	t := new(TumblrCrawler)
	if t.Queue == nil {
		t.Queue = list.New()
	}
	return t
}

func (t *TumblrCrawler) DownloadPhotos(site string) {
	baseUrl := "http://%s.tumblr.com/api/read?type=%s&num=%d&start=%d"

	start := START
	for true {
		mediaUrl := fmt.Sprintf(baseUrl, site, "photo", MEDIA_NUM, start)
		resp := ProxyHttpGet(mediaUrl)
		defer resp.Body.Close()
		if resp.StatusCode == 404 {
			break
		}
		body, _ := ioutil.ReadAll(resp.Body)
		result := new(TumblrResponse)
		xml.Unmarshal(body, result)
		for _, post := range result.Posts {
			for _, photoUrl := range post.PhotoUrl {
				if strings.Contains(photoUrl, "avatar") {
					continue
				}
				t.Queue.PushBack(photoUrl)
			}
		}
		start += MEDIA_NUM
	}
}
