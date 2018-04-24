package tumblr

import (
	"fmt"
	. "github.com/suyuanhxx/crawler/common"
	"io/ioutil"
	"encoding/xml"
	"strings"
)

type TumblrCrawler struct {
	ImageChannel chan string
	VideoChannel chan string
}

var (
	MEDIA_NUM = 50
	//START     = 0
	BASE_URL = "http://%s.tumblr.com/api/read?type=%s&num=%d&start=%d"
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
	t.getSourceUrl(site, PHOTO)
}

func (t *TumblrCrawler) DownloadVideo(site string) {
	go t.downLoadMedia(site, VIDEO)
	t.getSourceUrl(site, VIDEO)
}

func (t *TumblrCrawler) getSourceUrl(site string, mediaType string) {
	start := 0
	for true {
		mediaUrl := fmt.Sprintf(BASE_URL, site, mediaType, MEDIA_NUM, start)
		resp := ProxyHttpGet(mediaUrl)
		if resp.StatusCode == 404 || resp == nil {
			break
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		result := new(TumblrResponse)
		xml.Unmarshal(body, result)
		if len(result.Posts) == 0 || result.Posts == nil {
			break
		}
		for _, post := range result.Posts {
			t.parseUrl(mediaType, post)
		}
		start += MEDIA_NUM
	}
	fmt.Println("getSourceUrl finish!")
}

func (t *TumblrCrawler) parseUrl(mediaType string, post Post) {
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
			p, source := ParseVideoUrl(videoUrl)
			if p {
				t.VideoChannel <- source
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
