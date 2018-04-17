package main

import (
	"./tumblr"
	//. "./common"
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
	for {
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
		go func() {
			go t.DownloadVideo(site)
			image <- 1
		}()
		go func() {
			t.DownloadPhotos(site)
			video <- 1
		}()
	}
	<-image
	<-video
}
