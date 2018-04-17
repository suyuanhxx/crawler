package common

import (
	"testing"
)

func Test_DownLoadImage(t *testing.T) {
	img := "http://img.zcool.cn/community/018d4e554967920000019ae9df1533.jpg"

	DownLoadMedia(img, "travelingcolors", "photo")
}

func Test_DownLoadVideo(t *testing.T) {

	video := "https://vtt.tumblr.com/tumblr_p3foikwNKa1qjvnc4"

	DownLoadMedia(video, "travelingcolors", "video")
}
