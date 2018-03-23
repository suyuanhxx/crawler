package main

import (
	tumblr "./tumblr"
	. "./common"
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {

	img := "http://img.zcool.cn/community/018d4e554967920000019ae9df1533.jpg"

	DownLoadImage(img)


	siteFile, err := os.Open("sites.txt")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	siteArray := []string{}

	br := bufio.NewReader(siteFile)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		siteArray = append(siteArray, string(a))
	}
	t := tumblr.New()
	for _, site := range siteArray {
		t.DownloadPhotos(site)
	}

}
