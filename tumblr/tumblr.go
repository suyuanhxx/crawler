package tumblr

import (
	"fmt"
	. "github.com/suyuanhxx/crawler/common"
	"io/ioutil"
	"encoding/xml"
	"sync"
	"strings"
	//"net/http"
)

type TumblrCrawler struct {
	waitGroup *sync.WaitGroup
}

var (
	MEDIA_NUM = 50
	//START     = 0
	BASE_URL = "http://%s.tumblr.com/api/read?type=%s&num=%d&start=%d"
)

func New() (*TumblrCrawler) {
	t := new(TumblrCrawler)
	t.waitGroup = &sync.WaitGroup{}
	return t
}

func (t *TumblrCrawler) StartDownload(w *sync.WaitGroup, site string) {
	t.fetchSource(site, PHOTO)
	t.fetchSource(site, VIDEO)
	w.Done()
}

func (t *TumblrCrawler) fetchSource(site string, mediaType string) {
	start := 0

	for true {
		mediaUrl := fmt.Sprintf(BASE_URL, site, mediaType, MEDIA_NUM, start)
		resp,err := ProxyHttpGet(mediaUrl)
		//resp, err := http.Get(mediaUrl)
		if resp.StatusCode == 404 || resp == nil || err != nil {
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
			t.downLoad(site, mediaType, post)
		}
		start += MEDIA_NUM
	}
	fmt.Println("fetchSource finish!")
}

func (t *TumblrCrawler) downLoad(site, mediaType string, post Post) {
	switch mediaType {
	case PHOTO:
		t.downLoadPhoto(site, post)
	case VIDEO:
		t.downLoadVideo(site, post)
	}
}

func (t *TumblrCrawler) downLoadPhoto(site string, post Post) {
	for _, photoUrl := range post.PhotoUrl {
		if strings.Contains(photoUrl, "avatar") {
			continue
		}
		t.waitGroup.Add(1)
		go func(w *sync.WaitGroup, photoUrl, site string) {
			DownLoadMedia(w, photoUrl, site, PHOTO)
		}(t.waitGroup, photoUrl, site)
	}
	t.waitGroup.Wait()
}

func (t *TumblrCrawler) downLoadVideo(site string, post Post) {
	for _, videoUrl := range post.VideoUrl {
		ok, source := ParseVideoUrl(videoUrl)
		if ok {
			t.waitGroup.Add(1)
			go func(w *sync.WaitGroup, source, site string) {
				DownLoadMedia(w, source, site, VIDEO)
			}(t.waitGroup, source, site)
		}
	}
	t.waitGroup.Wait()
}
