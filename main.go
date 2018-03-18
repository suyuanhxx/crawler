package main

import (
	tumblr "./tumblr"
)

func main() {
	site := "musclecorps"
	//
	//baseUrl := "http://%s.tumblr.com/api/read?type=%s&num=%d&start=%d"
	//
	//mediaUrl := fmt.Sprintf(baseUrl, site, "photo", 50, 0)
	//
	//resp := ProxyHttpGet(mediaUrl)
	//
	//fmt.Print(resp)

	t := tumblr.New()
	t.DownloadPhotos(site)

}
