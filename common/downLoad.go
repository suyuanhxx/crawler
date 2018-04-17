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
	fmt.Println(url)
	if len(url) == 0 {
		return
	}
	resp := ProxyHttpGet(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var filename string
	if mediaType == PHOTO {
		filename = getImageName(url)
	} else if mediaType == VIDEO {
		filename = getVideoName(url)
	}

	out, _ := os.Create(GetPath(site, mediaType) + filename)
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
			fmt.Println()
		}
	}
	return path
}
