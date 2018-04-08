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

func DownLoadMedia(url string, dir string, mediaType string) {
	fmt.Println(url)
	resp := ProxyHttpGet(url)
	body, _ := ioutil.ReadAll(resp.Body)

	var filename string
	if mediaType == PHOTO {
		filename = getImageName(url)
	} else if mediaType == VIDEO {
		filename = getVideoName(url)
	}

	out, _ := os.Create(dir + filename)
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
