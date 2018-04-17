package tumblr

import (
	"fmt"
	. "../common"
	"io/ioutil"
	"encoding/xml"
	"strings"
)

type TumblrCrawler struct {
	ImageChannel chan string
	VideoChannel chan string
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
	if t.ImageChannel == nil {
		t.ImageChannel = make(chan string, 50)
	}
	if t.VideoChannel == nil {
		t.VideoChannel = make(chan string, 50)
	}
	return t
}

func (t *TumblrCrawler) DownloadPhotos(site string) {
	go t.downLoadMedia(site, PHOTO)
	t.saveMedia(site, PHOTO)
}

func (t *TumblrCrawler) DownloadVideo(site string) {
	go t.downLoadMedia(site, VIDEO)
	t.saveMedia(site, VIDEO)
}

func (t *TumblrCrawler) saveMedia(site string, mediaType string) {
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
			t.resolveUrl(mediaType, post)
		}
		start += MEDIA_NUM
	}
}

func (t *TumblrCrawler) resolveUrl(mediaType string, post Post) {
	switch mediaType {
	case "photo":
		for _, photoUrl := range post.PhotoUrl {
			if strings.Contains(photoUrl, "avatar") {
				continue
			}
			t.ImageChannel <- photoUrl
		}
	case "video":
		for _, videoUrl := range post.VideoUrl {
			i := strings.Index(videoUrl, "https")
			if i < len(videoUrl) && i > 0 {
				videoUrl = videoUrl[i:]
			}
			j := strings.Index(videoUrl, `"`)
			if j < len(videoUrl) && j > 0 {
				videoUrl = videoUrl[:j]
			}
			end := strings.LastIndex(videoUrl, "/")
			if end < len(videoUrl) && end > 0 {
				source := "https://vtt.tumblr.com" + videoUrl[end:]
				if len(source) > 0 {
					t.VideoChannel <- source
				}
			}
		}
	}
}

func (t *TumblrCrawler) downLoadMedia(site string, mediaType string) {
	switch mediaType {
	case "photo":
		for url := range t.ImageChannel {
			DownLoadMedia(url, site, PHOTO)
		}
	case "video":
		for url := range t.VideoChannel {
			DownLoadMedia(url, site, VIDEO)
		}
	}
}
