package common

import (
	"os"
	"strings"
	"fmt"
	"sync"
	"net/http"
	"io/ioutil"
	"io"
	"bytes"
)

var (
	PHOTO = "photo"
	VIDEO = "video"
)

func DownLoadMedia(w *sync.WaitGroup, url, site, mediaType string) {
	if len(url) == 0 {
		return
	}
	fmt.Println(url)
	downloadMedia(url, site, mediaType)
	w.Done()
}

func downloadMedia(url, site, mediaType string) {
	//resp,err := ProxyHttpGet(url)
	resp, err := http.Get(url)
	if err != nil || resp == nil || resp.Body == nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("DownLoad error:", url, body)
		return
	}
	defer resp.Body.Close()

	var filename string
	if mediaType == PHOTO {
		filename = getImageName(url)
	} else if mediaType == VIDEO {
		filename = getVideoName(url)
	}

	out, e := os.Create(GetPath(site, mediaType) + filename)
	if e != nil {
		fmt.Println("create file error:", url, out)
		return
	}
	io.Copy(out, bytes.NewReader(body))
}

func getImageName(imageUrl string) string {
	i := strings.LastIndex(imageUrl, "/")
	return string(imageUrl[i:])
}

func getVideoName(videoUrl string) string {
	i := strings.LastIndex(videoUrl, "/")
	return string(videoUrl[i:]) + ".mp4"
}

func GetPath(site string, mediaType string) string {
	dir, _ := os.Getwd()
	path := dir + "/" + site + "/" + mediaType
	_, err := os.Stat(path)
	if !os.IsExist(err) {
		if e := os.MkdirAll(path, os.ModePerm); e == nil {
		}
	}
	return path
}
