package common

import (
	"io/ioutil"
	"os"
	"io"
	"bytes"
	"strings"
	"fmt"
)

var (
	PHOTO = "photo"
	VIDEO = "video"
)

func DownLoadMedia(url string, site string, mediaType string) {
	if len(url) == 0 {
		return
	}
	fmt.Println(url)
	resp := ProxyHttpGet(url)
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
