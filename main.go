package main

import (
	"./tumblr"
	"os"
	"fmt"
	"bufio"
	"io"
	"sync"
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
	w := &sync.WaitGroup{}

	for _, site := range siteArray {
		w.Add(1)
		go func(site string) {
			t.StartDownload(w, site)
		}(site)
	}
	w.Wait()
}
