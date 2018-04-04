package tumblr

import (
	"regexp"
	"fmt"
	. "../common"
	"io/ioutil"
	"encoding/xml"
	"strings"
	"os"
)

type IVideo interface {
	VideoHd(url string) bool
	VideoDefault(url string) bool
}

type Match struct {
}

type TumblrCrawler struct {
	Queue      chan string
	VideoQueue chan string
}

type Post struct {
	PhotoUrl []string `xml:"photo-url"`
	VideoUrl []string `xml:"video-player"`
}

type TumblrResponse struct {
	XMLName    xml.Name `xml:"tumblr"`
	PhotoPosts []Post   `xml:"posts>post"`
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
		t.Queue = make(chan string)
		t.VideoQueue = make(chan string)
	}
	return t
}

func (t *TumblrCrawler) DownloadPhotos(site string) {

	go t.downLoadMedia(GetPath(site))

	t.saveMedia2Queue(site, "photo")
}

func (t *TumblrCrawler) DownloadVideo(site string) {

	go t.downLoadMedia(GetPath(site))

	t.saveMedia2Queue(site, "video")
}

func (t *TumblrCrawler) saveMedia2Queue(site string, mediaType string) {
	baseUrl := "http://%s.tumblr.com/api/read?type=%s&num=%d&start=%d"

	start := START
	for true {
		mediaUrl := fmt.Sprintf(baseUrl, site, mediaType, MEDIA_NUM, start)
		resp := ProxyHttpGet(mediaUrl)
		defer resp.Body.Close()
		if resp.StatusCode == 404 {
			break
		}
		body, _ := ioutil.ReadAll(resp.Body)
		result := new(TumblrResponse)
		xml.Unmarshal(body, result)
		for _, post := range result.PhotoPosts {
			if mediaType == "photo" {
				for _, photoUrl := range post.PhotoUrl {
					if strings.Contains(photoUrl, "avatar") {
						continue
					}
					t.Queue <- photoUrl
				}
			}
			if mediaType == "video" {
				for _, videoUrl := range post.VideoUrl {
					videoUrl = videoUrl[strings.Index(videoUrl, "https"):]
					source := videoUrl[:strings.Index(videoUrl, `"`)]
					t.VideoQueue <- source
				}
			}

		}
		start += MEDIA_NUM
	}
}

func (t *TumblrCrawler) downLoadMedia(dir string) {
	for i := range t.Queue { // chan关闭时，for循环会自动结束
		DownLoadImage(i, dir)
	}
}

func GetPath(site string) string {
	dir, _ := os.Getwd()
	path := dir + "/" + site
	_, err := os.Stat(path)
	if !os.IsExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	return path
}
