package main

import (
	"./tumblr"
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	Start()
}

func Start() {
	siteFile, err := os.Open("sites.txt")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	var siteArray []string
	br := bufio.NewReader(siteFile)
	for true {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		siteArray = append(siteArray, string(a))
	}

	t := tumblr.New()

	image := make(chan int, 1)
	video := make(chan int, 1)

	for _, site := range siteArray {
		go func(site string) {
			t.DownloadVideo(site)
			image <- 1
		}(site)
		go func(site string) {
			t.DownloadPhotos(site)
			video <- 1
		}(site)
	}
	<-image
	<-video
	close(t.ImageChannel)
	close(t.VideoChannel)
}
