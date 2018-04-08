package tumblr

import (
	"fmt"
	. "../common"
	"io/ioutil"
	"encoding/xml"
	"strings"
	"os"
)

type TumblrCrawler struct {
	ImageQueue chan string
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

var (
	MEDIA_NUM = 50
	START     = 0
)

func New() (*TumblrCrawler) {
	t := new(TumblrCrawler)
	if t.ImageQueue == nil {
		t.ImageQueue = make(chan string)
		t.VideoQueue = make(chan string)
	}
	return t
}

func (t *TumblrCrawler) DownloadPhotos(site string) {
	go t.downLoadMedia(GetPath(site), PHOTO)
	t.saveMedia2Queue(site, PHOTO)
}

func (t *TumblrCrawler) DownloadVideo(site string) {
	go t.downLoadMedia(GetPath(site), VIDEO)
	t.saveMedia2Queue(site, VIDEO)
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
					t.ImageQueue <- photoUrl
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

func (t *TumblrCrawler) downLoadMedia(dir string, mediaType string) {
	if mediaType == PHOTO {
		for i := range t.ImageQueue { // chan关闭时，for循环会自动结束
			DownLoadMedia(i, dir, mediaType)
		}
	}
	if mediaType == VIDEO {
		for i := range t.VideoQueue { // chan关闭时，for循环会自动结束
			DownLoadMedia(i, dir, mediaType)
		}
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
