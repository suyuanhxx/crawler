package tumblr

import (
	"encoding/xml"
	"strings"
	"fmt"
)

type Tumblr struct {
	XMLName   xml.Name `xml:"tumblr"`
	Tumblelog string   `xml:"tumblelog"`
}

type Tumblelog struct {
	Title string `xml:"title,attr"`
	Name  string `xml:"name,attr"`
}

type TumblrResponse struct {
	Tumblr
	Posts []Post `xml:"posts>post"`
}

type Post struct {
	PhotoUrl []string `xml:"photo-url"`
	VideoUrl []string `xml:"video-player"`
}

type VideoPlayer struct {
	XMLName xml.Name    `xml:"video"`
	Data    string      `xml:"data-crt-options"`
	Source  VideoSource `xml:"source"`
}

type VideoSource struct {
	Src  string `xml:"src,attr"`
	Type string `xml:"type,attr"`
}

func ParseVideoUrl(url string) (bool, string) {
	player := VideoPlayer{}

	if strings.Contains(url, "instagram-media") {
		return false, ""
	}

	if strings.Contains(url, "iframe") {
		return false, ""
	}

	playerString := strings.Replace(url, "&lt;", "<", -1)
	playerString = strings.Replace(playerString, "&gt;", ">", -1)
	playerString = strings.Replace(playerString, "'", "\"", -1)
	playerString = strings.Replace(playerString, "\r<", "<", -1)
	playerString = strings.Replace(playerString, "\n<", "<", -1)
	playerString = strings.Replace(playerString, ">\r", ">", -1)
	playerString = strings.Replace(playerString, ">\n", ">", -1)
	playerString = strings.Replace(playerString, "{\"", "{'", -1)
	playerString = strings.Replace(playerString, "\"}", "'}", -1)
	playerString = strings.Replace(playerString, "\":", "':", -1)
	playerString = strings.Replace(playerString, ",\"", ",'", -1)
	playerString = strings.Replace(playerString, "\",", "',", -1)
	playerString = strings.Replace(playerString, ":\"", ":'", -1)
	playerString = strings.Replace(playerString, "\">", "\"/>", -1)
	playerString = strings.Replace(playerString, "muted data-crt-video ", " ", -1)

	err := xml.Unmarshal([]byte(playerString), &player)

	if err != nil {
		fmt.Println(err, "playerString", playerString)
		return false, ""
	}

	return true, player.Source.Src
}
